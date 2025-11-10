package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/example/meihua_gpt/backend/models"
	"github.com/example/meihua_gpt/backend/services"
)

type DivinationRequest struct {
	Subject string `json:"subject" binding:"required"`
	Method  string `json:"method"`
}

type DivinationHandler struct {
	db *gorm.DB
}

func RegisterDivinationRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := &DivinationHandler{db: db}
	group := rg.Group("/divinations")
	group.GET("", handler.list)
	group.GET(":id", handler.get)
	group.POST("", handler.create)
}

func (h *DivinationHandler) list(c *gin.Context) {
	var items []models.Divination
	if err := h.db.Order("created_at desc").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *DivinationHandler) get(c *gin.Context) {
	var item models.Divination
	if err := h.db.First(&item, c.Param("id")).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "divination not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *DivinationHandler) create(c *gin.Context) {
	var req DivinationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	outcome := services.GenerateDivination(req.Subject, req.Method)
	record := models.Divination{
		Subject:       req.Subject,
		Method:        req.Method,
		PrimaryName:   outcome.PrimaryName,
		UpperTrigram:  outcome.UpperTrigram,
		LowerTrigram:  outcome.LowerTrigram,
		SecondaryName: outcome.SecondaryName,
		Commentary:    outcome.Commentary,
		Lines:         outcome.Lines,
		ChangingLines: outcome.ChangingLines,
	}

	if err := h.db.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}
