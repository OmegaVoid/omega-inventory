package model

import "github.com/bwmarrin/snowflake"

type PartCategory struct {
	ID          snowflake.ID    `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	PartsID     []*snowflake.ID `json:"parts"`
	ChildrenID  []*snowflake.ID `json:"children"`
	ParentID    *snowflake.ID   `json:"parent"`
	Root        *bool           `json:"root"`
}

func (PartCategory) IsNode() {}
