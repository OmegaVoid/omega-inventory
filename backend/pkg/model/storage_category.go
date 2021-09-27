package model

import "github.com/bwmarrin/snowflake"

type StorageLocationCategory struct {
	ID                 snowflake.ID    `json:"id"`
	Name               string          `json:"name"`
	Description        *string         `json:"description"`
	StorageLocationsID []*snowflake.ID `json:"storageLocations"`
	Root               *bool           `json:"root"`
	ParentID           *snowflake.ID   `json:"parent"`
	ChildrenID         []*snowflake.ID `json:"children"`
}

func (StorageLocationCategory) IsNode() {}
