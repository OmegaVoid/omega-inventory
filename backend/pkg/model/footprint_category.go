package model

import "github.com/bwmarrin/snowflake"

type FootprintCategory struct {
	ID          snowflake.ID `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Path        string
}

func (FootprintCategory) IsNode() {}
