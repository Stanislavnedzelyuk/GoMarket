package database

import (
	// Стандартный пакет Go для работы с базами данных
	"database/sql"
	// Пакет для логирования
	"log"
	// Импортируем драйвер PostgresSQL.
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres",
		"user=delta dbname=mydatabase sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	DB.Close()
}
