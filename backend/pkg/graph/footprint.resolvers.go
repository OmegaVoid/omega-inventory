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

func (r *footprintResolver) Category(ctx context.Context, obj *model.Footprint) (*model.FootprintCategory, error) {
	var footPrintCategory model.FootprintCategory
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &footPrintCategory, `select * from omegarogue.public.footprint_category where id=?;`, obj.CategoryID)
	if err != nil {
		return nil, errs.Wrap(err, "get category")
	}
	return &footPrintCategory, nil
}

func (r *footprintResolver) Attachments(ctx context.Context, obj *model.Footprint) ([]*model.File, error) {
	attachments := make([]*model.File, len(obj.AttachmentsID))
	for i, id := range obj.AttachmentsID {
		attachments[i] = &model.File{
			ID: *id, // TODO add other props
		}
	}
	return attachments, nil
}

func (r *footprintResolver) Image(ctx context.Context, obj *model.Footprint) (*model.File, error) {
	return &model.File{
		ID: *obj.ImageID, // TODO
	}, nil
}

func (r *footprintCategoryResolver) Parent(ctx context.Context, obj *model.FootprintCategory) (*model.FootprintCategory, error) {
	var parent model.FootprintCategory
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &parent, `select * from omegarogue.public.footprint_category where path = ?`, obj.ParentID)
	if err != nil {
		return nil, errs.Wrapf(err, "get parent")
	}
	return &parent, nil
}

func (r *footprintCategoryResolver) Children(ctx context.Context, obj *model.FootprintCategory) ([]*model.FootprintCategory, error) {
	var children []*model.FootprintCategory
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &children, `select * from omegarogue.public.footprint_category where path ~ '*.' || ? || '.*{1}'`, obj.ID)
	if err != nil {
		return nil, errs.Wrapf(err, "get parent")
	}
	return children, nil
}

func (r *footprintCategoryResolver) Footprints(ctx context.Context, obj *model.FootprintCategory) ([]*model.Footprint, error) {
	var footPrints []*model.Footprint
	// language=SQL
	err := pgxscan.Select(ctx, r.DbPool, &footPrints, `select * from omegarogue.public.footprint where category=?;`, obj.ID)
	if err != nil {
		return nil, errs.Wrap(err, "get footprints")
	}
	return footPrints, nil
}

// Footprint returns generated.FootprintResolver implementation.
func (r *Resolver) Footprint() generated.FootprintResolver { return &footprintResolver{r} }

// FootprintCategory returns generated.FootprintCategoryResolver implementation.
func (r *Resolver) FootprintCategory() generated.FootprintCategoryResolver {
	return &footprintCategoryResolver{r}
}

type footprintResolver struct{ *Resolver }
type footprintCategoryResolver struct{ *Resolver }
