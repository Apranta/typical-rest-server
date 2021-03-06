package typpostgres

import (
	"database/sql"
	"fmt"
)

func (m *Module) connect(cfg Config) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", m.dataSource(cfg))
	if err != nil {
		err = fmt.Errorf("Posgres: Connect: %w", err)
	}
	return
}

func (*Module) disconnect(db *sql.DB) (err error) {
	if err = db.Close(); err != nil {
		return fmt.Errorf("Postgres: Disconnect: %w", err)
	}
	return
}

func (m *Module) ping(db *sql.DB) (err error) {
	if err = db.Ping(); err != nil {
		return fmt.Errorf("Postgres: Ping: %w", err)
	}
	return
}

func (*Module) dataSource(c Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

func (*Module) adminDataSource(c Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, "template1")
}
