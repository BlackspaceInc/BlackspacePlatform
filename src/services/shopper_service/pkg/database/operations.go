package database

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/errors"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/grpc/proto"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/utils"
	saga "github.com/itimofeev/go-saga"
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
	DeleteBusinessAccount(ctx context.Context, id uint32) (bool, error)
	// DeleteBusinessAccounts deletes a set of business accounts by Id
	DeleteBusinessAccounts(ctx context.Context, ids []uint32) ([]bool, error)
	// GetBusinessAccount gets a business account by id
	GetBusinessAccount(ctx context.Context, id uint32) (*proto.BusinessAccount, error)
	// GetBusinessAccounts gets a set of business accounts by id
	GetBusinessAccounts(ctx context.Context, ids []uint32) ([]*proto.BusinessAccount, error)
}

// UpdateBusinessAccount updates a business account
func (db *Db) UpdateBusinessAccount(ctx context.Context, id uint32, account *proto.BusinessAccount) (*proto.BusinessAccount, error){
	db.Logger.For(ctx).Info(fmt.Sprintf("updated business account - id : %d", id))
	ctx, span := db.startRootSpan(ctx, UPDATE_ONE)
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB)(interface{}, error){
		db.Logger.For(ctx).Info("starting update account operation for account with id :%id", id)
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "update business account")
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

		// figure out if the email or password fields differ
		// first we check email
		if result.Email != account.Email {
			failCaseFunc := func() error {
				return errors.ErrFailedToUpdateAccountEmail
			}

			updateEmailFunc := func(from *proto.BusinessAccount, to *proto.BusinessAccount) error{
				to.Email = from.Email
				return nil
			}

			// we initiate a saga to update the record's email entry in the authentication service
			dtxTx := saga.NewSaga("update account email")
			store := saga.New()

			err := dtxTx.AddStep(&saga.Step{
				Name:           "update account email entry in authentication handler service",
				Func:           db.DistributedTxUpdateAccountEmail(ctx, id, account.Email, childSpan),
				CompensateFunc: failCaseFunc(),
			})

			if err != nil {
				db.Logger.Error(errors.ErrFailedToConfigureSaga, errors.ErrFailedToConfigureSaga.Error())
				return nil, errors.ErrFailedToConfigureSaga
			}

			// then we update the email field of the resulting record
			err = dtxTx.AddStep(&saga.Step{
				Name:           "update email field of record stored in db",
				Func:           updateEmailFunc(result, account),
				CompensateFunc: failCaseFunc(),
			})

			if err != nil {
				db.Logger.Error(errors.ErrFailedToConfigureSaga, errors.ErrFailedToConfigureSaga.Error())
				return nil, errors.ErrFailedToConfigureSaga
			}

			// TODO refactor saga logic into generic impl and move to its own function
			coordinator := saga.NewCoordinator(ctx, ctx, dtxTx, store)
			if result := coordinator.Play(); result != nil && (len(result.CompensateErrors) > 0 || result.ExecutionError != nil) {
				// log the saga operation errors
				db.Logger.Error(errors.ErrSagaFailedToExecuteSuccessfully, errors.ErrSagaFailedToExecuteSuccessfully.Error(),
					zap.Errors("compensate error", result.CompensateErrors), zap.Error(result.ExecutionError))

				// construct error
				errMsg := fmt.Sprintf("compensate errors : %s , execution errors %s", zap.Errors("compensate error",
					result.CompensateErrors).String+zap.Error(result.ExecutionError).String)
				err := errors.NewError(errMsg)
				return nil, err
			}
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
func (db *Db) DeleteBusinessAccounts(ctx context.Context, ids []uint32) ([]bool, error) {
	db.Logger.For(ctx).Info("delete business accounts")
	ctx, span := db.startRootSpan(ctx, DELETE_MANY)
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "delete business accounts - operation")
		defer childSpan.Finish()

		var deleteStatus = make([]bool, len(ids))
		for _, id := range ids {
			success, err := db.DeleteBusinessAccount(ctx, id)
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
func (db *Db) DeleteBusinessAccount(ctx context.Context, id uint32) (bool, error) {
	db.Logger.For(ctx).Info("delete business account")
	ctx, span := db.startRootSpan(ctx, DELETE_ONE)
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "delete business account - operation")
		defer childSpan.Finish()

		// check if business actually exists
		account := db.GetBusinessById(ctx, id)
		if account == nil {
			db.Logger.For(ctx).ErrorM(errors.ErrAccountDoesNotExist, errors.ErrAccountDoesNotExist.Error())
			return false, nil
		}

		accountOrm, err := account.ToORM(ctx)
		if err != nil {
			db.Logger.For(ctx).ErrorM(errors.ErrFailedToConvertFromOrmType, errors.ErrFailedToConvertToOrmType.Error())
			return false, err
		}

		// since we never truly delete the account from our backend we set the record to inactive
		// while also ensuring from the context of the authentication handler service the account is locked
		// define saga
		dstTx := saga.NewSaga("archive business account")
		// define coordinator store
		store := saga.New()

		// first operation is to perform a distributed transaction and lock the account if possible
		err = dstTx.AddStep(&saga.Step{
			Name:           "distributed unlock operation",
			Func:           db.DistributedTxLockAccount(ctx, accountOrm.AuthnId, childSpan),
			CompensateFunc: db.DistributedTxUnlockAccount(ctx, accountOrm.AuthnId, childSpan),
		})

		if err != nil {
			db.Logger.Error(errors.ErrFailedToConfigureSaga, errors.ErrFailedToConfigureSaga.Error())
			return nil, errors.ErrFailedToConfigureSaga
		}

		// second operation is to update the state of the account and save to database
		err = dstTx.AddStep(&saga.Step{
			Name:           "update business account and save to db operation",
			Func:           db.SetBusinessAccountStatusAndSave(ctx, accountOrm, false), // deactivate account
			CompensateFunc: db.SetBusinessAccountStatusAndSave(ctx, accountOrm, true),  // activate account
			Options:        nil,
		})

		if err != nil {
			db.Logger.Error(errors.ErrFailedToConfigureSaga, errors.ErrFailedToConfigureSaga.Error())
			return nil, errors.ErrFailedToConfigureSaga
		}

		coordinator := saga.NewCoordinator(ctx, ctx, dstTx, store)
		if result := coordinator.Play(); result != nil && (len(result.CompensateErrors) > 0 || result.ExecutionError != nil) {
			// log the saga operation errors
			db.Logger.Error(errors.ErrSagaFailedToExecuteSuccessfully, errors.ErrSagaFailedToExecuteSuccessfully.Error(),
				zap.Errors("compensate error", result.CompensateErrors), zap.Error(result.ExecutionError))

			// construct error
			errMsg := fmt.Sprintf("compensate errors : %s , execution errors %s", zap.Errors("compensate error",
				result.CompensateErrors).String+zap.Error(result.ExecutionError).String)
			err := errors.NewError(errMsg)
			return false, err
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
	ctx, span := db.startRootSpan(ctx, GET_MANY)
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

// GetBusinessAccount gets a singular business account
func (db *Db) GetBusinessAccount(ctx context.Context, id uint32) (*proto.BusinessAccount, error) {
	// define initial log entry
	db.Logger.For(ctx).Info("get business account")
	ctx, span := db.startRootSpan(ctx, GET_ONE)
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "get business account - operation")
		defer childSpan.Finish()

		account := db.GetBusinessById(ctx, id)
		if account == nil {
			db.Logger.For(ctx).Error(errors.ErrAccountDoesNotExist, errors.ErrAccountDoesNotExist.Error())
			return nil, errors.ErrAccountDoesNotExist
		}

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
	db.Logger.For(ctx).Info("get business account by id")

	ctx, span := db.startRootSpan(ctx, GET_ONE)
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "get business account by id - operation")
		defer childSpan.Finish()

		var businessAccountOrm proto.BusinessAccountORM

		// attempt to see if the record already exists
		recordNotFound := tx.Where(&proto.BusinessAccountORM{Id: id}).First(&businessAccountOrm).RecordNotFound()
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

	ctx, span := db.startRootSpan(ctx, GET_ONE)
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, "get business account by email - operation")
		defer childSpan.Finish()

		var businessAccountOrm proto.BusinessAccountORM

		// attempt to see if the record already exists
		recordNotFound := tx.Where(&proto.BusinessAccountORM{Email: email}).First(&businessAccountOrm).RecordNotFound()
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

func (db *Db) CreateBusinessAccount(ctx context.Context, account *proto.BusinessAccount) (*proto.BusinessAccount, error) {
	// 4. emit metrics
	db.Logger.For(ctx).Info("creating business account")

	ctx, span := db.startRootSpan(ctx, CREATE_ONE)
	defer span.Finish()

	// check if user record exists and if so check active status
	// if not active, communicate with auth handler service and then activate and save the record locally
	// if active return with an error message
	// else create the record through the below call
	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		// start child span
		db.Logger.For(ctx).Info("starting transaction")

		span := db.TracingEngine.CreateChildSpan(ctx, "create record operation")
		defer span.Finish()

		// validate account object
		if err := account.Validate(); err != nil {
			db.Logger.Error(errors.ErrInvalidAccount, err.Error())
			return nil, err
		}

		var businessAccount proto.BusinessAccountORM
		// attempt to see if the record already exists
		// no 2 records in our backend database can have the same email or company name
		recordNotFound := tx.Where(&proto.BusinessAccountORM{Email: account.Email, CompanyName: account.CompanyName}).First(&businessAccount).
			RecordNotFound()
		// if the record exists and the user is not active,
		// we reactivate the account by updating the account from the perspectice of the authentication handler service after which we update the
		// status of the operation in our backend db. Note this entire operation must be implemented in a saga
		if !recordNotFound && !businessAccount.IsActive {
			// call the authentication handler service but first ensure we trace the request
			childSpan := opentracing.StartSpan("authentication handler service - unlock operation", opentracing.ChildOf(span.Context()))
			defer childSpan.Finish()

			// define saga
			dstTx := saga.NewSaga("unlock business account")
			// define coordinator store
			store := saga.New()

			// first operation is to perform a distributed transaction and unlock the account if possible
			err := dstTx.AddStep(&saga.Step{
				Name:           "distributed unlock operation",
				Func:           db.DistributedTxUnlockAccount(ctx, businessAccount.Id, childSpan),
				CompensateFunc: db.DistributedTxLockAccount(ctx, businessAccount.Id, childSpan),
			})

			if err != nil {
				db.Logger.Error(errors.ErrFailedToConfigureSaga, errors.ErrFailedToConfigureSaga.Error())
				return nil, errors.ErrFailedToConfigureSaga
			}

			// second operation is to update the state of the account and save to database
			err = dstTx.AddStep(&saga.Step{
				Name:           "update business account and save to db operation",
				Func:           db.SetBusinessAccountStatusAndSave(ctx, businessAccount, true),  // activate account
				CompensateFunc: db.SetBusinessAccountStatusAndSave(ctx, businessAccount, false), // deactivate account
				Options:        nil,
			})

			if err != nil {
				db.Logger.Error(errors.ErrFailedToConfigureSaga, errors.ErrFailedToConfigureSaga.Error())
				return nil, errors.ErrFailedToConfigureSaga
			}

			coordinator := saga.NewCoordinator(ctx, ctx, dstTx, store)
			if result := coordinator.Play(); result != nil && (len(result.CompensateErrors) > 0 || result.ExecutionError != nil) {
				// log the saga operation errors
				db.Logger.Error(errors.ErrSagaFailedToExecuteSuccessfully, errors.ErrSagaFailedToExecuteSuccessfully.Error(),
					zap.Errors("compensate error", result.CompensateErrors), zap.Error(result.ExecutionError))

				// construct error
				errMsg := fmt.Sprintf("compensate errors : %s , execution errors %s", zap.Errors("compensate error",
					result.CompensateErrors).String+zap.Error(result.ExecutionError).String)
				err := errors.NewError(errMsg)
				return nil, err
			}

			db.Logger.Info("account successfully created",
				zap.String("id", string(businessAccount.Id)),
				zap.String("business account name", businessAccount.CompanyName),
				zap.String("email", businessAccount.Email))

			return &businessAccount, nil
		} else if !recordNotFound && businessAccount.IsActive {
			// account already exists and is active
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

		// TODO perform saga and allow interactions with the authentication handler service
		// set the activity status
		if err = db.SetBusinessAccountStatusAndSave(ctx, businessAccount, true); err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToUpdateAccountActiveStatus, err.Error())
			return nil, err
		}

		// hash password
		if businessAccount.Password, err = db.ValidateAndHashPassword(businessAccount.Password); err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToHashPassword, err.Error())
			return nil, err
		}

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

