package model

import "github.com/bwmarrin/snowflake"

type PartCategory struct {
	ID          snowflake.ID `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Path        string
	Root        *bool `json:"root"`
}

func (PartCategory) IsNode() {}
