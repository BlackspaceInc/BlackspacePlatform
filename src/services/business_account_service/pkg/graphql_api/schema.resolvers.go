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
	ID := obj.GetId()
	return HandleErrorsIfPresent(ID)
}

func (r *businessAccountResolver) MerchantType(ctx context.Context, obj *proto.BusinessAccount) (*models.MerchantType, error) {
	merchantKeyRef := obj.GetMerchantType()
	key := int32(merchantKeyRef)
	merchantType := models.MerchantType(proto.MerchantType_name[key])
	return &merchantType, nil
}

func (r *businessAccountResolver) ServicesManagedByBlackspace(ctx context.Context, obj *proto.BusinessAccount) (*models.ServicesManagedByBlackspace, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessAccountResolver) AuthnID(ctx context.Context, obj *proto.BusinessAccount) (*int, error) {
	ID := obj.GetAuthnId()
	return HandleErrorsIfPresent(ID)
}

func (r *businessTypeResolver) Category(ctx context.Context, obj *proto.BusinessType) (*models.BusinessCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *businessTypeResolver) SubCategory(ctx context.Context, obj *proto.BusinessType) (*models.BusinessSubCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mediaResolver) ID(ctx context.Context, obj *proto.Media) (*int, error) {
	ID := obj.GetId()
	return HandleErrorsIfPresent(ID)
}

func (r *paymentProcessingMethodsResolver) PaymentOptions(ctx context.Context, obj *proto.PaymentProcessingMethods) ([]*models.PaymentOptions, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *paymentProcessingMethodsResolver) Medium(ctx context.Context, obj *proto.PaymentProcessingMethods) ([]*models.PaymentMedium, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *phoneNumberResolver) Type(ctx context.Context, obj *proto.PhoneNumber) (*models.PhoneType, error) {
	phoneTypeKeyRef := obj.GetType()
	key := int32(phoneTypeKeyRef)
	phoneType := models.PhoneType(proto.PhoneType_name[key])
	// TODO handle
	return &phoneType, nil
}

func (r *topicsResolver) ID(ctx context.Context, obj *proto.Topics) (*int, error) {
	ID := obj.GetId()
	return HandleErrorsIfPresent(ID)
}

// BusinessAccount returns generated.BusinessAccountResolver implementation.
func (r *Resolver) BusinessAccount() generated.BusinessAccountResolver {
	return &businessAccountResolver{r}
}

// BusinessType returns generated.BusinessTypeResolver implementation.
func (r *Resolver) BusinessType() generated.BusinessTypeResolver { return &businessTypeResolver{r} }

// Media returns generated.MediaResolver implementation.
func (r *Resolver) Media() generated.MediaResolver { return &mediaResolver{r} }

// PaymentProcessingMethods returns generated.PaymentProcessingMethodsResolver implementation.
func (r *Resolver) PaymentProcessingMethods() generated.PaymentProcessingMethodsResolver {
	return &paymentProcessingMethodsResolver{r}
}

// PhoneNumber returns generated.PhoneNumberResolver implementation.
func (r *Resolver) PhoneNumber() generated.PhoneNumberResolver { return &phoneNumberResolver{r} }

// Topics returns generated.TopicsResolver implementation.
func (r *Resolver) Topics() generated.TopicsResolver { return &topicsResolver{r} }

type businessAccountResolver struct{ *Resolver }
type businessTypeResolver struct{ *Resolver }
type mediaResolver struct{ *Resolver }
type paymentProcessingMethodsResolver struct{ *Resolver }
type phoneNumberResolver struct{ *Resolver }
type topicsResolver struct{ *Resolver }
