package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/giantswarm/retry-go"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	saga "github.com/itimofeev/go-saga"
	"gorm.io/gorm/clause"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/errors"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/grpc/proto"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/utils"
)

type dbOperationType string

const (
	CREATE_ONE  dbOperationType = "CREATE_ONE"
	UPDATE_ONE  dbOperationType = "UPDATE_ONE"
	DELETE_ONE  dbOperationType = "DELETE_ONE"
	GET_ONE     dbOperationType = "GET_ONE"
	CREATE_MANY dbOperationType = "CREATE_MANY"
	UPDATE_MANY dbOperationType = "UPDATE_MANY"
	DELETE_MANY dbOperationType = "DELETE_MANY"
	GET_MANY    dbOperationType = "GET_MANY"
)

type DatabaseOperations interface {
	// CreateBusinessAccount creates a business account
	CreateBusinessAccount(ctx context.Context, account *proto.BusinessAccount) (*proto.BusinessAccount, error)
	// UpdateBusinessAccount updates a business account by Id
	UpdateBusinessAccount(ctx context.Context, id uint32, account *proto.BusinessAccount) (*proto.BusinessAccount, error)
	// DeleteBusinessAccount deletes a business account by Id
	ArchiveBusinessAccount(ctx context.Context, id uint32) (bool, error)
	// DeleteBusinessAccounts deletes a set of business accounts by Id
	ArchiveBusinessAccounts(ctx context.Context, ids []uint32) ([]bool, error)
	// GetBusinessAccount gets a business account by id
	GetBusinessAccount(ctx context.Context, id uint32) (*proto.BusinessAccount, error)
	// GetBusinessAccounts gets a set of business accounts by id
	GetBusinessAccounts(ctx context.Context, ids []uint32) ([]*proto.BusinessAccount, error)
}

// UpdateBusinessAccount updates a business account
func (db *Db) UpdateBusinessAccount(ctx context.Context, id uint32, account *proto.BusinessAccount) (*proto.BusinessAccount, error){
	db.Logger.For(ctx).Info(fmt.Sprintf("updated business account - id : %d", id))
	ctx, span := db.startRootSpan(ctx, "get_business_accounts_db_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB)(interface{}, error){
		db.Logger.For(ctx).Info("starting update account operation for account with id :%id", id)
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "update_business_accounts_db_tx")
		defer childSpan.Finish()

		// attempt to see if account exists
		result := db.GetBusinessById(ctx, id)
		if result == nil {
			db.Logger.ErrorM(errors.ErrAccountDoesNotExist, errors.ErrAccountDoesNotExist.Error())
			return nil, errors.ErrAccountDoesNotExist
		}

		// TODO compare the passwords and if they differ update the field through /password auth handler service call
		// As of now we do not allow users the ability to update their passwords through this call
		if !db.ComparePasswords(result.Password, []byte(account.Password)){
			db.Logger.ErrorM(errors.ErrCannotUpdatePassword, errors.ErrCannotUpdatePassword.Error())
			return nil, errors.ErrCannotUpdatePassword
		}

		// convert account record to ORM type
		businessAccountOrm, err := account.ToORM(ctx)
		if err != nil {
			db.Logger.Error(errors.ErrFailedToConvertToOrmType, err.Error())
			return nil, err
		}

		// save the account
		if err := tx.Save(&businessAccountOrm).Error; err != nil {
			db.Logger.Error(errors.ErrFailedToSaveUpdatedAccountRecord, err.Error())
			return nil, err
		}

		updatedAccount, err := businessAccountOrm.ToPB(ctx)
		if err != nil {
			db.Logger.Error(errors.ErrFailedToConvertFromOrmType, err.Error())
			return nil, err
		}

		return &updatedAccount, nil
	}

	res, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return res.(*proto.BusinessAccount), nil
}

