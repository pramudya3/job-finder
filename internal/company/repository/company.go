package repository

import (
	"context"
	model "job-finder/internal/company/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//go:generate mockery --name=CompanyRepository
type CompanyRepository interface {
	Create(ctx context.Context, company *model.Company) error
	Delete(ctx context.Context, companyID uuid.UUID) error
	FindByID(ctx context.Context, companyID uuid.UUID) (*model.Company, error)
}

type companyRepository struct {
	db *gorm.DB
}

func (c *companyRepository) Create(ctx context.Context, company *model.Company) error {
	c.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&company).Error
	})
	return nil
}

func (c *companyRepository) Delete(ctx context.Context, companyID uuid.UUID) error {
	c.db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", companyID).Delete(&model.Company{}).Error
	})

	return nil
}

func (c *companyRepository) FindByID(ctx context.Context, companyID uuid.UUID) (*model.Company, error) {
	company := &model.Company{}
	if err := c.db.Preload("Jobs").First(&company, "id = ?", companyID).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{
		db: db,
	}
}
