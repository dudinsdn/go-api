package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	DSN string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func Load() *Config {
	return &Config{
		DB:     DBConfig{DSN: "din:dbdin123@tcp(localhost:3306)/go_api"},
		Server: ServerConfig{Port: ":8080"},
	}
}

func InitDB(cfg DBConfig) *sql.DB {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		panic(err)
	}
	return db
}
