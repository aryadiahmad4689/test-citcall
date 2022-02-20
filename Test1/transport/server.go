package transport

import (
	"github.com/aryadiahmad4689/test-citcall/controller"
	"github.com/aryadiahmad4689/test-citcall/repository"
	"github.com/aryadiahmad4689/test-citcall/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(e *gin.Engine) *Server {
	repo := repository.NewCoutryRepository()
	service := services.NewCountryService(repo)
	controller.NewHandler(e, service)
	return &Server{
		engine: e,
	}
}

func (s *Server) Start() {

	_ = s.engine.Run(":" + "8080")

}
