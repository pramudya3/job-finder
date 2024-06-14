package model

import (
	"time"

	"github.com/google/uuid"
)

// find job by query from companyName, or title or description (contains from query)

type (
	Job struct {
		ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		CompanyID   uuid.UUID `json:"company_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
	}

	CreateJobReq struct {
		CompanyID   uuid.UUID `json:"company_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
	}
)
