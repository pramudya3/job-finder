package pagination

import (
	"math"
)

const (
	DefaultPageSize int64 = 20
)

type (
	Pagination struct {
		Page      int `json:"page"`
		Size      int `json:"size"`
		TotalPage int `json:"total_page"`
		TotalSize int `json:"total_size"`
		Offset    int `json:"offset,omitempty"`
	}

	paginationResponse struct {
		Page      int `json:"page"`
		Size      int `json:"size"`
		TotalPage int `json:"total_page"`
		TotalSize int `json:"total_size"`
	}
)

func Setup(page, size, total int) *Pagination {

	if page < 1 {
		page = 1
	}

	switch {
	case size < 1:
		size = 1
	case size > 100:
		size = 100
	}

	offset := (page - 1) * size

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	return &Pagination{
		Page:      page,
		Size:      size,
		TotalPage: totalPage,
		TotalSize: total,
		Offset:    offset,
	}
}

func PaginationResponse(page, size, totalPage, totalSize int) *paginationResponse {
	return &paginationResponse{
		Page:      page,
		Size:      size,
		TotalPage: totalPage,
		TotalSize: totalSize,
	}
}
