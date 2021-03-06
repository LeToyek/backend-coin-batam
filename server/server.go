package server

import (
	"project/handler"
	middleware "project/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Handler *handler.Handler
}

func (s *Server) StartServer() {
	s.Router.Use(CORSMiddleware())
	s.Router.Use(middleware.Authenticate())
	s.Router.POST("/register", handler.Register())
	s.Router.POST("/login", handler.Login())
	s.Router.GET("/index", s.Handler.TryIndex())
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
