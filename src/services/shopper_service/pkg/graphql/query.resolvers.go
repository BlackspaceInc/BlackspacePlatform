package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated/models"
)

func (r *queryResolver) GetCompanies(ctx context.Context) ([]*models.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPaginatedCompanies(ctx context.Context, limit int) ([]*models.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCompany(ctx context.Context, id int) (*models.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCompaniesByBusinessStage(ctx context.Context, stage *string) ([]*models.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCompaniesByBusinessType(ctx context.Context, businessType *string) ([]*models.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCompaniesByMerchantType(ctx context.Context, merchantType *string) ([]*models.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
