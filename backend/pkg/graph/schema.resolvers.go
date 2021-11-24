package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/pkg/model"
	"github.com/bwmarrin/snowflake"
	"github.com/georgysavva/scany/pgxscan"
	errs "github.com/pkg/errors"
)

func (r *mutationResolver) CreateFootprint(ctx context.Context, input model.FootprintInput) (*model.Footprint, error) {
	var footprint model.Footprint

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &footprint,
		`insert into footprint (fname, category, description, attachments, image) values (?, ?, ?, ?, ?) returning *;`,
		input.Name, input.Category, input.Description, input.Attachments, input.Image,
	)
	if err != nil {
		return nil, errs.Wrap(err, "create footprint")
	}
	return &footprint, nil
}

func (r *mutationResolver) UpdateFootprint(ctx context.Context, id snowflake.ID, input model.FootprintInput) (*model.Footprint, error) {
	var footprint model.Footprint

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &footprint,
		`update footprint set (fname, category, description, attachments, image) = (?, ?, ?, ?, ?) where id = ? returning *;`,
		input.Name, input.Category, input.Description, input.Attachments, input.Image, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update footprint")
	}
	return &footprint, nil
}

func (r *mutationResolver) DeleteFootprint(ctx context.Context, id snowflake.ID) (*model.Footprint, error) {
	var footprint model.Footprint
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &footprint, `delete from footprint where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete footprint")
	}
	return &footprint, nil
}

func (r *mutationResolver) CreateFootprintCategory(ctx context.Context, input model.FootprintCategoryInput) (*model.FootprintCategory, error) {
	var footprintCategory model.FootprintCategory

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &footprintCategory,
		`insert into footprint_category (cname, description, path) values (?, ?, ?) returning *;`,
		input.Name, input.Description, input.Parent,
	)
	if err != nil {
		return nil, errs.Wrap(err, "create footprint_category")
	}
	return &footprintCategory, nil
}
func (r *mutationResolver) UpdateFootprintCategory(ctx context.Context, id snowflake.ID, input model.FootprintCategoryInput) (*model.FootprintCategory, error) {
	var footprintCategory model.FootprintCategory

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &footprintCategory,
		`update footprint_category set (cname, description) = (?, ?) where id = ? returning *;`,
		input.Name, input.Description, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}

	r.DbPool.Exec(
		ctx,
		// language=SQL
		`
update footprint_category
set path = ( select p || $1::text
             from ( select path as p from storage_location_category where cname = ltree2text($1::text)::varchar ) T ) ||
           subpath(path, nlevel($2) - 1)
where path <@ $2;`,
		footprintCategory.ID, footprintCategory.Path,
	)

	return &footprintCategory, nil
}

func (r *mutationResolver) DeleteFootprintCategory(ctx context.Context, id snowflake.ID) (*model.FootprintCategory, error) {
	var footprintCategory model.FootprintCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &footprintCategory, `delete from footprint_category where id = ? returning *;`, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "delete footprint category")
	}
	return &footprintCategory, nil
}

func (r *mutationResolver) CreatePartMeasurementUnit(ctx context.Context, input model.PartMeasurementUnitInput) (*model.PartMeasurementUnit, error) {
	var unit model.PartMeasurementUnit

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &unit, `insert into part_unit (uname, short_name, is_default) values (?, ?, ?) returning *;`,
		input.Name, input.ShortName, false,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &unit, nil
}

func (r *mutationResolver) UpdatePartMeasurementUnit(ctx context.Context, id snowflake.ID, input model.PartMeasurementUnitInput) (*model.PartMeasurementUnit, error) {
	var unit model.PartMeasurementUnit

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &unit, `update part_unit set (uname, short_name) = (?, ?) where id = ? returning *;`,
		input.Name, input.ShortName,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &unit, nil
}

func (r *mutationResolver) DeletePartMeasurementUnit(ctx context.Context, id snowflake.ID) (*model.PartMeasurementUnit, error) {
	var partMeasurementUnit model.PartMeasurementUnit
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partMeasurementUnit, `delete from part_unit where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete part measurement unit")
	}
	return &partMeasurementUnit, nil
}

func (r *mutationResolver) CreatePartCategory(ctx context.Context, input model.PartCategoryInput) (*model.PartCategory, error) {
	var partCategory model.PartCategory

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &partCategory,
		`insert into part_category (cname, description, path) values (?, ?, ?) returning *;`,
		input.Name, input.Description, input.Prefixes,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &partCategory, nil
}

