package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/grpc/proto"
)

func (r *queryResolver) GetBusinessAccount(ctx context.Context, input proto.GetBusinessAccountRequest) ([]*proto.BusinessAccount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetBusinessAccounts(ctx context.Context, limit proto.GetBusinessAccountsRequest) ([]*proto.BusinessAccount, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
