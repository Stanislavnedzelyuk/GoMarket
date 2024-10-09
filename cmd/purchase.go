package cmd

type Purchase struct {
	ID        int // уникальный идентификатор покупки
	UserID    int // идентификатор пользователя, совершившего покупку
	ProductID int // идентификатор купленного товара
	Quantity  int // количество купленного товара
	Timestamp int // временная метка покупки
}
