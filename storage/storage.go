package storage

import (
	"database/sql"
	"embed"
	"fmt"

	"BudgetBotTG/middleware"
)

type Storage struct {
	conn        *sql.DB
	middlewares middleware.MWs
}

func CreateStorage(myConn *sql.DB, myMiddlewares middleware.MWs) Storage {
	return Storage{
		conn:        myConn,
		middlewares: myMiddlewares,
	}
}

//go:embed  migrations/1_innit_database.up.sql
var MigrationsFS embed.FS

func (s *Storage) MakeMigrations() {
	//todo const migrationsDir = "migrations" достать из конфига
	migrator := MustGetNewMigrator(MigrationsFS, "migrations")
	err := migrator.ApplyMigrations(s.conn)
	if err != nil {
		panic(err)
	}
	// todo refactor
	fmt.Println("Migrations applied!!")
}

func (s *Storage) Ping() error {
	return s.conn.Ping()
}

func (s *Storage) GetMiddlewares() middleware.MWs {
	middlewareMap := make(middleware.MWs)

	for key, mw := range s.middlewares {
		// todo Получаем указатель на конкретный тип middleware.MW
		switch key {
		case middleware.TimeChecker:
			mwPtr := mw.(*middleware.TimeCheck) // Предположим, что ConcreteMiddleware - это конкретная реализация интерфейса MW
			// Создаем копию mwPtr
			copiedMw := *mwPtr // Это создаст копию структуры ConcreteMiddleware
			// Записываем копию в middlewareMap
			middlewareMap[key] = &copiedMw // Сохраняем указатель на копию структуры в map
		}

	}
	return middlewareMap
}