func (r *mutationResolver) UpdatePartCategory(ctx context.Context, id snowflake.ID, input model.PartCategoryInput) (*model.PartCategory, error) {
	var partCategory model.PartCategory

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &partCategory, `update unit set (uname, symbol, prefixes) = (?, ?, ?) where id = ? returning *;`,
		input.Name, input.Description, input.Prefixes, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &partCategory, nil
}

func (r *mutationResolver) DeletePartCategory(ctx context.Context, id snowflake.ID) (*model.PartCategory, error) {
	var partCategory model.PartCategory
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partCategory, `delete from part_category where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete part category")
	}
	return &partCategory, nil
}

func (r *mutationResolver) CreatePart(ctx context.Context, input model.PartInput) (*model.Part, error) {
	var part model.Part

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &part, `insert into unit (uname, symbol, prefixes) values (?, ?, ?) returning *;`,
		input.Name, input.Symbol, input.Prefixes,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &part, nil
}

func (r *mutationResolver) UpdatePart(ctx context.Context, id snowflake.ID, input model.PartInput) (*model.Part, error) {
	var part model.Part

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &part, `update unit set (uname, symbol, prefixes) = (?, ?, ?) where id = ? returning *;`,
		input.Name, input.Symbol, input.Prefixes, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &part, nil
}

func (r *mutationResolver) DeletePart(ctx context.Context, id snowflake.ID) (*model.Part, error) {
	var part model.Part
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &part, `delete from part where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete part")
	}
	return &part, nil
}

func (r *mutationResolver) CreatePartParameter(ctx context.Context, input model.PartParameterInput) (*model.PartParameter, error) {
	var partParameter model.PartParameter

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &partParameter, `insert into unit (uname, symbol, prefixes) values (?, ?, ?) returning *;`,
		input.Name, input.Symbol, input.Prefixes,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &partParameter, nil
}

func (r *mutationResolver) UpdatePartParameter(ctx context.Context, id snowflake.ID, input model.PartParameterInput) (*model.PartParameter, error) {
	var partParameter model.PartParameter

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &partParameter,
		`update unit set (uname, symbol, prefixes) = (?, ?, ?) where id = ? returning *;`,
		input.Name, input.Symbol, input.Prefixes, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &partParameter, nil
}

func (r *mutationResolver) DeletePartParameter(ctx context.Context, id snowflake.ID) (*model.PartParameter, error) {
	var partParameter model.PartParameter
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partParameter, `delete from part parameter where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete part parameter")
	}
	return &partParameter, nil
}

func (r *mutationResolver) CreateStorageLocationCategory(ctx context.Context, input model.StorageLocationCategoryInput) (*model.StorageLocationCategory, error) {
	var storageLocationCategory model.StorageLocationCategory

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocationCategory,
		`insert into unit (uname, symbol, prefixes) values (?, ?, ?) returning *;`,
		input.Name, input.Symbol, input.Prefixes,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &storageLocationCategory, nil
}

func (r *mutationResolver) UpdateStorageLocationCategory(ctx context.Context, id snowflake.ID, input model.StorageLocationCategoryInput) (*model.StorageLocationCategory, error) {
	var storageLocationCategory model.StorageLocationCategory

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocationCategory,
		`update unit set (uname, symbol, prefixes) = (?, ?, ?) where id = ? returning *;`,
		input.Name, input.Symbol, input.Prefixes, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &storageLocationCategory, nil
}

func (r *mutationResolver) DeleteStorageLocationCategory(ctx context.Context, id snowflake.ID) (*model.StorageLocationCategory, error) {
	var storageLocationCategory model.StorageLocationCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocationCategory, `delete from storage_location_category where id = ? returning *;`, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "delete storage location category")
	}
	return &storageLocationCategory, nil
}

func (r *mutationResolver) CreateStorageLocation(ctx context.Context, input model.StorageLocationInput) (*model.StorageLocation, error) {
	var storageLocation model.StorageLocation

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocation, `insert into unit (uname, symbol, prefixes) values (?, ?, ?) returning *;`,
		input.Name, input.Symbol, input.Prefixes,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &storageLocation, nil
}

func (r *mutationResolver) UpdateStorageLocation(ctx context.Context, id snowflake.ID, input model.StorageLocationInput) (*model.StorageLocation, error) {
	var storageLocation model.StorageLocation

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocation,
		`update unit set (uname, symbol, prefixes) = (?, ?, ?) where id = ? returning *;`,
		input.Name, input.Symbol, input.Prefixes, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &storageLocation, nil
}

func (r *mutationResolver) DeleteStorageLocation(ctx context.Context, id snowflake.ID) (*model.StorageLocation, error) {
	var storageLocation model.StorageLocation
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &storageLocation, `delete from storage_location where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete storage location")
	}
	return &storageLocation, nil
}

