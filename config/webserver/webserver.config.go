package webserver

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	futils "github.com/gofiber/fiber/v2/utils"
	"github.com/rs/zerolog"
	"go-fiber-api/app/database"
	"go-fiber-api/app/database/seeds"
	md "go-fiber-api/app/middleware"
	"go-fiber-api/app/router"
	"go-fiber-api/config/config"
	"go-fiber-api/utils/response"
	"go.uber.org/fx"
	"os"
	"runtime"
	"strings"
	"time"
)

// NewFiber : initialize the webserver
func NewFiber(cfg *config.Config) *fiber.App {
	// setup
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.App.Name,
		AppName:               cfg.App.Name,
		Prefork:               cfg.App.Prefork,
		ErrorHandler:          response.ErrorHandler,
		IdleTimeout:           cfg.App.IdleTimeout * time.Second,
		EnablePrintRoutes:     cfg.App.PrintRoutes,
		DisableStartupMessage: true,
	})

	// pass production config to check it
	response.IsProduction = cfg.App.Production

	return app
}

// Start : function to start webserver
func Start(lifecycle fx.Lifecycle, cfg *config.Config, fiber *fiber.App, router *router.Router, middlewares *md.Middleware, db *database.Database, log zerolog.Logger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				middlewares.Register()
				router.Register()

				// Custom Startup Messages
				host, port := config.ParseAddress(cfg.App.Port)
				if host == "" {
					if fiber.Config().Network == "tcp6" {
						host = "[::1]"
					} else {
						host = "0.0.0.0"
					}
				}

				// ASCII Art
				ascii, err := os.ReadFile("./storage/ascii_art.txt")
				if err != nil {
					log.Debug().Err(err).Msg("An unknown error occurred when to print ASCII art!")
				}

				for _, line := range strings.Split(futils.UnsafeString(ascii), "\n") {
					log.Info().Msg(line)
				}

				// Information message
				log.Info().Msg(fiber.Config().AppName + " is running at the moment!")

				// Debug informations
				if !cfg.App.Production {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !cfg.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					log.Debug().Msgf("Version: %s", "-")
					log.Debug().Msgf("Host: %s", host)
					log.Debug().Msgf("Port: %s", port)
					log.Debug().Msgf("Prefork: %s", prefork)
					log.Debug().Msgf("Handlers: %d", fiber.HandlersCount())
					log.Debug().Msgf("Processes: %d", procs)
					log.Debug().Msgf("PID: %d", os.Getpid())
				}

				// Listen the app (with TLS Support)
				//if cfg.App.TLS.Enable {
				//	log.Debug().Msg("TLS support was enabled.")
				//
				//	if err := fiber.ListenTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {
				//		log.Error().Err(err).Msg("An unknown error occurred when to run server!")
				//	}
				//}

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()

				db.ConnectDatabase()

				migrateFlag := flag.Bool("migrate", db.Cfg.DB.Postgres.Migrate, "migrate the database")
				seedFlag := flag.Bool("seed", db.Cfg.DB.Postgres.Seed, "seed the database")
				flag.Parse()

				// read flag -migrate to migrate the database
				if *migrateFlag {
					db.MigrateModels()
				}
				// read flag -seed to seed the database
				if *seedFlag {

					// init seed models
					masterStatusSeeder := seeds.MasterStatusSeeder{}
					allSeeders := []database.Seeder{masterStatusSeeder}

					db.SeedModels(allSeeders)
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info().Msg("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic().Err(err).Msg("")
				}

				log.Info().Msg("Running cleanup tasks...")
				log.Info().Msg("1- Shutdown the database")
				db.ShutdownDatabase()
				log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
				log.Info().Msg("\u001b[96msee you again👋\u001b[0m")

				return nil
			},
		},
	)
}
