package api

import (
	"employe/pkg/store"
	"math"
)

type PaginatedEmployees struct {
	Data       []store.Employee `json:"data"`
	Total      int              `json:"total"`
	Page       int              `json:"page"`
	PageSize   int              `json:"pageSize"`
	TotalPages int              `json:"totalPages"`
}

func paginateEmployees(employees []store.Employee, page, pageSize int) PaginatedEmployees {
	total := len(employees)
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	start := (page - 1) * pageSize
	if start > total {
		start = total
	}

	end := start + pageSize
	if end > total {
		end = total
	}

	paginated := PaginatedEmployees{
		Data:       employees[start:end],
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
	return paginated
}
