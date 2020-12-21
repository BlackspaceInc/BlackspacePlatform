package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	svcErrors "github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/errors"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/grpc/proto"
	"github.com/itimofeev/go-saga"
	opentracing "github.com/opentracing/opentracing-go"
)

/*
TODO
	1) emit metrics
		- latency, failures, success
	2) Resolver retain a reference to metrics object, tracing object, and logging object
	3) Handle password changes
	4) send out emails on account creation, ... etc
 */

func (r *mutationResolver) CreateBusinessAcount(ctx context.Context, input proto.CreateBusinessAccountRequest) (*proto.BusinessAccount, error) {
	r.Db.Logger.For(ctx).Info(fmt.Sprintf("create business accounts api op"))
	sp, ctx := opentracing.StartSpanFromContext(ctx, "create_business_account_api_op")
	defer sp.Finish()

	if &input == nil || input.BusinessAccount == nil || input.BusinessAccount.Validate() != nil {
		r.Db.Logger.For(ctx).Error(svcErrors.ErrInvalidInputArguments, svcErrors.ErrInvalidInputArguments.Error())
		return nil, svcErrors.ErrInvalidInputArguments
	}

	// attempt to obtain business account form backend based on email
	account := r.Db.GetBusinessByEmail(ctx, input.BusinessAccount.Email)
	if account != nil {
		// attempt to see if account is inactive
		if !account.IsActive {
			// we reactivate the account via a saga
			distributedUnlockOpStep := saga.Step{
				Name:           "unlock_business_account_distributed_tx",
				Func:           r.Db.DistributedTxUnlockAccount(ctx, account.AuthnId, sp),
				CompensateFunc: r.Db.DistributedTxLockAccount(ctx, account.AuthnId, sp),
				Options:        nil,
			}

			// activate account activity status
			reactivateAccountOpStep := saga.Step{
				Name:           "reactivate_business_account",
				Func:           r.Db.SetBusinessAccountStatusAndSave(ctx, account, true),
				CompensateFunc: r.Db.SetBusinessAccountStatusAndSave(ctx, account, false),
				Options:        nil,
			}

			if err := r.Db.Saga.RunSaga(ctx, "unlock_business_account", distributedUnlockOpStep, reactivateAccountOpStep); err != nil {
				r.Db.Logger.For(ctx).Error(err, err.Error())
				return nil, err
			}

			return account, nil
		}
	}

	// now we attempt to create the account from the current context of this service since for this api to be invoked the api
	// gateway must have already created an entry in the authentication handler service referencing this account record
	// perform this operation as a retryable one
	account, err := r.Db.CreateBusinessAccount(ctx, account, input.AuthnId)
	if err != nil {
		r.Db.Logger.For(ctx).Error(err, err.Error())
		return nil, err
	}

	return account, nil
}

