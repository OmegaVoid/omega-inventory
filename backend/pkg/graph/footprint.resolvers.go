package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/pkg/model"
)

func (r *footprintResolver) Category(ctx context.Context, obj *model.Footprint) (*model.FootprintCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *footprintResolver) Attachments(ctx context.Context, obj *model.Footprint) (
	[]*model.FootprintAttachment,
	error,
) {
	panic(fmt.Errorf("not implemented"))
}

func (r *footprintResolver) Image(ctx context.Context, obj *model.Footprint) (*model.FootprintAttachment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *footprintCategoryResolver) Parent(ctx context.Context, obj *model.FootprintCategory) (
	*model.FootprintCategory,
	error,
) {
	panic(fmt.Errorf("not implemented"))
}

func (r *footprintCategoryResolver) Children(
	ctx context.Context,
	obj *model.FootprintCategory,
) ([]*model.FootprintCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *footprintCategoryResolver) Footprints(ctx context.Context, obj *model.FootprintCategory) (
	[]*model.Footprint,
	error,
) {
	panic(fmt.Errorf("not implemented"))
}

// Footprint returns generated.FootprintResolver implementation.
func (r *Resolver) Footprint() generated.FootprintResolver { return &footprintResolver{r} }

// FootprintCategory returns generated.FootprintCategoryResolver implementation.
func (r *Resolver) FootprintCategory() generated.FootprintCategoryResolver {
	return &footprintCategoryResolver{r}
}

type footprintResolver struct{ *Resolver }
type footprintCategoryResolver struct{ *Resolver }
