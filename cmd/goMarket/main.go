package main

import (
	"github.com/gin-gonic/gin"
	"gomarket/handlers"
	"log"
)

func main() {
	// Подключение к базе данных Postgres SQL
	db.Connect()
	defer db.Close()

	log.Println("Starting GoMarket server...")
	r := gin.Default()

	// Эндпоинты для пользователей
	r.POST("/register", handlers.RegisterUser)  // Регистрация нового пользователя
	r.POST("/login", handlers.LoginUser)        // Авторизация пользователя
	r.GET("/user/:id", handlers.GetUserProfile) // Получение профиля пользователя

	// Эндпоинты для управления товарами
	r.POST("/products", handlers.AddProduct)          // Добавление нового товара
	r.GET("/products", handlers.GetAllProducts)       // Получение списка всех товаров
	r.GET("/products/:id", handlers.GetProductByID)   // Получение товара по ID
	r.PUT("/products/:id", handlers.UpdateProduct)    // Обновление информации о товаре
	r.DELETE("/products/:id", handlers.DeleteProduct) // Удаление товара

	// Эндпоинты для управления покупками
	r.POST("/purchase", handlers.MakePurchase)              // Совершение покупки
	r.GET("/user/:id/purchases", handlers.GetUserPurchases) // Получение списка покупок пользователя

	r.Run() // Запуск сервера на порту 8080
}
