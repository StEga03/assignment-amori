package entity

import (
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
)

type ServerConfig struct {
	PortHTTP  string
	MachineID uint16
}

type OpenAIConfig struct {
	APIKey string
}

type JWTConfig struct {
	SecretKey string
}

type AppConfig struct {
	ServerConfig
	OpenAIConfig
	JWTConfig
	Database     pgx.DBConfig
	ErrorWrapper errorwrapper.Config
}
