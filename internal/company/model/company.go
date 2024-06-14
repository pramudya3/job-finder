package model

import (
	"job-finder/internal/job/model"

	"github.com/google/uuid"
)

type (
	Company struct {
		ID   uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		Name string    `json:"name" gorm:"unique"`

		Jobs []*model.Job `json:"jobs" gorm:"foreignKey:CompanyID"`
	}

	CreateCompany struct {
		Name string `json:"name" validate:"required"`
	}

	CompanyResponse struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}
)
