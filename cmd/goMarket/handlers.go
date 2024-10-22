package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gomarket/db"
	"gomarket/models"
)

func RegisterUser(c *gin.Context) {
	log.Println("Attempting to register a new user...")
	var newUser models.User
	// Привязка JSON данных к структуре пользователя
	if err := c.BindJSON(&newUser); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерация уникального ID для пользователя
	newUser.ID = uuid.New().String()
	// Хэширование пароля перед сохранением
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = hashedPassword

	// Сохранение пользователя в базе данных
	_, err = db.DB.Exec("INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4)", newUser.ID, newUser.Username, newUser.Password, newUser.Email)
	if err != nil {
		log.Printf("Error inserting user into database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	log.Printf("User %s registered successfully", newUser.Username)
	c.JSON(http.StatusCreated, newUser)
}

func LoginUser(c *gin.Context) {
	log.Println("Attempting to log in user...")
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	// Привязка JSON данных к структуре данных для логина
	if err := c.BindJSON(&loginData); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Поиск пользователя по имени
	var user models.User
	err := db.DB.QueryRow("SELECT id, username, password, email FROM users WHERE username=$1", loginData.Username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Проверка пароля
	if verifyPassword(user.Password, loginData.Password) != nil {
		log.Printf("Failed login attempt for username: %s", loginData.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	log.Printf("User %s logged in successfully", user.Username)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Fetching profile for user ID: %s", id)
	// Получение пользователя по ID
	var user models.User
	err := db.DB.QueryRow("SELECT id, username, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Printf("User ID %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
