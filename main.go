package main

import (
	"job-finder/app/server"
	company "job-finder/internal/company/model"
	job "job-finder/internal/job/model"
	"job-finder/pkg/config"
	psql "job-finder/pkg/dbs/postgres"
	"log"
)

//	@title			JobFinder Swagger API
//	@version		1.0
//	@description	Swagger API for JobFinder.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Nanda Rizky Pramudya
//	@contact.email	pramudya500@gmail.com

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/api/v1
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed load config, err: %v", err)
	}

	db, err := psql.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed initiate database, err: %v", err)
	}

	if err := db.AutoMigrate(&company.Company{}, &job.Job{}); err != nil {
		log.Fatalf("failed migrate database, err: %v", err)
	}

	server.NewServer(db, cfg).Run()
}
