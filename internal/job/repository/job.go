package repository

import (
	"context"
	"job-finder/internal/job/model"
	"job-finder/pkg/pagination"

	"gorm.io/gorm"
)

//go:generate mockery --name=JobRepository
type JobRepository interface {
	Create(ctx context.Context, job *model.Job) error
	FindJob(ctx context.Context, query map[string]interface{}) ([]*model.Job, interface{}, error)
}

type jobRepository struct {
	db *gorm.DB
}

func (j *jobRepository) Create(ctx context.Context, job *model.Job) error {
	if err := j.db.Save(job).Error; err != nil {
		return err
	}
	return nil
}

func (j *jobRepository) FindJob(ctx context.Context, query map[string]interface{}) ([]*model.Job, interface{}, error) {
	var (
		total       int64
		jobs        = []*model.Job{}
		page        = query["page"].(int)
		size        = query["size"].(int)
		keyword     = query["keyword"].(string)
		companyName = query["companyName"].(string)
	)

	database := j.db.Model(jobs).Joins("left join companies on companies.id = jobs.company_id")

	if keyword != "" {
		database.Where("jobs.title ~* ? OR jobs.description ~* ?", keyword, keyword)
	}

	if companyName != "" {
		database.Where("companies.name ~* ?", companyName)
	}

	if err := database.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	pg := pagination.Setup(page, size, int(total))

	if err := database.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Offset(pg.Offset).Limit(pg.Size)
	}).Find(&jobs).Error; err != nil {
		return nil, nil, err
	}

	return jobs, pagination.PaginationResponse(pg.Page, pg.Size, pg.TotalPage, pg.TotalSize), nil
}
func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{
		db: db,
	}
}
