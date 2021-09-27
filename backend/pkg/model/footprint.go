package model

import "github.com/bwmarrin/snowflake"

type Footprint struct {
	ID            snowflake.ID    `json:"id"`
	Name          string          `json:"name"`
	CategoryID    *snowflake.ID   `json:"category"`
	Description   *string         `json:"description"`
	AttachmentsID []*snowflake.ID `json:"attachments"`
	ImageID       *snowflake.ID   `json:"image"`
}

func (Footprint) IsNode() {}
