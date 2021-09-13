package graph

import (
	"github.com/OmegaVoid/omega-inventory/pkg/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	measurementUnits          []*model.PartMeasurementUnit
	parts                     []*model.Part
	partCategories            []*model.PartCategory
	storageLocations          []*model.StorageLocation
	storageLocationCategories []*model.StorageLocationCategory
}
