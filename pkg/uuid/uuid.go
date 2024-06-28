package uuid

import "github.com/google/uuid"

var GenUUID = uuid.New

// UUID return method of UUID generator.
//
//go:generate mockgen -destination ./../../gen/mock/pkg/uuid/uuid.go -package mock_uuid . UUID
type UUID interface {
	GenUUID() string
}

type uuidImp struct{}

// New represent function for initialize uuid function.
func New() UUID {
	return &uuidImp{}
}

// GenUUID represent a function for generate UUID and return it as a string.
func (u *uuidImp) GenUUID() string {
	return uuid.New().String()
}
