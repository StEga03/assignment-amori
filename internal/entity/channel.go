package entity

import "github.com/assignment-amori/internal/entity/generic"

type Channel struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
	generic.MetaInfo
}

type ChannelUCRequest struct {
	Name          string                   `json:"name"`
	MessageSource []MessageSourceUCRequest `json:"messageSource"`
}

type ChannelParams struct {
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
}
