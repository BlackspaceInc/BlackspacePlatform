package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/models"
)

func (r *companyResolver) ID(ctx context.Context, obj *models.Company) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mediaResolver) ID(ctx context.Context, obj *models.Media) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *topicsResolver) ID(ctx context.Context, obj *models.Topics) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Company returns generated.CompanyResolver implementation.
func (r *Resolver) Company() generated.CompanyResolver { return &companyResolver{r} }

// Media returns generated.MediaResolver implementation.
func (r *Resolver) Media() generated.MediaResolver { return &mediaResolver{r} }

// Topics returns generated.TopicsResolver implementation.
func (r *Resolver) Topics() generated.TopicsResolver { return &topicsResolver{r} }

type companyResolver struct{ *Resolver }
type mediaResolver struct{ *Resolver }
type topicsResolver struct{ *Resolver }
