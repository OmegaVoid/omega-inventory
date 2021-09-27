package model

import "github.com/bwmarrin/snowflake"

type Unit struct {
	ID         snowflake.ID    `json:"id"`
	Name       string          `json:"name"`
	Symbol     string          `json:"symbol"`
	PrefixesID []*snowflake.ID `json:"prefixes"`
}

func (Unit) IsNode() {}
