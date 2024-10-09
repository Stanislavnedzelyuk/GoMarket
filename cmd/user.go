package cmd

type User struct {
	ID       int    // уникальный идентификатор пользователя
	Username string // логин пользователя
	Password string // хэшированный пароль
	Email    string // адрес электронной почты (опционально)
}
