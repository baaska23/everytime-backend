package auth

import (
	"everytime-backend/internal/shared/apierror"
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
		apierror.WriteGin(c, apierror.BadRequest("userId is required"))
		return
	}

	user, err := h.service.FindOrCreateUser(userId)
	if err != nil {
		apierror.WriteGin(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) GetUserById(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		apierror.WriteGin(c, apierror.BadRequest("userId is required"))
		return
	}

	user, err := h.service.GetUserById(userId)
	if err != nil {
		apierror.WriteGin(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
