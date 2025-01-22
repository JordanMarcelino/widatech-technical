package provider

import (
	"github.com/JordanMarcelino/widatech-technical/internal/config"
	"github.com/JordanMarcelino/widatech-technical/internal/database"
	"github.com/JordanMarcelino/widatech-technical/internal/repository"
	"github.com/jmoiron/sqlx"
)

var (
	db        *sqlx.DB
	dataStore repository.DataStore
)

func InitGlobal(cfg *config.Config) {
	db = database.InitPostgres(cfg)
	dataStore = repository.NewDataStore(db)
}
