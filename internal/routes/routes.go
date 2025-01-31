package routes

import (
	"go-todo/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register all routes
func RegisterRoutes(router *gin.Engine, db *gorm.DB){
	h := &handlers.Handler{DB: db}

	router.POST("/todo", h.CreateTodoHandler)
	router.GET("todo", h.ListToDo)
}