package model

import "github.com/bwmarrin/snowflake"

type StorageLocationCategory struct {
	ID          snowflake.ID `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Root        *bool        `json:"root"`
	Path        string
}

func (StorageLocationCategory) IsNode() {}
