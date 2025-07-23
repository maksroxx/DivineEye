package alert

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlertHandler struct {
	AlertService AlertService
}

func NewAlertHandler(svc AlertService) *AlertHandler {
	return &AlertHandler{AlertService: svc}
}

func (h *AlertHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")

	var req struct {
		Coin      string  `json:"coin"`
		Direction string  `json:"direction"`
		Price     float64 `json:"price"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	id, err := h.AlertService.Create(c.Request.Context(), userID, req.Coin, req.Direction, req.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *AlertHandler) List(c *gin.Context) {
	userID := c.GetString("user_id")

	alerts, err := h.AlertService.Get(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alerts)
}

func (h *AlertHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}

	if err := h.AlertService.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": id})
}
