package entity

import "github.com/assignment-amori/internal/entity/generic"

type Channel struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
	generic.MetaInfo
}

type NewChannelUCRequest struct {
	Name          string                   `json:"name"`
	MessageSource []MessageSourceUCRequest `json:"messageSource"`
}

type NewChannelParams struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
}
