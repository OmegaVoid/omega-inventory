package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/pkg/middlewares"
	"github.com/OmegaVoid/omega-inventory/pkg/model"
	"github.com/bwmarrin/snowflake"
	errs "github.com/pkg/errors"
)

func (r *mutationResolver) CreatePartMeasurementUnit(
	ctx context.Context,
	input model.PartMeasurementUnitInput,
) (*model.PartMeasurementUnit, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errs.Wrap(err, "new snowflake node")
	}
	unit := model.PartMeasurementUnit{
		ID:        node.Generate().String(),
		Name:      input.Name,
		ShortName: input.ShortName,
	}
	r.measurementUnits = append(r.measurementUnits, &unit)
	return &unit, nil
}

func (r *mutationResolver) UpdatePartMeasurementUnit(
	ctx context.Context,
	id string,
	input model.PartMeasurementUnitInput,
) (*model.PartMeasurementUnit, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeletePartMeasurementUnit(ctx context.Context, id string) (
	*model.PartMeasurementUnit,
	error,
) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreatePartCategory(ctx context.Context, input model.PartCategoryInput) (
	*model.PartCategory,
	error,
) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errs.Wrap(err, "new snowflake node")
	}
	category := model.PartCategory{
		ID:          node.Generate().String(),
		Name:        input.Name,
		Description: input.Description,
	}
	r.partCategories = append(r.partCategories, &category)
	return &category, nil
}

