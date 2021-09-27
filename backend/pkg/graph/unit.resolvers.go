package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/pkg/model"
)

func (r *unitResolver) Prefixes(ctx context.Context, obj *model.Unit) ([]*model.SiPrefix, error) {
	panic(fmt.Errorf("not implemented"))
}

// Unit returns generated.UnitResolver implementation.
func (r *Resolver) Unit() generated.UnitResolver { return &unitResolver{r} }

type unitResolver struct{ *Resolver }
