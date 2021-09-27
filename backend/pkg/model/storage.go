package model

import "github.com/bwmarrin/snowflake"

type StorageLocation struct {
	ID         snowflake.ID    `json:"id"`
	Name       string          `json:"name"`
	ImageID    *snowflake.ID   `json:"Image"`
	CategoryID *snowflake.ID   `json:"category"`
	PartsID    []*snowflake.ID `json:"parts"`
}

func (StorageLocation) IsNode() {}
