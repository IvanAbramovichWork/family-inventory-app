package handlers

import (
	"log"
	"net/http"

	"github.com/IvanAbramovichWork/family-inventory-app/app/database"

	"github.com/IvanAbramovichWork/family-inventory-app/app/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	DB *sqlx.DB
}

// NewUserHandler создает новый экземпляр UserHandler
func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// RegisterUser регистрирует нового пользователя
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Хеширование пароля может быть добавлено здесь
	err := database.CreateUser(h.DB, user)
	if err != nil {
		log.Println("Error registering user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// GetUser получает информацию о пользователе по его ID
func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := database.GetUserByID(h.DB, userID)
	if err != nil {
		log.Println("Error fetching user:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
