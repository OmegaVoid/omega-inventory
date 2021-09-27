package graph

import (
	"github.com/OmegaVoid/omega-inventory/pkg/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	measurementUnits          []*model.PartMeasurementUnit
	footprints                []*model.Footprint
	footprintCategories       []*model.FootprintCategory
	footprintAttachments      []*model.FootprintAttachment
	parts                     []*model.Part
	partCategories            []*model.PartCategory
	storageLocations          []*model.StorageLocation
	storageLocationCategories []*model.StorageLocationCategory

	DbPool *pgxpool.Pool
}
