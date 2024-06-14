package usecase

import (
	"context"
	"job-finder/internal/job/model"
	"job-finder/internal/job/repository"
	"job-finder/pkg/config"
	"time"

	"github.com/google/uuid"
)

//go:generate mockery --name=JobUsecase
type JobUsecase interface {
	Create(ctx context.Context, job *model.CreateJobReq) error
	FindJob(ctx context.Context, query map[string]interface{}) ([]*model.Job, interface{}, error)
}

type jobUsecase struct {
	jobRepository repository.JobRepository
	timeout       time.Duration
}

func (j *jobUsecase) Create(ctx context.Context, job *model.CreateJobReq) error {
	ctx, cancel := context.WithTimeout(ctx, j.timeout)
	defer cancel()

	newJob := &model.Job{
		ID:          uuid.New(),
		CompanyID:   job.CompanyID,
		Title:       job.Title,
		Description: job.Description,
		CreatedAt:   time.Now(),
	}

	return j.jobRepository.Create(ctx, newJob)
}

func (j *jobUsecase) FindJob(ctx context.Context, query map[string]interface{}) ([]*model.Job, interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, j.timeout)
	defer cancel()

	return j.jobRepository.FindJob(ctx, query)
}

func NewJobUsecase(jobRepo repository.JobRepository, cfg *config.Config) JobUsecase {
	return &jobUsecase{
		jobRepository: jobRepo,
		timeout:       time.Duration(cfg.CtxTimeout) * time.Second,
	}
}
