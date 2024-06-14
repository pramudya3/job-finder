package usecase

import (
	"context"
	model "job-finder/internal/company/model"
	"job-finder/internal/company/repository"
	"job-finder/pkg/config"
	"time"

	"github.com/google/uuid"
)

//go:generate mockery --name=CompanyUsecase
type CompanyUsecase interface {
	Register(ctx context.Context, company *model.CreateCompany) (*model.CompanyResponse, error)
	Delete(ctx context.Context, companyID uuid.UUID) error
	FindByID(ctx context.Context, companyID uuid.UUID) (*model.Company, error)
}

type companyUsecase struct {
	companyRepository repository.CompanyRepository
	timeout           time.Duration
}

func (c *companyUsecase) Register(ctx context.Context, company *model.CreateCompany) (*model.CompanyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	newCompany := &model.Company{
		ID:   uuid.New(),
		Name: company.Name,
	}
	if err := c.companyRepository.Create(ctx, newCompany); err != nil {
		return nil, err
	}

	return &model.CompanyResponse{
		ID:   newCompany.ID,
		Name: newCompany.Name,
	}, nil
}

func (c *companyUsecase) Delete(ctx context.Context, companyID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return c.companyRepository.Delete(ctx, companyID)
}

func (c *companyUsecase) FindByID(ctx context.Context, companyID uuid.UUID) (*model.Company, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return c.companyRepository.FindByID(ctx, companyID)
}

func NewCompanyUsecase(companyRepo repository.CompanyRepository, cfg *config.Config) CompanyUsecase {
	return &companyUsecase{
		companyRepository: companyRepo,
		timeout:           time.Duration(cfg.CtxTimeout) * time.Second,
	}
}
