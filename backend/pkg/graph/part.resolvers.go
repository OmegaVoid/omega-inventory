package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/pkg/model"
)

func (r *partResolver) Category(ctx context.Context, obj *model.Part) (*model.PartCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partResolver) Footprint(ctx context.Context, obj *model.Part) (*model.Footprint, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partResolver) Unit(ctx context.Context, obj *model.Part) (*model.PartMeasurementUnit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partResolver) StorageLocation(ctx context.Context, obj *model.Part) (*model.StorageLocation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partResolver) Attachments(ctx context.Context, obj *model.Part) ([]*model.PartAttachment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partResolver) Parameters(ctx context.Context, obj *model.Part) ([]*model.PartParameter, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partCategoryResolver) Parts(ctx context.Context, obj *model.PartCategory) ([]*model.PartCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partCategoryResolver) Children(ctx context.Context, obj *model.PartCategory) ([]*model.PartCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partCategoryResolver) Parent(ctx context.Context, obj *model.PartCategory) (*model.PartCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partParameterResolver) Part(ctx context.Context, obj *model.PartParameter) (*model.Part, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partParameterResolver) Unit(ctx context.Context, obj *model.PartParameter) (*model.Unit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partParameterResolver) ValueType(ctx context.Context, obj *model.PartParameter) (model.ValueType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partParameterResolver) SiPrefix(ctx context.Context, obj *model.PartParameter) (*model.SiPrefix, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partParameterResolver) MinSiPrefix(ctx context.Context, obj *model.PartParameter) (*model.SiPrefix, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *partParameterResolver) MaxSiPrefix(ctx context.Context, obj *model.PartParameter) (*model.SiPrefix, error) {
	panic(fmt.Errorf("not implemented"))
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