// DeleteBusinessAccounts deletes a set of business accounts by ids
func (db *Db) ArchiveBusinessAccounts(ctx context.Context, ids []uint32) ([]bool, error) {
	db.Logger.For(ctx).Info("delete business accounts")
	ctx, span := db.startRootSpan(ctx, "archive_business_accounts_db_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "archive_business_accounts_db_tx")
		defer childSpan.Finish()

		var deleteStatus = make([]bool, len(ids))
		for _, id := range ids {
			success, err := db.ArchiveBusinessAccount(ctx, id)
			if !success {
				db.Logger.For(ctx).Error(err, fmt.Sprintf("%s - for id %d ", errors.ErrFailedToDeleteBusinessAccount.Error(), id))
			}

			deleteStatus = append(deleteStatus, success)
		}

		return deleteStatus, nil
	}

	res, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return res.([]bool), nil
}

// DeleteBusinessAccount archives a business account by id
func (db *Db) ArchiveBusinessAccount(ctx context.Context, id uint32) (bool, error) {
	db.Logger.For(ctx).Info("delete business account")
	ctx, span := db.startRootSpan(ctx, "archive_business_account_db_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "archive_business_account_db_tx")
		defer childSpan.Finish()

		// check if business actually exists
		account := db.GetBusinessById(ctx, id)
		if account == nil {
			db.Logger.For(ctx).ErrorM(errors.ErrAccountDoesNotExist, errors.ErrAccountDoesNotExist.Error())
			return false, nil
		}

		// deactivate account activity status
		deactivateAccountOpStep := saga.Step{
			Name:           "deactivate_business_account",
			Func:           db.SetBusinessAccountStatusAndSave(ctx, account, false),
			CompensateFunc: db.SetBusinessAccountStatusAndSave(ctx, account, true),
			Options:        nil,
		}

		if err := db.Saga.RunSaga(ctx, "deactivate_business_account", deactivateAccountOpStep); err != nil {
			db.Logger.For(ctx).Error(err, err.Error())
			return nil, err
		}

		return true, nil
	}

	res, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return false, err
	}

	return res.(bool), nil
}

// GetBusinessAccount gets a set of business accounts by ids
func (db *Db) GetBusinessAccounts(ctx context.Context, ids []uint32) ([]*proto.BusinessAccount, error) {
	// define initial log entry
	db.Logger.For(ctx).Info("get business accounts")
	ctx, span := db.startRootSpan(ctx, "get_business_accounts_db_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "get business accounts - operation")
		defer childSpan.Finish()

		var accounts = make([]*proto.BusinessAccount, len(ids)+1)
		for _, id := range ids {
			account := db.GetBusinessById(ctx, id)
			if account == nil {
				db.Logger.For(ctx).Error(errors.ErrAccountDoesNotExist, fmt.Sprintf("%s - for id %d ", errors.ErrAccountDoesNotExist.Error(), id))
			} else {
				accounts = append(accounts, account)
			}
		}

		return accounts, nil
	}

	res, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return res.([]*proto.BusinessAccount), nil
}

func (db *Db) GetPaginatedBusinessAccounts(ctx context.Context, limit int64)([]*proto.BusinessAccount, error){
	// define initial log entry
	db.Logger.For(ctx).Info("get business accounts")
	ctx, span := db.startRootSpan(ctx, "get_paginated_business_accounts_db_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "get_paginated_business_account_db_tx")
		defer childSpan.Finish()

		var result = make([]*proto.BusinessAccount, limit)
		if err := db.PreloadTx(tx).Limit(limit).Find(&result).Error; err != nil {
			db.Logger.For(ctx).Error(errors.ErrUnableToObtainBusinessAccounts, err.Error())
			return nil, err
		}

		return result, nil
	}

	res, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return res.([]*proto.BusinessAccount), nil
}

// GetBusinessAccount gets a singular business account
func (db *Db) GetBusinessAccount(ctx context.Context, id uint32) (*proto.BusinessAccount, error) {
	// define initial log entry
	db.Logger.For(ctx).Info(fmt.Sprintf("get business account - id : %d", id))
	ctx, span := db.startRootSpan(ctx, "get_business_account_db_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "get_business_account_db_tx")
		defer childSpan.Finish()

		account := db.GetBusinessById(ctx, id)
		if account == nil {
			db.Logger.For(ctx).Error(errors.ErrAccountDoesNotExist, errors.ErrAccountDoesNotExist.Error())
			return nil, errors.ErrAccountDoesNotExist
		}

		db.Logger.For(ctx).Info("successfully obtained business account", zap.String("id", string(id)), zap.String("name", account.CompanyName))
		return account, nil
	}

	res, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return res.(*proto.BusinessAccount), nil
}