func (r *mutationResolver) CreateSiPrefix(ctx context.Context, input model.SiPrefixInput) (*model.SiPrefix, error) {
	var siPrefix model.SiPrefix

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &siPrefix, `insert into unit (uname, symbol, prefixes) values (?, ?, ?) returning *;`,
		input.Name, input.Symbol, input.Prefixes,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &siPrefix, nil
}

func (r *mutationResolver) UpdateSiPrefix(ctx context.Context, id snowflake.ID, input model.SiPrefixInput) (*model.SiPrefix, error) {
	var siPrefix model.SiPrefix

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &siPrefix, `update unit set (uname, symbol, prefixes) = (?, ?, ?) where id = ? returning *;`,
		input.Name, input.Symbol, input.Prefixes, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &siPrefix, nil
}

func (r *mutationResolver) DeleteSiPrefix(ctx context.Context, id snowflake.ID) (*model.SiPrefix, error) {
	var siPrefix model.SiPrefix
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &siPrefix, `delete from si_prefix where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete si prefix")
	}
	return &siPrefix, nil
}

func (r *mutationResolver) CreateUnit(ctx context.Context, input model.UnitInput) (*model.Unit, error) {
	var unit model.Unit

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &unit, `insert into unit (uname, symbol, prefixes) values (?, ?, ?) returning *;`,
		input.Name, input.Symbol, input.Prefixes,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &unit, nil
}

func (r *mutationResolver) UpdateUnit(ctx context.Context, id snowflake.ID, input model.UnitInput) (*model.Unit, error) {
	var unit model.Unit

	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &unit, `update unit set (uname, symbol, prefixes) = (?, ?, ?) where id = ? returning *;`,
		input.Name, input.Symbol, input.Prefixes, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "update file")
	}
	return &unit, nil
}

func (r *mutationResolver) DeleteUnit(ctx context.Context, id snowflake.ID) (*model.Unit, error) {
	var unit model.Unit
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &unit, `delete from unit where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete unit")
	}
	return &unit, nil
}

func (r *mutationResolver) CreateFile(ctx context.Context, input model.FileInput) (*model.File, error) {
	var file model.File
	var snow snowflake.ID
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &snow, `select id_generator()`)
	if err != nil {
		return nil, errs.Wrap(err, "get snowflake")
	}

	isImage := strings.HasPrefix(input.Content.ContentType, "image")
	extension := filepath.Ext(input.Content.Filename)
	name := strings.TrimSuffix(input.Content.Filename, extension)

	err = r.Fs.WriteReader(fmt.Sprintf("%v.%s", snow, extension), input.Content.File)
	if err != nil {
		return nil, errs.Wrap(err, "create file")
	}

	if input.Description != nil {
		// language=SQL
		err = pgxscan.Get(
			ctx, r.DbPool, &file, `insert into attachment values (?, ?, ?, ?, ?, ?, ?, ?) returning *;`,
			snow, snow, name, input.Content.ContentType, input.Description, isImage, extension, input.Content.Size,
		)
		if err != nil {
			return nil, errs.Wrap(err, "create file")
		}
	} else {
		// language=SQL
		err = pgxscan.Get(
			ctx, r.DbPool, &file,
			`insert into attachment (id, filename, original_name, mimetype, is_image, extension, size) values (?, ?, ?, ?, ?, ?) returning *;`,
			snow, snow, name, input.Content.ContentType, isImage, extension, input.Content.Size,
		)
		if err != nil {
			return nil, errs.Wrap(err, "create file")
		}
	}
	return &file, nil
}

