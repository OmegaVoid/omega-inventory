package model

import "github.com/bwmarrin/snowflake"

type FootprintCategory struct {
	ID           snowflake.ID    `json:"id"`
	Name         string          `json:"name"`
	Description  *string         `json:"description"`
	ParentID     *snowflake.ID   `json:"parent"`
	ChildrenID   []*snowflake.ID `json:"children"`
	FootprintsID []*snowflake.ID `json:"footprints"`
}

func (FootprintCategory) IsNode() {}
