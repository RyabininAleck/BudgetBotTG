package sqliteDB

import (
	"database/sql"
	"log"

	"BudgetBotTG/config"
	"BudgetBotTG/middleware"
	"BudgetBotTG/storage"
)

func GetStorage(config config.ServerConfig, myMiddlewares map[string]middleware.MW) storage.Storage {
	connectionStr := config.DB //todo path to sqliteDB
	connectionStr = "./storage/sqlite_storage.sqliteDB"
	conn, err := sql.Open("sqlite3", connectionStr)
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}

	return storage.CreateStorage(conn, myMiddlewares)

}
