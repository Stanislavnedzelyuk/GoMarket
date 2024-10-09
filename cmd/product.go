package cmd

type Product struct {
	ID          int    // уникальный идентификатор товара
	Name        string // название товара
	Description string // описание товара
	Price       int    // цена товара
	Quantity    string // количество товара на складе
}
