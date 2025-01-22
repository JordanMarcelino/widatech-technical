package provider

import (
	"github.com/JordanMarcelino/widatech-technical/internal/config"
	"github.com/JordanMarcelino/widatech-technical/internal/database"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func InitGlobal(cfg *config.Config) {
	db = database.InitPostgres(cfg)
}
