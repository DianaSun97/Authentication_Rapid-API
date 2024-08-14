package dbrepo

import (
	"database/sql"
	"github.com/DianaSun97/elluliin_booking/internal/config"
	"github.com/DianaSun97/elluliin_booking/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
