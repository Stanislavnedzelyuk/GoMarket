package internal

import "database/sql"

type Product struct {
	ID          int    // Уникальный идентификатор товара
	Name        string // Название товара
	Description string // Описание товара
	Price       int    // Цена товара
	Quantity    string // Количество товара на складе
}

func AddProduct(db *sql.DB, name, description string, price float64, quantity int) error {
	query := `INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, name, description, price, quantity)
	return err
}
