package main

import "github.com/gin-gonic/gin"

type Server struct {
	status int32
}

func NewServer() *Server {
	s := &Server{}
	return s
}

func (srv *Server) Run(port string) {
	engine := gin.New()

	engine.GET("/api/:word", srv.GetWordDefinition)
	engine.GET("/api/:word/minify=:minify", srv.GetWordDefinition)
	//engine.Any("health", srv.Health)

	engine.Run(port)
}
