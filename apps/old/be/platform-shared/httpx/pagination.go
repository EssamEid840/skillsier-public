package httpx

import (
	"net/http"
	"strconv"
)

// PaginationParams represents pagination parameters from query string
type PaginationParams struct {
	// Page is the current page number (1-based)
	Page int
	
	// Limit is the number of items per page
	Limit int
	
	// Offset is calculated from Page and Limit
	Offset int
	
	// SortBy is the field to sort by
	SortBy string
	
	// SortOrder is the sort direction (asc, desc)
	SortOrder string
}

// PaginationMeta represents pagination metadata in responses
type PaginationMeta struct {
	// CurrentPage is the current page number
	CurrentPage int `json:"current_page"`
	
	// PerPage is the number of items per page
	PerPage int `json:"per_page"`
	
	// Total is the total number of items
	Total int64 `json:"total"`
	
	// TotalPages is the total number of pages
	TotalPages int `json:"total_pages"`
	
	// HasNext indicates if there's a next page
	HasNext bool `json:"has_next"`
	
	// HasPrev indicates if there's a previous page
	HasPrev bool `json:"has_prev"`
}

// DefaultPaginationParams returns default pagination parameters
func DefaultPaginationParams() *PaginationParams {
	return &PaginationParams{
		Page:      1,
		Limit:     20,
		Offset:    0,
		SortBy:    "created_at",
		SortOrder: "desc",
	}
}

// ParsePaginationParams extracts pagination parameters from HTTP request
func ParsePaginationParams(r *http.Request) *PaginationParams {
	params := DefaultPaginationParams()
	
	// Parse page
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			params.Page = page
		}
	}
	
	// Parse limit
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			params.Limit = limit
			// Cap at 100 items per page
			if params.Limit > 100 {
				params.Limit = 100
			}
		}
	}
	
	// Calculate offset
	params.Offset = (params.Page - 1) * params.Limit
	
	// Parse sort_by
	if sortBy := r.URL.Query().Get("sort_by"); sortBy != "" {
		params.SortBy = sortBy
	}
	
	// Parse sort_order
	if sortOrder := r.URL.Query().Get("sort_order"); sortOrder != "" {
		if sortOrder == "asc" || sortOrder == "desc" {
			params.SortOrder = sortOrder
		}
	}
	
	return params
}

// NewPaginationMeta creates pagination metadata
func NewPaginationMeta(page, limit int, total int64) *PaginationMeta {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	if totalPages < 1 {
		totalPages = 1
	}
	
	return &PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrev:     page > 1,
	}
}

// WritePaginatedResponse writes a paginated JSON response
func WritePaginatedResponse(w http.ResponseWriter, r *http.Request, data interface{}, params *PaginationParams, total int64) {
	meta := NewPaginationMeta(params.Page, params.Limit, total)
	WriteSuccessWithMeta(w, r, data, meta)
}