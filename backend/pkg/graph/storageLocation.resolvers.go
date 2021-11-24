package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/pkg/model"
	"github.com/georgysavva/scany/pgxscan"
	errs "github.com/pkg/errors"
)

func (r *storageLocationResolver) Image(ctx context.Context, obj *model.StorageLocation) (*model.File, error) {
	return r.File(ctx, obj.ImageID)
}

func (r *storageLocationResolver) Category(ctx context.Context, obj *model.StorageLocation) (*model.StorageLocationCategory, error) {
	var storageLocationCategory model.StorageLocationCategory
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &storageLocationCategory, `select * from omegarogue.public.storage_location_category where id=?;`, obj.CategoryID)
	if err != nil {
		return nil, errs.Wrap(err, "get category")
	}
	return &storageLocationCategory, nil
}

func (r *storageLocationResolver) Parts(ctx context.Context, obj *model.StorageLocation) ([]*model.Part, error) {
	var parts []*model.Part
	// language=SQL
	err := pgxscan.Select(ctx, r.DbPool, &parts, `select * from omegarogue.public.part where category=?;`, obj.ID)
	if err != nil {
		return nil, errs.Wrap(err, "get parts")
	}
	return parts, nil
}

func (r *storageLocationCategoryResolver) StorageLocations(ctx context.Context, obj *model.StorageLocationCategory) ([]*model.StorageLocation, error) {
	var storageLocations []*model.StorageLocation
	// language=SQL
	err := pgxscan.Select(ctx, r.DbPool, &storageLocations, `select * from omegarogue.public.storage_location where category=?;`, obj.ID)
	if err != nil {
		return nil, errs.Wrap(err, "get storage locations")
	}
	return storageLocations, nil
}

func (r *storageLocationCategoryResolver) Parent(ctx context.Context, obj *model.StorageLocationCategory) (*model.StorageLocationCategory, error) {
	var parent model.StorageLocationCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &parent,
		`select * from omegarogue.public.storage_location_category where id::text = subpath(?, -1, 1)`, obj.Path,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get parent")
	}
	return &parent, nil
}

func (r *storageLocationCategoryResolver) Children(ctx context.Context, obj *model.StorageLocationCategory) ([]*model.StorageLocationCategory, error) {
	var children []*model.StorageLocationCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &children, `select * from omegarogue.public.storage_location_category where path ~ ? || '.*{1}'`,
		obj.Path,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get children")
	}
	return children, nil
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
