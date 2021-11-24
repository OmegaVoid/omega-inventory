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

func (r *partResolver) Category(ctx context.Context, obj *model.Part) (*model.PartCategory, error) {
	var category model.PartCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &category, `select * from omegarogue.public.part_category where id = ?`, obj.CategoryID,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get category")
	}
	return &category, nil
}

func (r *partResolver) Footprint(ctx context.Context, obj *model.Part) (*model.Footprint, error) {
	var footprint model.Footprint
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &footprint, `select * from omegarogue.public.footprint where id = ?`, obj.FootprintID,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get footprint")
	}
	return &footprint, nil
}

func (r *partResolver) Unit(ctx context.Context, obj *model.Part) (*model.PartMeasurementUnit, error) {
	var unit model.PartMeasurementUnit
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &unit, `select * from omegarogue.public.part_unit where id = ?`, obj.UnitID)
	if err != nil {
		return nil, errs.Wrapf(err, "get unit")
	}
	return &unit, nil
}

func (r *partResolver) StorageLocation(ctx context.Context, obj *model.Part) (*model.StorageLocation, error) {
	var storageLocation model.StorageLocation
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocation, `select * from omegarogue.public.storage_location where id = ?`,
		obj.StorageLocationID,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get storage location")
	}
	return &storageLocation, nil
}

func (r *partResolver) Attachments(ctx context.Context, obj *model.Part) ([]*model.File, error) {
	return r.Files(ctx, obj.AttachmentsID)
}

func (r *partResolver) Parameters(ctx context.Context, obj *model.Part) ([]*model.PartParameter, error) {
	var partParameters []*model.PartParameter
	// language=SQL
	err := pgxscan.Select(
		ctx, r.DbPool, &partParameters, `select * from omegarogue.public.part_parameter where id=any(?);`,
		obj.ParametersID,
	)
	if err != nil {
		return nil, errs.Wrap(err, "get part parameters")
	}
	return partParameters, nil
}

func (r *partCategoryResolver) Parts(ctx context.Context, obj *model.PartCategory) ([]*model.Part, error) {
	var parts []*model.Part
	// language=SQL
	err := pgxscan.Select(ctx, r.DbPool, &parts, `select * from omegarogue.public.part where category=?;`, obj.ID)
	if err != nil {
		return nil, errs.Wrap(err, "get parts")
	}
	return parts, nil
}

func (r *partCategoryResolver) Children(ctx context.Context, obj *model.PartCategory) ([]*model.PartCategory, error) {
	var children []*model.PartCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &children, `select * from omegarogue.public.part_category where path ~ ? || '.*{1}'`, obj.Path,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get parent")
	}
	return children, nil
}

func (r *partCategoryResolver) Parent(ctx context.Context, obj *model.PartCategory) (*model.PartCategory, error) {
	var parent model.PartCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &parent, `select * from omegarogue.public.part_category where id::text = subpath(?, -1, 1)`,
		obj.Path,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get parent")
	}
	return &parent, nil
}

func (r *partParameterResolver) Part(ctx context.Context, obj *model.PartParameter) (*model.Part, error) {
	var part model.Part
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &part, `select * from omegarogue.public.part where ? = any(parameters)`, obj.ID)
	if err != nil {
		return nil, errs.Wrapf(err, "get part")
	}
	return &part, nil
}

func (r *partParameterResolver) Unit(ctx context.Context, obj *model.PartParameter) (*model.Unit, error) {
	var unit model.Unit
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &unit, `select * from omegarogue.public.part_category where id = ?`, obj.UnitID)
	if err != nil {
		return nil, errs.Wrapf(err, "get unit")
	}
	return &unit, nil
}

func (r *partParameterResolver) SiPrefix(ctx context.Context, obj *model.PartParameter) (*model.SiPrefix, error) {
	var siPrefix model.SiPrefix
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &siPrefix, `select * from omegarogue.public.part_category where id = ?`, obj.SiPrefixID,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get si prefix")
	}
	return &siPrefix, nil
}

func (r *partParameterResolver) MinSiPrefix(ctx context.Context, obj *model.PartParameter) (*model.SiPrefix, error) {
	var siPrefix model.SiPrefix
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &siPrefix, `select * from omegarogue.public.part_category where id = ?`, obj.MinSiPrefixID,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get min si prefix")
	}
	return &siPrefix, nil
}

func (r *partParameterResolver) MaxSiPrefix(ctx context.Context, obj *model.PartParameter) (*model.SiPrefix, error) {
	var siPrefix model.SiPrefix
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &siPrefix, `select * from omegarogue.public.part_category where id = ?`, obj.MaxSiPrefixID,
	)
	if err != nil {
		return nil, errs.Wrapf(err, "get max si prefix")
	}
	return &siPrefix, nil
}

// Part returns generated.PartResolver implementation.
func (r *Resolver) Part() generated.PartResolver { return &partResolver{r} }

// PartCategory returns generated.PartCategoryResolver implementation.
func (r *Resolver) PartCategory() generated.PartCategoryResolver { return &partCategoryResolver{r} }

// PartParameter returns generated.PartParameterResolver implementation.
func (r *Resolver) PartParameter() generated.PartParameterResolver { return &partParameterResolver{r} }

type partResolver struct{ *Resolver }
type partCategoryResolver struct{ *Resolver }
type partParameterResolver struct{ *Resolver }
