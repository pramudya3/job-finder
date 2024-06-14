package server

import (
	companyApi "job-finder/internal/company/api/http"
	jobApi "job-finder/internal/job/api/http"
	"job-finder/pkg/config"
	"log"

	res "net/http"

	_ "job-finder/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Server struct {
	ginEngine *gin.Engine
	db        *gorm.DB
	cfg       *config.Config
}

func NewServer(db *gorm.DB, cfg *config.Config) *Server {
	return &Server{
		ginEngine: gin.Default(),
		db:        db,
		cfg:       cfg,
	}
}

func (s *Server) Run() error {
	s.MapRoutes()
	s.ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.ginEngine.GET("/healthcheck", func(c *gin.Context) {
		c.Status(res.StatusOK)
	})

	log.Fatalln(s.ginEngine.Run(s.cfg.ServerAddr))
	return nil
}

func (s *Server) MapRoutes() {
	v1 := s.ginEngine.Group("/api/v1")
	companyApi.CompanyRoutes(v1, s.db, s.cfg)
	jobApi.JobRoutes(v1, s.db, s.cfg)
}
