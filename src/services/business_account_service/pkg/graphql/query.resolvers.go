package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	svcErrors "github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/errors"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql/generated"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/grpc/proto"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (r *queryResolver) GetBusinessAccount(ctx context.Context, input proto.GetBusinessAccountRequest) ([]*proto.BusinessAccount, error) {
	r.Db.Logger.For(ctx).Info(fmt.Sprintf("get business account api op"))
	sp, ctx := opentracing.StartSpanFromContext(ctx, "get_business_account_api_op")
	defer sp.Finish()

	// ensure input is not nil or misconfigured
	if &input == nil || input.Id == 0 {
		r.Db.Logger.For(ctx).Error(svcErrors.ErrInvalidInputArguments, svcErrors.ErrInvalidInputArguments.Error())
		return nil, svcErrors.ErrInvalidInputArguments
	}

	var accounts = make([]*proto.BusinessAccount, 1)

	account, err := r.Db.GetBusinessAccount(ctx, input.Id)
	if err != nil {
		r.Db.Logger.For(ctx).ErrorM(err, err.Error())
		return nil, err
	}

	accounts = append(accounts, account)

	r.Db.Logger.For(ctx).Info(fmt.Sprintf("successfully obtain business account - id ; %d", input.Id), zap.String("company name", account.CompanyName))
	return accounts, nil
}

func (r *queryResolver) GetBusinessAccounts(ctx context.Context, limit proto.GetBusinessAccountsRequest) ([]*proto.BusinessAccount, error) {
	r.Db.Logger.For(ctx).Info(fmt.Sprintf("get paginated business accounts api op"))
	sp, ctx := opentracing.StartSpanFromContext(ctx, "get_paginated_business_account_api_op")
	defer sp.Finish()

	if &limit == nil || limit.Limit == 0 {
		r.Db.Logger.For(ctx).Error(svcErrors.ErrInvalidInputArguments, svcErrors.ErrInvalidInputArguments.Error())
		return nil, svcErrors.ErrInvalidInputArguments
	}

	accounts, err := r.Db.GetPaginatedBusinessAccounts(ctx, limit.Limit)
	if err != nil {
		r.Db.Logger.For(ctx).Error(err, err.Error())
		return nil, err
	}

	r.Db.Logger.For(ctx).Info(fmt.Sprintf("successfully obtain %d business accounts", limit.Limit))
	return accounts, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
