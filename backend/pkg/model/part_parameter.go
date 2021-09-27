package model

import "github.com/bwmarrin/snowflake"

type PartParameter struct {
	ID                 snowflake.ID  `json:"id"`
	PartID             *snowflake.ID `json:"part"`
	Name               string        `json:"name"`
	Description        *string       `json:"description"`
	UnitID             *snowflake.ID `json:"unit"`
	Value              *float64      `json:"value"`
	NormalizedValue    *float64      `json:"normalizedValue"`
	MaxValue           *float64      `json:"maxValue"`
	NormalizedMaxValue *float64      `json:"normalizedMaxValue"`
	MinValue           *float64      `json:"minValue"`
	NormalizedMinValue *float64      `json:"normalizedMinValue"`
	StringValue        *string       `json:"stringValue"`
	ValueTypeID        snowflake.ID  `json:"valueType"`
	SiPrefixID         *snowflake.ID `json:"siPrefix"`
	MinSiPrefixID      *snowflake.ID `json:"minSiPrefix"`
	MaxSiPrefixID      *snowflake.ID `json:"maxSiPrefix"`
}

func (PartParameter) IsNode() {}
