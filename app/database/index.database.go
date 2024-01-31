package database

import (
	"github.com/rs/zerolog"
	"go-fiber-api/app/database/entity"
	"go-fiber-api/config/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database setup database with gorm
type Database struct {
	DB  *gorm.DB
	Log zerolog.Logger
	Cfg *config.Config
}

type Seeder interface {
	Seed(*gorm.DB) error
	Count(*gorm.DB) (int, error)
}

func NewDatabase(cfg *config.Config, log zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

// ConnectDatabase connect database
func (_db *Database) ConnectDatabase() {
	conn, err := gorm.Open(postgres.Open(_db.Cfg.DB.Postgres.DSN), &gorm.Config{})
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to connect the database!")
	} else {
		_db.Log.Info().Msg("Connected the database succesfully!")
	}

	_db.DB = conn
}

// ShutdownDatabase shutdown database
func (_db *Database) ShutdownDatabase() {
	sqlDB, err := _db.DB.DB()
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to shutdown the database!")
	} else {
		_db.Log.Info().Msg("Shutdown the database succesfully!")
	}
	sqlDB.Close()
}

// MigrateModels migrate models
func (_db *Database) MigrateModels() {
	err := _db.DB.AutoMigrate(
		Models()...,
	)
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to migrate the database!")
	} else {
		_db.Log.Info().Msg("Migrate the database entity succesfully!")
	}
}

// Models list of models for migration
func Models() []interface{} {
	return []interface{}{
		entity.MasterStatus{},
		entity.User{},
		entity.UserRole{},
		entity.UserRoleAccess{},
	}
}

// SeedModels seed data
func (_db *Database) SeedModels(seeder []Seeder) {
	for _, seed := range seeder {
		count, err := seed.Count(_db.DB)
		if err != nil {
			_db.Log.Error().Err(err).Msg("An unknown error occurred when to seed the database!")
		}

		if count == 0 {
			if err := seed.Seed(_db.DB); err != nil {
				_db.Log.Error().Err(err).Msg("An unknown error occurred when to seed the database!")
			}

			_db.Log.Info().Msg("Seeded the database successfully!")
		} else {
			_db.Log.Info().Msg("Database is already seeded!")
		}
	}

	_db.Log.Info().Msg("Seeded the database succesfully!")
}
