package http

import (
	"job-finder/internal/company/repository"
	"job-finder/internal/company/usecase"
	"job-finder/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CompanyRoutes(g *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	companyRepo := repository.NewCompanyRepository(db)
	companyUc := usecase.NewCompanyUsecase(companyRepo, cfg)
	companyHandler := NewCompanyHandler(companyUc)

	route := g.Group("/companies")
	{
		route.POST("/", companyHandler.Create)
		route.DELETE("/:id", companyHandler.DeleteByID)
		route.GET("/:id", companyHandler.FindByID)
	}
}
