package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hiodoshi-go/internal/models"
	"hiodoshi-go/internal/services"
)

type HiodoshiHandler struct {
	service *services.HiodoshiService
}

func NewHiodoshiHandler(service *services.HiodoshiService) *HiodoshiHandler {
	return &HiodoshiHandler{service: service}
}

func (h *HiodoshiHandler) GetHiodoshi(c *gin.Context) {
	hiodoshi := h.service.GenerateRandomHiodoshi()

	
	hiodoshiResult := models.Hiodoshi{
		Message: hiodoshi,
	}

	c.JSON(http.StatusOK, hiodoshiResult)
}