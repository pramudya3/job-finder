package http

import (
	"job-finder/internal/job/repository"
	"job-finder/internal/job/usecase"
	"job-finder/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JobRoutes(g *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	jobRepo := repository.NewJobRepository(db)
	jobUc := usecase.NewJobUsecase(jobRepo, cfg)
	jobHandler := NewJobHandler(jobUc)

	route := g.Group("/jobs")
	{
		route.POST("/", jobHandler.Create)
		route.GET("/", jobHandler.FindJob)
	}
}