// SetBusinessAccountStatusAndSave updates the active status of a business account in the backend database
func (db *Db) SetBusinessAccountStatusAndSave(ctx context.Context, businessAccount proto.BusinessAccountORM,
	accountActive bool) error {

	db.Logger.For(ctx).Info(fmt.Sprintf("updating business account active status to %v", accountActive))
	span := db.TracingEngine.CreateChildSpan(ctx, "update business account active status operation")
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) error {
		businessAccount.IsActive = accountActive
		err := tx.Save(&businessAccount).Error
		if err != nil {
			db.Logger.For(ctx).Error(errors.ErrFailedToUpdateAccountActiveStatus, err.Error())
			return err
		}

		return nil
	}

	return db.PerformTransaction(ctx, tx)
}

// DistributedTxUnlockAccount unlocks an account in a distributed transaction
func (db *Db) DistributedTxUnlockAccount(ctx context.Context, id uint32, childSpan opentracing.Span) error {
	return func() error {
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
	}()
}

// DistributedTxLockAccount locks an account in a distributed transaction
func (db *Db) DistributedTxLockAccount(ctx context.Context, id uint32, childSpan opentracing.Span) error {
	return func() error {
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
	}()
}

// DistributedTxUpdateAccountEmail updates the account record's email entry in a distributed transaction
func (db *Db) DistributedTxUpdateAccountEmail(ctx context.Context, id uint32, email string, childSpan opentracing.Span) error {
	return func() error {
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
