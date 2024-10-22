package internal

type User struct {
	ID       int    // Уникальный идентификатор пользователя
	Username string // Логин пользователя
	Password string // Хэшированный пароль
	Email    string // Адрес электронной почты (опционально)
}
