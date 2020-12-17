package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/grpc/proto"
)

func (r *mutationResolver) CreateBusinessAcount(ctx context.Context, input proto.CreateBusinessAccountRequest) (*proto.BusinessAccount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBusinessAccount(ctx context.Context, input proto.UpdateBusinessAccountRequest) (*proto.BusinessAccount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBusinessAccount(ctx context.Context, id proto.DeleteBusinessAccountRequest) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBusinessAccounts(ctx context.Context, ids proto.DeleteBusinessAccountsRequest) ([]*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