func (r *mutationResolver) UpdateFile(ctx context.Context, id snowflake.ID, input model.FileInput) (
	*model.File, error,
) {
	var file model.File

	if input.Content != nil {
		isImage := strings.HasPrefix(input.Content.ContentType, "image")
		extension := filepath.Ext(input.Content.Filename)
		name := strings.TrimSuffix(input.Content.Filename, extension)

		err := r.Fs.WriteReader(fmt.Sprintf("%v.%s", id, extension), input.Content.File)
		if err != nil {
			return nil, errs.Wrap(err, "update file")
		}

		if input.Description != nil {
			// language=SQL
			err = pgxscan.Get(
				ctx, r.DbPool, &file,
				`update attachment set (extension, description, is_image, mimetype, original_name, size) = (?, ?, ?, ?, ?, ?) where id = ? returning *;`,
				extension, input.Description, isImage, input.Content.ContentType, name, input.Content.Size, id,
			)
			if err != nil {
				return nil, errs.Wrap(err, "update file")
			}
		} else {
			// language=SQL
			err = pgxscan.Get(
				ctx, r.DbPool, &file,
				`update attachment set (extension, is_image, mimetype, original_name, size) = (?, ?, ?, ?, ?) where id = ? returning *;`,
				extension, isImage, input.Content.ContentType, name, input.Content.Size, id,
			)
			if err != nil {
				return nil, errs.Wrap(err, "update file")
			}
		}

	} else if input.Description != nil {
		// language=SQL
		err := pgxscan.Get(
			ctx, r.DbPool, &file, `update attachment set description = ? where id = ? returning *;`, input.Description,
			id,
		)
		if err != nil {
			return nil, errs.Wrap(err, "update file")
		}
	}

	return &file, nil
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id snowflake.ID) (*model.File, error) {
	var file model.File
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &file, `delete from attachment where id = ? returning *;`, id)
	if err != nil {
		return nil, errs.Wrap(err, "delete file")
	}
	return &file, nil
}

func (r *queryResolver) Footprints(ctx context.Context) ([]*model.Footprint, error) {
	var footprints []*model.Footprint
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &footprints, `select * from omegarogue.public.footprint`)
	if err != nil {
		return nil, errs.Wrap(err, "get footprints")
	}
	return footprints, nil
}

func (r *queryResolver) Footprint(ctx context.Context, id snowflake.ID) (*model.Footprint, error) {
	var footprint model.Footprint
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &footprint, `select * from omegarogue.public.footprint where id=?`, id)
	if err != nil {
		return nil, errs.Wrap(err, "get footprint")
	}
	return &footprint, nil
}

func (r *queryResolver) FootprintCategories(ctx context.Context) ([]*model.FootprintCategory, error) {
	var footprintCategories []*model.FootprintCategory
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &footprintCategories, `select * from omegarogue.public.footprint`)
	if err != nil {
		return nil, errs.Wrap(err, "get footprint categories")
	}
	return footprintCategories, nil
}

func (r *queryResolver) FootprintCategory(ctx context.Context, id snowflake.ID) (*model.FootprintCategory, error) {
	var footprintCategory model.FootprintCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &footprintCategory, `select * from omegarogue.public.footprint_category where id=?`, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "get footprint category")
	}
	return &footprintCategory, nil
}

func (r *queryResolver) Parts(ctx context.Context, category *model.PartCategoryInput) ([]*model.Part, error) {
	var parts []*model.Part
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &parts, `select * from omegarogue.public.part`)
	if err != nil {
		return nil, errs.Wrap(err, "get parts")
	}
	return parts, nil
}

func (r *queryResolver) Part(ctx context.Context, id snowflake.ID) (*model.Part, error) {
	var part model.Part
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &part, `select * from omegarogue.public.part where id=?`, id)
	if err != nil {
		return nil, errs.Wrap(err, "get part")
	}
	return &part, nil
}

func (r *queryResolver) PartCategories(ctx context.Context, category *model.PartCategoryInput) ([]*model.PartCategory, error) {
	var partCategories []*model.PartCategory
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partCategories, `select * from omegarogue.public.part_category`)
	if err != nil {
		return nil, errs.Wrap(err, "get part categories")
	}
	return partCategories, nil
}

func (r *queryResolver) PartCategory(ctx context.Context, id snowflake.ID) (*model.PartCategory, error) {
	var partCategory model.PartCategory
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partCategory, `select * from omegarogue.public.part_category where id=?`, id)
	if err != nil {
		return nil, errs.Wrap(err, "get part category")
	}
	return &partCategory, nil
}

func (r *queryResolver) PartMeasurementUnits(ctx context.Context) ([]*model.PartMeasurementUnit, error) {
	var partMeasurementUnits []*model.PartMeasurementUnit
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partMeasurementUnits, `select * from omegarogue.public.part_unit`)
	if err != nil {
		return nil, errs.Wrap(err, "get part units")
	}
	return partMeasurementUnits, nil
}

