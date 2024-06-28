package helper

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/locale"
	"github.com/assignment-amori/pkg/sony"
	"github.com/assignment-amori/pkg/sql/pgx"
	timeutils "github.com/assignment-amori/pkg/time_utils"
	"github.com/assignment-amori/pkg/validator"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

type GenericModulesResult struct {
	AppConfig entity.AppConfig
	SonyFlake *sonyflake.Sonyflake
}

func AppInitializer(ctx context.Context) (*GenericModulesResult, error) {
	var (
		result *GenericModulesResult
		err    error
	)

	// Initialize config.
	appConfig, err := initAppConfig()
	if err != nil {
		log.Fatalf("Failed to init config: %v", err)
		return result, err
	}

	// Init errorwrapper.
	err = errorwrapper.NewErrorWrapper(appConfig.ErrorWrapper)
	if err != nil {
		return result, err
	}

	// Init localizer.
	err = locale.NewLocale()
	if err != nil {
		return result, err
	}

	// Init validator.
	err = validator.New()
	if err != nil {
		return result, err
	}

	// Initialize time location.
	err = timeutils.NewTimeLocation()
	if err != nil {
		return nil, err
	}

	// Init Sony Flake.
	sf := sony.NewIDGenerator(appConfig.ServerConfig.MachineID)

	return &GenericModulesResult{
		AppConfig: appConfig,
		SonyFlake: sf,
	}, nil
}

func initAppConfig() (entity.AppConfig, error) {
	// Parse the string as an unsigned integer with base 10.
	parsedMachineID, err := strconv.ParseUint(os.Getenv("MACHINE_ID"), 10, 16)
	if err != nil {
		return entity.AppConfig{}, errors.Wrap(err, "failed to parse machine id")
	}

	// Convert to uint16.
	machineID := uint16(parsedMachineID)

	return entity.AppConfig{
		ServerConfig: entity.ServerConfig{
			PortHTTP:  os.Getenv("PORT_HTTP"),
			MachineID: machineID,
		},
		OpenAIConfig: entity.OpenAIConfig{
			APIKey: os.Getenv("OPENAI_API_KEY"),
		},
		JWTConfig: entity.JWTConfig{
			SecretKey: os.Getenv("JWT_SECRET_KEY"),
		},
		Database: pgx.DBConfig{
			DSN: os.Getenv("DB_DSN"),
		},
	}, nil
}
