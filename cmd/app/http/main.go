package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/assignment-amori/cmd/helper"
	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	channelHdlr "github.com/assignment-amori/internal/handler/channel/http"
	fileHdlr "github.com/assignment-amori/internal/handler/file/http"
	userHdlr "github.com/assignment-amori/internal/handler/user/http"
	channelRepository "github.com/assignment-amori/internal/repository/channel"
	messageRepository "github.com/assignment-amori/internal/repository/message"
	openaiRepository "github.com/assignment-amori/internal/repository/openai"
	userRepository "github.com/assignment-amori/internal/repository/user"
	openaiService "github.com/assignment-amori/internal/service/openai"
	channelUsecase "github.com/assignment-amori/internal/usecase/channel"
	fileUsecase "github.com/assignment-amori/internal/usecase/file"
	userUsecase "github.com/assignment-amori/internal/usecase/user"
	handhelp "github.com/assignment-amori/middleware/http"
	"github.com/assignment-amori/pkg/consistency/enforcer"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	var (
		err error
		ctx = context.Background()
	)

	// Load .envrc file for dev env.
	err = godotenv.Load(".envrc")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return
	}

	// App Initializer.
	genericMod, err := helper.AppInitializer(ctx)
	if err != nil {
		log.Fatalf("Failed to init app: %v", err)
		return
	}

	slog.Info("Starting Application")
	err = startApp(ctx, genericMod)
	return
}

func startApp(ctx context.Context, genericMod *helper.GenericModulesResult) error {
	// Init database.
	slog.Info("Connect to Database")
	db, err := pgx.New(ctx, genericMod.AppConfig.Database)
	if err != nil {
		return errorwrapper.Wrap(err, errorwrapper.ErrIDDatabaseInit)
	}
	defer db.Close()

	consistency := enforcer.New(db)

	//------------------------------------------------------
	slog.Info("Start Initializing Service")
	//------------------------------------------------------

	// OpenAI service.
	openaiSvc := openaiService.New(genericMod.AppConfig.OpenAIConfig)

	//------------------------------------------------------
	slog.Info("Start Initializing Repository")
	//------------------------------------------------------

	// User Repo.
	userRepo := userRepository.New(db, genericMod.SonyFlake, genericMod.AppConfig.JWTConfig.SecretKey)

	// Channel Repo.
	channelRepo := channelRepository.New(db, genericMod.SonyFlake)

	// Message Repo.
	messageRepo := messageRepository.New(db, genericMod.SonyFlake)

	// OpenAI Repo.
	openaiRepo := openaiRepository.New(openaiSvc)

	//------------------------------------------------------
	slog.Info("Start Initializing Usecase")
	//------------------------------------------------------

	// Channel UC.
	channelUC := channelUsecase.New(
		consistency,
		channelRepo,
		messageRepo,
		openaiRepo,
		userRepo,
		genericMod.SonyFlake,
	)

	// File UC.
	fileUC := fileUsecase.New()

	// User UC.
	userUC := userUsecase.New(userRepo)

	//------------------------------------------------------
	slog.Info("Start Initializing Handler")
	//------------------------------------------------------

	// Channel Handler.
	channelHandler := channelHdlr.New(channelUC)

	// Upload Handler.
	uploadHandler := fileHdlr.New(fileUC)

	// User Handler.
	userHandler := userHdlr.New(userUC)

	slog.Info("Creating endpoint route")
	router := newRoutes(
		genericMod.AppConfig,
		channelHandler,
		uploadHandler,
		userHandler,
	)

	slog.Info("Application Start")
	return http.ListenAndServe(fmt.Sprintf(":%s", genericMod.AppConfig.ServerConfig.PortHTTP), router)
}

func newRoutes(
	appConfig entity.AppConfig,
	channelHandler *channelHdlr.Handler,
	fileHandler *fileHdlr.Handler,
	userHandler *userHdlr.Handler,
) *chi.Mux {
	router := chi.NewRouter()
	helperModule := handhelp.New(
		appConfig,
	)

	// Create a new CORS middleware with default options.
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers.
	})

	router.Use(corsMiddleware.Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(httprate.LimitByIP(100, 1*time.Minute))

	router.Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {
			r.Use(helperModule.APIVersion("v1"))

			r.Group(func(r chi.Router) {
				r.Mount(helper.GenModulePattern(constant.ModuleChannels), channelHandler.Routes(helperModule))
				r.Mount(helper.GenModulePattern(constant.ModuleFiles), fileHandler.Routes(helperModule))
				r.Mount(helper.GenModulePattern(constant.ModuleUsers), userHandler.Routes(helperModule))
			})
		})
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("ok"))
	})

	return router
}