func (r *queryResolver) PartMeasurementUnit(ctx context.Context, id snowflake.ID) (*model.PartMeasurementUnit, error) {
	var partMeasurementUnit model.PartMeasurementUnit
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partMeasurementUnit, `select * from omegarogue.public.part_unit where id=?`, id)
	if err != nil {
		return nil, errs.Wrap(err, "get part unit")
	}
	return &partMeasurementUnit, nil
}

func (r *queryResolver) PartParameters(ctx context.Context) ([]*model.PartParameter, error) {
	var partParameters []*model.PartParameter
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partParameters, `select * from omegarogue.public.part_parameter`)
	if err != nil {
		return nil, errs.Wrap(err, "get part parameters")
	}
	return partParameters, nil
}

func (r *queryResolver) PartParameter(ctx context.Context, id snowflake.ID) (*model.PartParameter, error) {
	var partParameter model.PartParameter
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &partParameter, `select * from omegarogue.public.part_parameter where id=?`, id)
	if err != nil {
		return nil, errs.Wrap(err, "get part parameter")
	}
	return &partParameter, nil
}

func (r *queryResolver) StorageLocations(ctx context.Context, category *model.StorageLocationCategoryInput) ([]*model.StorageLocation, error) {
	var storageLocations []*model.StorageLocation
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &storageLocations, `select * from omegarogue.public.storage_location`)
	if err != nil {
		return nil, errs.Wrap(err, "get storage locations")
	}
	return storageLocations, nil
}

func (r *queryResolver) StorageLocation(ctx context.Context, id snowflake.ID) (*model.StorageLocation, error) {
	var storageLocation model.StorageLocation
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocation, `select * from omegarogue.public.storage_location where id=?`, id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "get storage location")
	}
	return &storageLocation, nil
}

func (r *queryResolver) StorageLocationCategories(ctx context.Context, category *model.StorageLocationCategoryInput) ([]*model.StorageLocationCategory, error) {
	var storageLocationCategories []*model.StorageLocationCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocationCategories, `select * from omegarogue.public.storage_location_category`,
	)
	if err != nil {
		return nil, errs.Wrap(err, "get storage location categories")
	}
	return storageLocationCategories, nil
}

func (r *queryResolver) StorageLocationCategory(ctx context.Context, id snowflake.ID) (*model.StorageLocationCategory, error) {
	var storageLocationCategory model.StorageLocationCategory
	// language=SQL
	err := pgxscan.Get(
		ctx, r.DbPool, &storageLocationCategory, `select * from omegarogue.public.storage_location_category where id=?`,
		id,
	)
	if err != nil {
		return nil, errs.Wrap(err, "get storage location category")
	}
	return &storageLocationCategory, nil
}

func (r *queryResolver) SiPrefixes(ctx context.Context, params *model.QuerySiPrefixInput) ([]*model.SiPrefix, error) {
	var siPrefixes []*model.SiPrefix
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &siPrefixes, `select * from omegarogue.public.si_prefix`)
	if err != nil {
		return nil, errs.Wrap(err, "get si prefixes")
	}
	return siPrefixes, nil
}

func (r *queryResolver) SiPrefix(ctx context.Context, id snowflake.ID) (*model.SiPrefix, error) {
	var siPrefix model.SiPrefix
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &siPrefix, `select * from omegarogue.public.si_prefix where id=?`, id)
	if err != nil {
		return nil, errs.Wrap(err, "get si prefix")
	}
	return &siPrefix, nil
}

func (r *queryResolver) Units(ctx context.Context, params *model.QueryUnitInput) ([]*model.Unit, error) {
	var units []*model.Unit
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &units, `select * from omegarogue.public.unit`)
	if err != nil {
		return nil, errs.Wrap(err, "get units")
	}
	return units, nil
}

func (r *queryResolver) Unit(ctx context.Context, id snowflake.ID) (*model.Unit, error) {
	var unit model.Unit
	// language=SQL
	err := pgxscan.Get(ctx, r.DbPool, &unit, `select * from omegarogue.public.unit where id=?`, id)
	if err != nil {
		return nil, errs.Wrap(err, "get unit")
	}
	return &unit, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*model.File, error) {
	return r.Resolver.Files(ctx)
}

func (r *queryResolver) File(ctx context.Context, id snowflake.ID) (*model.File, error) {
	return r.Resolver.File(ctx, &id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
