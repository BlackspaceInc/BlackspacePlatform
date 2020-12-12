package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated"
	models1 "github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated/models"
)

func (r *mutationResolver) UpdateCompany(ctx context.Context, input models1.CompanyInput) (*models1.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCompany(ctx context.Context, input models1.CompanyInput) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCompany(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCompanies(ctx context.Context, ids []*int) ([]*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
