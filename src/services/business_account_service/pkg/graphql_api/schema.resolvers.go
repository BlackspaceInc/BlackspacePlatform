package graphql_api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api/generated"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api/models"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api/proto"
)

func (r *businessAccountResolver) ID(ctx context.Context, obj *proto.BusinessAccount) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) PhoneNumber(ctx context.Context, obj *proto.BusinessAccount) (*models.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) Media(ctx context.Context, obj *proto.BusinessAccount) (*models.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) TypeOfBusiness(ctx context.Context, obj *proto.BusinessAccount) (*models.BusinessType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) MerchantType(ctx context.Context, obj *proto.BusinessAccount) (*models.MerchantType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) PaymentDetails(ctx context.Context, obj *proto.BusinessAccount) (*models.PaymentProcessingMethods, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) ServicesManagedByBlackspace(ctx context.Context, obj *proto.BusinessAccount) (*models.ServicesManagedByBlackspace, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) FounderAddress(ctx context.Context, obj *proto.BusinessAccount) (*models.Address, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) SubscribedTopics(ctx context.Context, obj *proto.BusinessAccount) (*models.Topics, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) AuthnID(ctx context.Context, obj *proto.BusinessAccount) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// BusinessAccount returns generated.BusinessAccountResolver implementation.
func (r *Resolver) BusinessAccount() generated.BusinessAccountResolver {
	return &businessAccountResolver{r}
}

type businessAccountResolver struct{ *Resolver }