func (r *mutationResolver) UpdateBusinessAccount(ctx context.Context, input proto.UpdateBusinessAccountRequest) (*proto.BusinessAccount, error) {
	r.Db.Logger.For(ctx).Info(fmt.Sprintf("update business account api op"))
	sp, ctx := opentracing.StartSpanFromContext(ctx, "update_business_account_api_op")
	defer sp.Finish()

	// validate the input
	if &input == nil ||
		input.UpdatedBusinessAccount == nil ||
		input.UpdatedBusinessAccount.Validate() != nil ||
		input.Id == 0 ||
		input.Validate() != nil {
		r.Db.Logger.For(ctx).Error(svcErrors.ErrInvalidInputArguments, svcErrors.ErrInvalidInputArguments.Error())
		return nil, svcErrors.ErrInvalidInputArguments
	}

	var newBusinessAccount = input.UpdatedBusinessAccount
	var businessAccountId = input.Id

	// attempt obtain the business account stored in the backend db first
	oldBusinessAccount := r.Db.GetBusinessById(ctx, businessAccountId)
	if oldBusinessAccount == nil {
		r.Db.Logger.For(ctx).ErrorM(svcErrors.ErrAccountDoesNotExist, fmt.Sprintf("business account with id %d does not exist", businessAccountId))
		return nil, svcErrors.ErrAccountDoesNotExist
	}

	var transactionalSteps = make([]saga.Step, 3)
	var updatedAccount = make(chan *proto.BusinessAccount)
	// TODO: handle password updates via authentication handler service in the future - Need to implement this too
	// TODO: send out an email to the account owner that the email or password has been changed
	// define a saga step tailored to saving the new business account record in our backend db
	updateAndSaveAccountStep := saga.Step{
		Name: "update_business_account",
		Func: r.Db.UpdateBusinessAccount(ctx, businessAccountId, newBusinessAccount),
		CompensateFunc: func(ctx context.Context) error { // no compensating function just return an error if this operation fails
			return svcErrors.ErrFailedToSaveUpdatedAccountRecord
		},
	}
	transactionalSteps = append(transactionalSteps, updateAndSaveAccountStep)

	// check if the email is updated
	if oldBusinessAccount.Email != newBusinessAccount.Email {
		// perform distributed tx via saga
		var authnId = oldBusinessAccount.AuthnId
		var newEmail = newBusinessAccount.Email

		compensatingFunc := func(ctx context.Context, id uint32, account *proto.BusinessAccount) error {
			acc, err := r.Db.UpdateBusinessAccount(ctx, businessAccountId, oldBusinessAccount) // reset business account to original state
			if err != nil {
				return err
			}

			updatedAccount <- acc
			return nil
		}

		// update the email account from the context of the authentication handler service
		updateEmailInDtxStep := saga.Step{
			Name:           "update_business_account_email_distributed_tx",
			Func:           r.Db.DistributedTxUpdateAccountEmail(ctx, authnId, newEmail, sp),
			CompensateFunc: compensatingFunc(ctx, authnId, oldBusinessAccount),
		}

		transactionalSteps = append(transactionalSteps, updateEmailInDtxStep)
	}

	// run the saga
	if err := r.Db.Saga.RunSaga(ctx, "update_business_account", transactionalSteps...); err != nil {
		r.Db.Logger.For(ctx).Error(err, err.Error())
		return nil, err
	}

	return <-updatedAccount, nil
}

func (r *mutationResolver) DeleteBusinessAccount(ctx context.Context, id proto.DeleteBusinessAccountRequest) (bool, error) {
	r.Db.Logger.For(ctx).Info(fmt.Sprintf("delete business account api op"))
	sp, ctx := opentracing.StartSpanFromContext(ctx, "delete_business_account_api_op")
	defer sp.Finish()

	// validate the input
	if id.Id == 0 || id.Validate() != nil {
		r.Db.Logger.For(ctx).Error(svcErrors.ErrInvalidInputArguments, svcErrors.ErrInvalidInputArguments.Error())
		return false, svcErrors.ErrInvalidInputArguments
	}

	accountId := id.Id
	account := r.Db.GetBusinessById(ctx, id.Id)
	if account == nil {
		r.Db.Logger.For(ctx).Error(svcErrors.ErrAccountDoesNotExist, svcErrors.ErrAccountDoesNotExist.Error())
		return false, svcErrors.ErrAccountDoesNotExist
	}

	var (
		transactionalSteps                = make([]saga.Step, 3)
		dtxLockOpStep, archiveAccountStep saga.Step
	)

	// we perform this operation as a distributed transaction
	// since we never truly delete the account from our backend we set the record to inactive
	// while also ensuring from the context of the authentication handler service the account is locked
	// define saga

	// first operation is to perform a distributed transaction and lock the account if possible
	dtxLockOpStep = saga.Step{
		Name:           "distributed lock operation",
		Func:           r.Db.DistributedTxLockAccount(ctx, account.AuthnId, sp),
		CompensateFunc: r.Db.DistributedTxUnlockAccount(ctx, account.AuthnId, sp),
	}

	// second operation is to update the state of the account and save to database
	archiveAccountStep = saga.Step{
		Name:           "archive business account operation",
		Func:           r.Db.ArchiveBusinessAccount(ctx, accountId),              // archive business account
		CompensateFunc: r.Db.SetBusinessAccountStatusAndSave(ctx, account, true), // activate account
		Options:        nil,
	}
	transactionalSteps = append(transactionalSteps, dtxLockOpStep, archiveAccountStep)

	// run the saga
	if err := r.Db.Saga.RunSaga(ctx, "archive_business_account", dtxLockOpStep, archiveAccountStep); err != nil {
		r.Db.Logger.For(ctx).Error(err, err.Error())
		return false, err
	}

	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
