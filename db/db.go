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
	DB       *sqlx.DB
}

// Connect to Postgres
func (s *SQL) Connect() {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.Username, s.Password, s.DbName)
	s.DB = sqlx.MustConnect("postgres", dataSource)

	if err := s.DB.Ping(); err != nil {
		log.Error(err.Error())
	}
}

// Close db
func (s *SQL) Close() {
	s.DB.Close()
}
