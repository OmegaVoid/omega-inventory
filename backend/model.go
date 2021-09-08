package main

import (

	"github.com/beego/beego/v2/client/orm"
)

type Orm struct {
	Id int
}


type PartCategory struct {
	Orm
	Name  string
	Description string
	Parent *PartCategory `orm:"null;rel(fk)"`
	Children      []*PartCategory `orm:"null;reverse(many)"`
	Parts []*Part `orm:"null;reverse(many)"`
}

type PartMeasurementUnit struct {
	Orm
	Name string
	ShortName string
	Parts []*Part `orm:"reverse(many)"`
}

type StorageLocation struct {
	Orm
	Name  string
	// TODO Image StorageLocationImage
	Category *StorageLocationCategory `orm:"rel(fk)"`
	Parts []*Part `orm:"null;reverse(many)"`
}

type StorageLocationCategory struct {
	Orm
	Parent *StorageLocationCategory `orm:"null;rel(fk)"`
	Name  string
	Description string
	Children      []*StorageLocationCategory `orm:"null;reverse(many)"`
	StorageLocations []*StorageLocation `orm:"null;reverse(many)"`
}


type Part struct {
	Orm
	Name  string
	Category *PartCategory `orm:"rel(fk)"`
	Description string
	// TODO Footprint Footprint
	Unit *PartMeasurementUnit `orm:"rel(fk)"`
	StorageLocation *StorageLocation `orm:"rel(fk)"`
	// TODO Manufacturers []PartManufacturer
	// TODO Distributors []PartDistributor
	// TODO Attachments []PartAttachment
	Comment string
	StockLevel float64
	MinStockLevel float64
	// TODO ProjectParts []ProjectPart
	InternalPartNumber string
}


func init() {
	orm.RegisterModel(new(PartCategory),new(PartMeasurementUnit),new(StorageLocation),new(StorageLocationCategory),new(Part))
}