func (r *mutationResolver) UpdatePartCategory(
	ctx context.Context,
	id string,
	input model.PartCategoryInput,
) (*model.PartCategory, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeletePartCategory(ctx context.Context, id string) (*model.PartCategory, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreateStorageLocationCategory(
	ctx context.Context,
	input model.StorageLocationCategoryInput,
) (*model.StorageLocationCategory, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errs.Wrap(err, "new snowflake node")
	}
	storageLocationCategory := model.StorageLocationCategory{
		ID:          node.Generate().String(),
		Name:        input.Name,
		Description: input.Description,
	}
	r.storageLocationCategories = append(r.storageLocationCategories, &storageLocationCategory)
	return &storageLocationCategory, nil
}

func (r *mutationResolver) UpdateStorageLocationCategory(
	ctx context.Context,
	id string,
	input model.StorageLocationCategoryInput,
) (*model.StorageLocationCategory, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeleteStorageLocationCategory(
	ctx context.Context,
	id string,
) (*model.StorageLocationCategory, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreateStorageLocation(
	ctx context.Context,
	input model.StorageLocationInput,
) (*model.StorageLocation, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errs.Wrap(err, "new snowflake node")
	}
	storageLocation := model.StorageLocation{
		ID:   node.Generate().String(),
		Name: input.Name,
	}
	r.storageLocations = append(r.storageLocations, &storageLocation)
	return &storageLocation, nil
}

func (r *mutationResolver) UpdateStorageLocation(
	ctx context.Context,
	id string,
	input model.StorageLocationInput,
) (*model.StorageLocation, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeleteStorageLocation(ctx context.Context, id string) (*model.StorageLocation, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreatePart(ctx context.Context, input model.PartInput) (*model.Part, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errs.Wrap(err, "new snowflake node")
	}
	part := model.Part{
		ID:                 node.Generate().String(),
		Name:               input.Name,
		Description:        input.Description,
		Comment:            input.Comment,
		StockLevel:         input.StockLevel,
		MinStockLevel:      input.MinStockLevel,
		InternalPartNumber: input.InternalPartNumber,
	}
	r.parts = append(r.parts, &part)
	return &part, nil
}

func (r *mutationResolver) UpdatePart(ctx context.Context, id string, input model.PartInput) (*model.Part, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeletePart(ctx context.Context, id string) (*model.Part, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreateSiPrefix(ctx context.Context, input model.SiPrefixInput) (*model.SiPrefix, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) UpdateSiPrefix(ctx context.Context, id string, input model.SiPrefixInput) (
	*model.SiPrefix,
	error,
) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeleteSiPrefix(ctx context.Context, id string) (*model.SiPrefix, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreateUnit(ctx context.Context, input model.UnitInput) (*model.Unit, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) UpdateUnit(ctx context.Context, id string, input model.UnitInput) (*model.Unit, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeleteUnit(ctx context.Context, id string) (*model.Unit, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreatePartAttachment(
	ctx context.Context,
	input model.PartAttachmentInput,
) (*model.PartAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) UpdatePartAttachment(
	ctx context.Context,
	id string,
	input model.PartAttachmentInput,
) (*model.PartAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeletePartAttachment(ctx context.Context, id string) (*model.PartAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreateFootprint(ctx context.Context, input model.FootprintInput) (*model.Footprint, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) UpdateFootprint(
	ctx context.Context,
	id string,
	input model.FootprintInput,
) (*model.Footprint, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeleteFootprint(ctx context.Context, id string) (*model.Footprint, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreateFootprintCategory(
	ctx context.Context,
	input model.FootprintCategoryInput,
) (*model.FootprintCategory, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) UpdateFootprintCategory(
	ctx context.Context,
	id string,
	input model.FootprintCategoryInput,
) (*model.FootprintCategory, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeleteFootprintCategory(ctx context.Context, id string) (*model.FootprintCategory, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreateFootprintAttachment(
	ctx context.Context,
	input model.FootprintAttachmentInput,
) (*model.FootprintAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) UpdateFootprintAttachment(
	ctx context.Context,
	id string,
	input model.FootprintAttachmentInput,
) (*model.FootprintAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeleteFootprintAttachment(ctx context.Context, id string) (
	*model.FootprintAttachment,
	error,
) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) CreatePartParameter(
	ctx context.Context,
	input model.PartParameterInput,
) (*model.PartParameter, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) UpdatePartParameter(
	ctx context.Context,
	id string,
	input model.PartParameterInput,
) (*model.PartParameter, error) {
	return nil, errs.New("not implemented")
}

func (r *mutationResolver) DeletePartParameter(ctx context.Context, id string) (*model.PartParameter, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) Parts(ctx context.Context, category *model.PartCategoryInput) ([]*model.Part, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return r.parts, nil
}

func (r *queryResolver) Part(ctx context.Context, id string) (*model.Part, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) PartCategories(ctx context.Context, category *model.PartCategoryInput) (
	[]*model.PartCategory,
	error,
) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return r.partCategories, nil
}

func (r *queryResolver) PartCategory(ctx context.Context, id string) (*model.PartCategory, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) StorageLocations(
	ctx context.Context,
	category *model.StorageLocationCategoryInput,
) ([]*model.StorageLocation, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return r.storageLocations, nil
}

func (r *queryResolver) StorageLocation(ctx context.Context, id string) (*model.StorageLocation, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) StorageLocationCategories(
	ctx context.Context,
	category *model.StorageLocationCategoryInput,
) ([]*model.StorageLocationCategory, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return r.storageLocationCategories, nil
}

func (r *queryResolver) StorageLocationCategory(ctx context.Context, id string) (
	*model.StorageLocationCategory,
	error,
) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) PartMeasurementUnits(ctx context.Context) ([]*model.PartMeasurementUnit, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "get gin context")
	}
	_ = gc
	return r.measurementUnits, nil
}

func (r *queryResolver) PartMeasurementUnit(ctx context.Context, id string) (*model.PartMeasurementUnit, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) SiPrefixes(ctx context.Context) ([]*model.SiPrefix, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) SiPrefix(ctx context.Context, id string) (*model.SiPrefix, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) Units(ctx context.Context) ([]*model.Unit, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) Unit(ctx context.Context, id string) (*model.Unit, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) PartAttachments(ctx context.Context) ([]*model.PartAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) PartAttachment(ctx context.Context, id string) (*model.PartAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) Footprints(ctx context.Context) ([]*model.Footprint, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) Footprint(ctx context.Context, id string) (*model.Footprint, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) FootprintCategories(ctx context.Context) ([]*model.FootprintCategory, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) FootprintCategory(ctx context.Context, id string) (*model.FootprintCategory, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) FootprintAttachments(ctx context.Context) ([]*model.FootprintAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) FootprintAttachment(ctx context.Context, id string) (*model.FootprintAttachment, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) PartParameters(ctx context.Context) ([]*model.PartParameter, error) {
	return nil, errs.New("not implemented")
}

func (r *queryResolver) PartParameter(ctx context.Context, id string) (*model.PartParameter, error) {
	return nil, errs.New("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
