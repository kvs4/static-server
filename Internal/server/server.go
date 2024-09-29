package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kvs4/static-server/Internal/config"
	"github.com/kvs4/static-server/Internal/storage"
)

type Server struct {
	engin *gin.Engine
	port  string
}

func New(cfg *config.Config) *Server {
	fs := storage.NewFileService(cfg.StaticDir)
	router := NewRouter(fs)

	return &Server{
		engin: router,
		port:  cfg.Port,
	}
}

func (s *Server) ListenAndServe() error {
	return s.engin.Run(fmt.Sprintf(":%s", s.port))
}
