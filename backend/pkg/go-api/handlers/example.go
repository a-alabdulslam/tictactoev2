package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tictactoe/pkg/go-api/repository"
)

func NewExampleHandler() IExampleHandler {
	repo := repository.NewExampleRepository()
	return &ExampleHandler{repo}
}

type IExampleHandler interface {
	HelloName(c *gin.Context)
}

type ExampleHandler struct {
	repo repository.IExampleRepository
}

func (h *ExampleHandler) HelloName(c *gin.Context) {
	name := c.Param("name")
	// validate name and return 400 if invalid
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
	}
	namefromdb := h.repo.GetName(name)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + namefromdb,
	})
}
