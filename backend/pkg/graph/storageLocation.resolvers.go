package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/pkg/model"
)

func (r *storageLocationResolver) Image(ctx context.Context, obj *model.StorageLocation) (
	*model.StorageLocationImage,
	error,
) {
	panic(fmt.Errorf("not implemented"))
}

func (r *storageLocationResolver) Category(
	ctx context.Context,
	obj *model.StorageLocation,
) (*model.StorageLocationCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *storageLocationResolver) Parts(ctx context.Context, obj *model.StorageLocation) ([]*model.Part, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *storageLocationCategoryResolver) StorageLocations(
	ctx context.Context,
	obj *model.StorageLocationCategory,
) ([]*model.StorageLocation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *storageLocationCategoryResolver) Parent(
	ctx context.Context,
	obj *model.StorageLocationCategory,
) (*model.StorageLocationCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *storageLocationCategoryResolver) Children(
	ctx context.Context,
	obj *model.StorageLocationCategory,
) ([]*model.StorageLocationCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

// StorageLocation returns generated.StorageLocationResolver implementation.
func (r *Resolver) StorageLocation() generated.StorageLocationResolver {
	return &storageLocationResolver{r}
}

// StorageLocationCategory returns generated.StorageLocationCategoryResolver implementation.
func (r *Resolver) StorageLocationCategory() generated.StorageLocationCategoryResolver {
	return &storageLocationCategoryResolver{r}
}

type storageLocationResolver struct{ *Resolver }
type storageLocationCategoryResolver struct{ *Resolver }