// GetBusinessById gets a business account by id from the backend database
func (db *Db) GetBusinessById(ctx context.Context, id uint32) *proto.BusinessAccount {
	// define initial log entry
	db.Logger.For(ctx).Info(fmt.Sprintf("get business account by id - id : %d", id))
	ctx, span := db.startRootSpan(ctx, "get_business_account_by_id_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "get_business_account_by_id_tx")
		defer childSpan.Finish()

		var businessAccountOrm proto.BusinessAccountORM
		// attempt to see if the record already exists
		recordNotFound := db.PreloadTx(tx).
							 Where(&proto.BusinessAccountORM{Id: id}).
							 First(&businessAccountOrm).RecordNotFound()
		if recordNotFound {
			db.Logger.For(ctx).Error(errors.ErrAccountDoesNotExist, "account does not exist")
			return nil, errors.ErrAccountDoesNotExist
		}

		// transform orm type to account type
		account, err := businessAccountOrm.ToPB(ctx)
		if err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToConvertFromOrmType, err.Error())
			return nil, err
		}

		db.Logger.For(ctx).Info("successfully obtained business account", zap.String("id", string(id)))
		return &account, nil
	}

	res, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil
	}

	return res.(*proto.BusinessAccount)
}

// GetBusinessByEmail gets a business account by email from the backend database
func (db *Db) GetBusinessByEmail(ctx context.Context, email string) *proto.BusinessAccount {
	// define initial log entry
	db.Logger.For(ctx).Info("get business account by email")

	ctx, span := db.startRootSpan(ctx, "get_business_account_by_email_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "get_business_account_by_email_tx")
		defer childSpan.Finish()

		var businessAccountOrm proto.BusinessAccountORM

		// attempt to see if the record already exists
		recordNotFound := db.PreloadTx(tx).
							 Where(&proto.BusinessAccountORM{Email: email}).
							 First(&businessAccountOrm).RecordNotFound()
		if recordNotFound {
			db.Logger.For(ctx).Error(errors.ErrAccountDoesNotExist, "account does not exist")
			return nil, errors.ErrAccountDoesNotExist
		}

		// transform orm type to account type
		account, err := businessAccountOrm.ToPB(ctx)
		if err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToConvertFromOrmType, err.Error())
			return nil, err
		}

		db.Logger.For(ctx).Info("successfully obtained business account", zap.String("email", email))
		return &account, nil
	}

	res, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil
	}

	return res.(*proto.BusinessAccount)
}

// CreateBusinessAccount creates a business account and saves it to the database
func (db *Db) CreateBusinessAccount(ctx context.Context, account *proto.BusinessAccount, authnid uint32) (*proto.BusinessAccount, error) {
	db.Logger.For(ctx).Info("creating business account")
	ctx, span := db.startRootSpan(ctx, "create_business_account_op")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting transaction")
		span := db.TracingEngine.CreateChildSpan(ctx, "create_business_account_tx")
		defer span.Finish()

		// validate account object
		if err := account.Validate(); err != nil {
			db.Logger.Error(errors.ErrInvalidAccount, err.Error())
			return nil, err
		}

		var businessAccount proto.BusinessAccountORM
		// attempt to see if the record already exists
		// no 2 records in our backend database can have the same email or company name
		recordNotFound := db.PreloadTx(tx).Where(&proto.BusinessAccountORM{Email: account.Email,
				CompanyName: account.CompanyName}).First(&businessAccount).
			RecordNotFound()

		if !recordNotFound{
			// account already exists
			db.Logger.ErrorM(errors.ErrAccountAlreadyExist, errors.ErrAccountAlreadyExist.Error())
			return nil, errors.ErrAccountAlreadyExist
		}

		// if the account does not exist we save it in the db
		// convert it first to orm type
		businessAccount, err := account.ToORM(ctx)
		if err != nil {
			db.Logger.ErrorM(errors.ErrFailedToConvertToOrmType, err.Error())
			return nil, err
		}

		// hash password
		if businessAccount.Password, err = db.ValidateAndHashPassword(businessAccount.Password); err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToHashPassword, err.Error())
			return nil, err
		}

		// activate account and assign authn id relation
		businessAccount.IsActive = true
		businessAccount.AuthnId = authnid

		// save the account record in the db
		if err := tx.Create(&businessAccount).Error; err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToCreateAccount, err.Error())
			return nil, err
		}

		createdAccount, err := businessAccount.ToPB(ctx)
		if err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToConvertFromOrmType, err.Error())
			return nil, err
		}

		return &createdAccount, nil
	}

	result, err := db.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	createdAccount := result.(*proto.BusinessAccount)
	return createdAccount, nil
}

