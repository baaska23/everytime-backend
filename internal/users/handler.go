package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) FindOrCreateUser(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	user, err := h.service.FindOrCreateUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find or create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) GetUserById(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	user, err := h.service.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user by id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
