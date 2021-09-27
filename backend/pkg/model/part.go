package model

import "github.com/bwmarrin/snowflake"

type Part struct {
	ID                 snowflake.ID    `json:"id"`
	Name               string          `json:"name"`
	CategoryID         *snowflake.ID   `json:"category"`
	Description        *string         `json:"description"`
	FootprintID        *snowflake.ID   `json:"footprint"`
	UnitID             *snowflake.ID   `json:"unit"`
	StorageLocationID  *snowflake.ID   `json:"storageLocation"`
	AttachmentsID      []*snowflake.ID `json:"attachments"`
	Comment            *string         `json:"comment"`
	StockLevel         float64         `json:"stockLevel"`
	MinStockLevel      float64         `json:"minStockLevel"`
	ParametersID       []*snowflake.ID `json:"parameters"`
	InternalPartNumber string          `json:"internalPartNumber"`
}

func (Part) IsNode() {}
