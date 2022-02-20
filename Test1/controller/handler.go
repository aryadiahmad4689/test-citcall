package controller

import (
	"fmt"

	"github.com/aryadiahmad4689/test-citcall/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	postService services.CountryService
}

func NewHandler(e *gin.Engine, postService services.CountryService) {
	handler := &Handler{postService: postService}
	e.GET("/countrys", handler.GetAll)
}

func (h *Handler) GetAll(c *gin.Context) {
	fmt.Println("Handler GetAll")
	countrys, err := h.postService.GetAll()
	if err != nil {
		c.HTML(500, "error.tmpl", gin.H{})
	}
	c.HTML(200, "index.html", gin.H{"countrys": countrys})
}
