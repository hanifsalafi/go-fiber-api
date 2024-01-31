package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type app = struct {
	Name        string        `toml:"name"`
	Port        string        `toml:"port"`
	PrintRoutes bool          `toml:"print-routes"`
	Prefork     bool          `toml:"prefork"`
	Production  bool          `toml:"production"`
	IdleTimeout time.Duration `toml:"idle-timeout"`
}

// db struct config
type db = struct {
	Postgres struct {
		DSN     string `toml:"dsn"`
		Migrate bool   `toml:"migrate"`
		Seed    bool   `toml:"seed"`
	}
}

// log struct config
type logger = struct {
	TimeFormat string        `toml:"time-format"`
	Level      zerolog.Level `toml:"level"`
	Prettier   bool          `toml:"prettier"`
}

// middleware
type middleware = struct {
	Compress struct {
		Enable bool
		Level  compress.Level
	}

	Recover struct {
		Enable bool
	}

	Monitor struct {
		Enable bool
		Path   string
	}

	Pprof struct {
		Enable bool
	}

	Limiter struct {
		Enable     bool
		Max        int
		Expiration time.Duration `toml:"expiration_seconds"`
	}
}

type Config struct {
	App        app
	DB         db
	Logger     logger
	Middleware middleware
}

// NewConfig : initialize config
func NewConfig() *Config {
	config, err := ParseConfig("config")
	if err != nil && !fiber.IsChild() {
		// panic if config is not found
		log.Panic().Err(err).Msg("config not found")
	}

	return config
}

// ParseConfig : func to parse config
func ParseConfig(name string, debug ...bool) (*Config, error) {
	var (
		contents *Config
		file     []byte
		err      error
	)

	if len(debug) > 0 {
		file, err = os.ReadFile(name)
	} else {
		_, b, _, _ := runtime.Caller(0)
		// get base path
		path := filepath.Dir(filepath.Dir(filepath.Dir(b)))
		file, err = os.ReadFile(filepath.Join(path, "./config/toml/", name+".toml"))
	}

	if err != nil {
		return &Config{}, err
	}

	err = toml.Unmarshal(file, &contents)

	return contents, err
}

// ParseAddress : func to parse address
func ParseAddress(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i > 0 {
		return raw[:i], raw[i+1:]
	}

	return raw, ""
}
