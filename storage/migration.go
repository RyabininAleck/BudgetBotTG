package storage

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// Migrator структура для применения миграций.
type Migrator struct {
	srcDriver source.Driver // Драйвер источника миграций.
}

// MustGetNewMigrator создает новый экземпляр Migrator с встроенными SQL-файлами миграций.
// В случае ошибки вызывает panic.
func MustGetNewMigrator(sqlFiles embed.FS, dirName string) *Migrator {
	// Создаем новый драйвер источника миграций с встроенными SQL-файлами.
	d, err := iofs.New(sqlFiles, dirName)
	if err != nil {
		panic(err)
	}
	return &Migrator{
		srcDriver: d,
	}
}

// ApplyMigrations применяет миграции к базе данных.
func (m *Migrator) ApplyMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("unable to create sqliteDB instance: %v", err)
	}

	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", m.srcDriver, "sqlite_db", driver)
	if err != nil {
		return fmt.Errorf("unable to create migration: %v", err)
	}

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to apply migrations %v", err)
	}

	return nil
}
