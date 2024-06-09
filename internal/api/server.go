package api

import (
	"github.com/bobbybof/timesheet/config"
	"github.com/bobbybof/timesheet/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	store  repository.Store
	router *gin.Engine
}

func NewServer(config config.Config, store repository.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.getRoutes()
	return server, nil
}

func (server *Server) getRoutes() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With", "X-XSRF-TOKEN"},
		AllowMethods: []string{"POST", "PUT", "GET", "DELETE", "PATCH"},
	}))

	router.GET("user/:id", server.GetUser)
	router.PUT("user/:id", server.UpdateUser)

	router.POST("activity", server.CreateActivity)
	router.PUT("activity/:id", server.UpdateActivity)
	router.DELETE("activity/:id", server.DeleteActivity)
	router.GET("activities", server.GetUserActivities)

	router.GET("projects", server.GetProjects)
	router.POST("project", server.CreateProject)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
