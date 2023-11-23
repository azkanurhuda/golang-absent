package postgres

import (
	"github.com/azkanurhuda/golang-absent/pkg/conf"
	"github.com/davecgh/go-spew/spew"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewCon() (*gorm.DB, error) {
	dsn := conf.Postgres.DSN()
	spew.Dump(dsn)
	dialector := pg.Open(dsn)
	cnf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}

	return gorm.Open(dialector, cnf)
}
