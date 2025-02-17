package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service_notification/internal/models"
	"service_notification/internal/service"
)

// NotificationController - структура контроллера
type NotificationController struct {
	service service.Service
}

func NewNotificationController(service *service.Service) *NotificationController {
	return &NotificationController{service: *service}
}

// CreateNotificationHandler - обработчик создания уведомления
func (nc *NotificationController) CreateNotificationHandler(c *gin.Context) {
	var req models.СreateNotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Создание уведомления в БД
	ctx := c.Request.Context()
	err := nc.service.SendNotifications(ctx, req.ID, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}