func (db *Db) PreloadTx(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Media.SubscribedTopics").
		Preload(clause.Associations)
}

// SetBusinessAccountStatusAndSave updates the active status of a business account in the backend database
func (db *Db) SetBusinessAccountStatusAndSave(ctx context.Context, businessAccount *proto.BusinessAccount,
	activateAccount bool) error {

	db.Logger.For(ctx).Info(fmt.Sprintf("updating business account active status to %v", activateAccount))
	span := db.TracingEngine.CreateChildSpan(ctx, "update business account active status operation")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) error {
		// convert to orm type
		account, err := businessAccount.ToORM(ctx)
		if err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToConvertToOrmType, err.Error())
			return err
		}

		// set account active status
		account.IsActive = activateAccount
		err = tx.Save(&account).Error
		if err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToUpdateAccountActiveStatus, err.Error())
			return err
		}

		return nil
	}

	f := func() error {
		return db.PerformTransaction(ctx, tx)
	}

	// should perform this as a retryable operation in case of errors
	return  db.PerformRetryableOperation(f)
}

// DistributedTxUnlockAccount unlocks an account in a distributed transaction
func (db *Db) DistributedTxUnlockAccount(ctx context.Context, id uint32, childSpan opentracing.Span) error {
	f :=  func() error {
		subSpan, ctx := opentracing.StartSpanFromContext(ctx, "unlock account", opentracing.ChildOf(childSpan.Context()))
		defer subSpan.Finish()

		// perform call to the authentication handler service
		httpClient := &http.Client{}
		url := db.AuthenticationHandlerServiceBaseEndpoint + "/unlock/" + fmt.Sprint(id)
		httpReq, _ := http.NewRequest("POST", url, nil)

		// Transmit the span's TraceContext as HTTP headers on our
		// outbound request.
		_ = opentracing.GlobalTracer().Inject(
			childSpan.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(httpReq.Header))

		resp, err := httpClient.Do(httpReq)
		if err != nil {
			db.Logger.For(ctx).Error(errors.ErrDistributedTransactionError, err.Error())
			return err
		}

		if resp.StatusCode != http.StatusOK {
			db.Logger.For(ctx).Error(errors.ErrDistributedTransactionError, errors.ErrDistributedTransactionError.Error()+" authentication handler service")
			return errors.ErrDistributedTransactionError
		}
		return nil
	}

	return db.PerformRetryableOperation(f)
}

func (db *Db) PerformRetryableOperation(f func() error) error {
	return retry.Do(func() error {
		return f()
	},
		retry.MaxTries(db.MaxRetriesPerOperation),
		// TODO: emit metrics
		// retry.AfterRetryLimit()
		retry.Timeout(db.RetryTimeOut),
		retry.Sleep(db.OperationSleepInterval),
	)
}

// DistributedTxLockAccount locks an account in a distributed transaction
func (db *Db) DistributedTxLockAccount(ctx context.Context, id uint32, childSpan opentracing.Span) error {
	f := func() error {
		subSpan, ctx := opentracing.StartSpanFromContext(ctx, "lock account", opentracing.ChildOf(childSpan.Context()))
		defer subSpan.Finish()

		// perform call to the authentication handler service
		httpClient := &http.Client{}
		url := db.AuthenticationHandlerServiceBaseEndpoint + "/lock/" + fmt.Sprint(id)
		httpReq := db.createRequestAndPropagateTraces(url, subSpan, nil)

		resp, err := db.performHttpRequest(httpClient, httpReq, ctx)
		if err != nil || db.processResponseStatusCode(resp, ctx) {
			// todo: retry the operation with exponential backoff
			return err
		}

		return nil
	}

	return db.PerformRetryableOperation(f)
}

