package abstraction

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const (
	defaultPageSize = 100
)

var (
	defaultSortBy  = "modified_at"
	defaultOrderBy = "desc"
)

type PaginationCursor struct {
	PageSize int     `json:"page_size" query:"page_size" default:"100"`
	Cursor   string  `json:"cursor"    query:"cursor"`
	SortBy   *string `json:"sort_by"   query:"sort_by"                 example:"id"`
	OrderBy  *string `json:"order_by"  query:"order_by"                             enums:"asc,desc"`
}

func (p *PaginationCursor) Limit() int {
	if p == nil || p.PageSize <= 0 || p.PageSize > defaultPageSize {
		return defaultPageSize
	}

	return p.PageSize
}

func (p *PaginationCursor) GetPageSize() int {
	if p == nil || p.PageSize <= 0 || p.PageSize > defaultPageSize {
		return defaultPageSize
	}
	return p.PageSize
}

type Pagination struct {
	Page     int     `json:"page" query:"page" default:"1"`
	PageSize int     `json:"page_size" query:"page_size" default:"100"`
	Cursor   string  `json:"cursor" query:"cursor"`
	SortBy   *string `json:"sort_by" query:"sort_by" example:"id"`
	OrderBy  *string `json:"order_by" query:"order_by" enums:"asc,desc"`
}

// set default before further processing
// do it after binding the pagination object in handler layer
func (p *Pagination) SetDefault() {
	if p == nil || p.PageSize <= 0 {
		p.PageSize = defaultPageSize
	}
	if p == nil || p.Page <= 0 {
		p.Page = 1
	}

	var sort string

	if p == nil {
		return
	}

	if IsStringBlank(p.SortBy) {
		p.SortBy = &defaultSortBy
	}

	if *p.SortBy == "order" {
		*p.SortBy = fmt.Sprintf("\"%s\"", "order")
	}

	if IsStringBlank(p.OrderBy) {
		p.OrderBy = &defaultOrderBy
	}

	switch *p.OrderBy {
	case "asc":
		sort = "asc"
	case "desc":
		sort = "desc"
	default:
		sort = "desc"
	}
	p.OrderBy = &sort
}

func (p *Pagination) Limit() int {
	if p == nil || p.PageSize <= 0 {
		return defaultPageSize
	}
	return p.PageSize
}

func (p *Pagination) Offset() int {
	if p == nil || p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit()
}

func (p *Pagination) GetPage() int {
	if p == nil || p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p *Pagination) GetPageSize() int {
	if p == nil || p.PageSize <= 0 {
		return defaultPageSize
	}
	return p.PageSize
}

func (p *Pagination) Apply(db *gorm.DB) {
	if p != nil {
		db.Limit(p.Limit())
		db.Offset(p.Offset())
		// if p.OrderBy != nil {
		db.Order(p.GetSortBy())
		// }
	}
}

func IsStringBlank(s *string) bool {
	if s == nil || strings.TrimSpace(*s) == "" {
		return true
	}
	return false
}

func (p *Pagination) ChangeDefaultSortingClause(sortBy string, orderBy *string) {
	// because this function is to populate the sorting
	// sortBy can't be empty string
	// orderBy might be empty
	// defaultorderBy := "desc"
	if sortBy == "" {
		return
	}
	if IsStringBlank(p.SortBy) {
		p.SortBy = &defaultSortBy
	}

	if *p.SortBy == defaultSortBy {
		p.SortBy = &sortBy
	}

	if !IsStringBlank(orderBy) && IsStringBlank(p.OrderBy) {
		p.OrderBy = orderBy
	}
}

func (p *Pagination) SetDefaultSorting() {
	var sort string

	if p == nil {
		return
	}

	if IsStringBlank(p.SortBy) {
		p.SortBy = &defaultSortBy
	}

	if *p.SortBy == "order" {
		*p.SortBy = fmt.Sprintf("\"%s\"", "order")
	}

	if IsStringBlank(p.OrderBy) {
		p.OrderBy = &defaultOrderBy
	}

	switch *p.OrderBy {
	case "asc":
		sort = "asc"
	case "desc":
		sort = "desc"
	default:
		sort = "desc"
	}
	p.OrderBy = &sort
}

// Return SortBy value if not nil with OrderBy value with default asc
func (p *Pagination) GetSortBy() string {
	return *p.SortBy + " " + *p.OrderBy
}

func (p *Pagination) GetSorting() *Sorting {
	if IsStringBlank(p.OrderBy) {
		p.OrderBy = &defaultOrderBy
	}
	if p.SortBy != nil {
		return &Sorting{
			*p.SortBy,
			*p.OrderBy,
		}
	}

	return nil
}

type Sorting struct {
	SortBy  string `query:"sort_by"`
	OrderBy string `query:"order_by" enums:"asc,desc" default:"asc"`
}

func NewSorting(sortBy, sort string) *Sorting {
	if sort != "desc" {
		sort = "asc"
	}
	return &Sorting{
		SortBy:  sortBy,
		OrderBy: sort,
	}
}

type PaginationInfo struct {
	*Pagination
	*Sorting
	Count         int    `json:"count"`
	MoreRecords   bool   `json:"more_records"`
	NextCursor    string `json:"next_cursor"`
	TotalPageSize int    `json:"total_page"`
	TotalCount    int    `json:"total_count,omitempty"`
}

func (p *Pagination) CreatePageInfo(count int64) *PaginationInfo {
	p.SetDefault()
	return &PaginationInfo{
		Pagination:    p,
		MoreRecords:   false,
		Count:         int(count),
		TotalPageSize: int((count + int64(p.PageSize) - 1) / int64(p.PageSize)),
	}
}
