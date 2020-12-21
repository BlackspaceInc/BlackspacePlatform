package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql/generated"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/grpc/proto"
)

func (r *businessAccountResolver) ID(ctx context.Context, obj *proto.BusinessAccount) (*int, error) {
	id := int(obj.Id)
	return handleErrorIfPresent(id)}

func (r *businessAccountResolver) AuthnID(ctx context.Context, obj *proto.BusinessAccount) (*int, error) {
	id := int(obj.AuthnId)
	return handleErrorIfPresent(id)
}

func (r *mediaResolver) ID(ctx context.Context, obj *proto.Media) (*int, error) {
	id := int(obj.Id)
	return handleErrorIfPresent(id)
}

func (r *topicsResolver) ID(ctx context.Context, obj *proto.Topics) (*int, error) {
	id := int(obj.Id)
	return handleErrorIfPresent(id)
}

// BusinessAccount returns generated.BusinessAccountResolver implementation.
func (r *Resolver) BusinessAccount() generated.BusinessAccountResolver {
	return &businessAccountResolver{r}
}

// Media returns generated.MediaResolver implementation.
func (r *Resolver) Media() generated.MediaResolver { return &mediaResolver{r} }

// Topics returns generated.TopicsResolver implementation.
func (r *Resolver) Topics() generated.TopicsResolver { return &topicsResolver{r} }

type businessAccountResolver struct{ *Resolver }
type mediaResolver struct{ *Resolver }
type topicsResolver struct{ *Resolver }
