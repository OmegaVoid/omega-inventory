package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/OmegaVoid/omega-inventory/pkg/model"
	"github.com/bwmarrin/snowflake"
	"github.com/georgysavva/scany/pgxscan"
	errs "github.com/pkg/errors"
)

func (r *Resolver) Files(ctx context.Context, ids ...[]*snowflake.ID) ([]*model.File, error) {
	var files []*model.File
	if len(ids) == 0 {
		// language=SQL
		err := pgxscan.Select(ctx, r.DbPool, &files, `select * from omegarogue.public.attachment;`)
		if err != nil {
			return nil, errs.Wrap(err, "get files")
		}
		return files, nil
	}

	files = make([]*model.File, len(ids))

	// language=SQL
	err := pgxscan.Select(ctx, r.DbPool, &files, `select * from omegarogue.public.attachment where id=any(?);`, ids)
	if err != nil {
		return nil, errs.Wrap(err, "get files")
	}
	return files, nil
}

func (r *Resolver) File(ctx context.Context, id *snowflake.ID) (*model.File, error) {
	var file model.File
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &file, `select * from omegarogue.public.attachment where id=?;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "get file")
	}
	return &file, nil
}