// DistributedTxUpdateAccountEmail updates the account record's email entry in a distributed transaction
func (db *Db) DistributedTxUpdateAccountEmail(ctx context.Context, id uint32, email string, childSpan opentracing.Span) error {
	f := func() error {
		subSpan, ctx := opentracing.StartSpanFromContext(ctx, "lock account", opentracing.ChildOf(childSpan.Context()))
		defer subSpan.Finish()

		// perform call to the authentication handler service
		httpClient := &http.Client{}
		url := db.AuthenticationHandlerServiceBaseEndpoint + "/update/" + fmt.Sprint(id)

		type updateEmailReq struct {
			Email string `json:"email"`
		}

		reqBody := updateEmailReq{
			Email: email,
		}

		body, err := utils.CreateRequestBody(reqBody)
		if err != nil {
			return err
		}

		httpReq := db.createRequestAndPropagateTraces(url, subSpan, body)

		resp, err := db.performHttpRequest(httpClient, httpReq, ctx)
		if err != nil || db.processResponseStatusCode(resp, ctx) {
			// todo: retry the operation with exponential backoff
			return err
		}

		return nil
	}

	return db.PerformRetryableOperation(f)
}

// DistributedTxCreateAccount creates the account record in a distributed transaction
func (db *Db) DistributedTxCreateAccount(ctx context.Context, email, password string, childSpan opentracing.Span) (*uint32, error) {
	return func() (*uint32, error) {
		subSpan, ctx := opentracing.StartSpanFromContext(ctx, "create account", opentracing.ChildOf(childSpan.Context()))
		defer subSpan.Finish()

		// perform call to the authentication handler service
		httpClient := &http.Client{}
		url := db.AuthenticationHandlerServiceBaseEndpoint + "/create"

		type createAccountReq struct {
			Email string `json:"email"`
			Password string `json:"password"`
		}

		type createAccountRes struct {
			Error string `json:"error"`
			Id uint32 `json:"id"`
		}

		var result createAccountRes
		reqBody := createAccountReq{
			Email: email,
			Password: password,
		}

		body, err := utils.CreateRequestBody(reqBody)
		if err != nil {
			return nil, err
		}

		httpReq := db.createRequestAndPropagateTraces(url, subSpan, body)

		resp, err := db.performHttpRequest(httpClient, httpReq, ctx)
		if err != nil || db.processResponseStatusCode(resp, ctx) {
			// todo: retry the operation with exponential backoff
			return nil, err
		}

		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}

		// update the id by reference
		return &result.Id, nil
	}()
}

// createRequestAndPropagateTraces creates a request and propagates the contexts to tracing engine
func (db *Db) createRequestAndPropagateTraces(url string, childSpan opentracing.Span, body io.Reader) *http.Request {
	httpReq, _ := http.NewRequest("POST", url, body)

	// Transmit the span's TraceContext as HTTP headers on our
	// outbound request.
	opentracing.GlobalTracer().Inject(
		childSpan.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(httpReq.Header))
	return httpReq
}

// processResponseStatusCode processes response code and return wether the status code was faulty or not
func (db *Db) processResponseStatusCode(resp *http.Response, ctx context.Context) bool {
	if resp.StatusCode != http.StatusOK {
		db.Logger.For(ctx).Error(errors.ErrDistributedTransactionError, errors.ErrDistributedTransactionError.Error()+" authentication handler service")
		return true
	}
	return false
}

// performHttpRequest performs an http request and returns the request response status and any occuring errors
func (db *Db) performHttpRequest(httpClient *http.Client, httpReq *http.Request, ctx context.Context) (*http.Response, error) {
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		db.Logger.For(ctx).Error(errors.ErrDistributedTransactionError, err.Error())
		return nil, err
	}
	return resp, nil
}

// startRootSpan starts a root span object
func (db *Db) startRootSpan(ctx context.Context, operationType dbOperationType) (context.Context, opentracing.Span) {
	return utils.StartRootOperationSpan(ctx, string(operationType), db.TracingEngine, db.Logger)
}
