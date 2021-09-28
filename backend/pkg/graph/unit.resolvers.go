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

func (r *unitResolver) Prefixes(ctx context.Context, obj *model.Unit) ([]*model.SiPrefix, error) {
	var siPrefixes []*model.SiPrefix
	// language=SQL
	err := pgxscan.Select(ctx, r.DbPool, &siPrefixes, `select * from omegarogue.public.si_prefix where id=any(?);`, obj.PrefixesID)
	if err != nil {
		return nil, errs.Wrap(err, "get prefixes")
	}
	return siPrefixes, nil
}

// Unit returns generated.UnitResolver implementation.
func (r *Resolver) Unit() generated.UnitResolver { return &unitResolver{r} }

type unitResolver struct{ *Resolver }
