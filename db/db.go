package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"

	// Postgres driver
	_ "github.com/lib/pq"
)

// SQL struct
type SQL struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
	Db       *sqlx.DB
}

// Connect to Postgres
func (s *SQL) Connect() {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.Username, s.Password, s.DbName)
	s.Db = sqlx.MustConnect("postgres", dataSource)

	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
	}
}

// Close db
func (s *SQL) Close() {
	s.Db.Close()
}
