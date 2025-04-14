package abstraction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagination_Limit(t *testing.T) {
	type fields struct {
		Page     int
		PageSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "PageSize = 0",
			fields: fields{
				PageSize: 0,
			},
			want: defaultPageSize,
		},
		{
			name: "PageSize = 1",
			fields: fields{
				PageSize: 1,
			},
			want: 1,
		},
		{
			name: "PageSize = 15",
			fields: fields{
				PageSize: 7,
			},
			want: 7,
		},
		{
			name: "PageSize = negative",
			fields: fields{
				PageSize: -12,
			},
			want: defaultPageSize,
		},
		{
			name: "PageSize = some big number",
			fields: fields{
				PageSize: 1000,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				Page:     tt.fields.Page,
				PageSize: tt.fields.PageSize,
			}
			if got := p.Limit(); got != tt.want {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_Offset(t *testing.T) {
	type fields struct {
		Page     int
		PageSize int
	}
	tests := []struct {
		name string
		data fields
		want int
	}{
		{
			name: "page = neg and pageSize = neg",
			data: fields{
				Page:     -1,
				PageSize: -1,
			},
			want: 0,
		},
		{
			name: "page = 0 and pageSize = neg",
			data: fields{
				Page:     0,
				PageSize: -1,
			},
			want: 0,
		},
		{
			name: "page = pos and pageSize = neg",
			data: fields{
				Page:     3,
				PageSize: -1,
			},
			want: 200,
		},
		{
			name: "page = neg and pageSize = zero",
			data: fields{
				Page:     -1,
				PageSize: 0,
			},
			want: 0,
		},
		{
			name: "page = 0 and pageSize = zero",
			data: fields{
				Page:     0,
				PageSize: 0,
			},
			want: 0,
		},
		{
			name: "page = pos and pageSize = zero",
			data: fields{
				Page:     3,
				PageSize: 0,
			},
			want: 200,
		},
		{
			name: "page = neg and pageSize = pos",
			data: fields{
				Page:     -1,
				PageSize: 4,
			},
			want: 0,
		},
		{
			name: "page = 0 and pageSize = pos",
			data: fields{
				Page:     0,
				PageSize: 4,
			},
			want: 0,
		},
		{
			name: "page = pos and pageSize = pos",
			data: fields{
				Page:     3,
				PageSize: 4,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				Page:     tt.data.Page,
				PageSize: tt.data.PageSize,
			}
			if got := p.Offset(); got != tt.want {
				t.Errorf("Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_SetDefault_pagetest(t *testing.T) {
	type fields struct {
		Page     int
		PageSize int
	}
	test := []struct {
		name string
		data fields
		want fields
	}{
		{
			name: "page = neg and pageSize = neg",
			data: fields{
				Page:     -1,
				PageSize: -1,
			},
			want: fields{
				Page:     1,
				PageSize: defaultPageSize,
			},
		},
		{
			name: "page = 0 and pageSize = neg",
			data: fields{
				Page:     0,
				PageSize: -1,
			},
			want: fields{
				Page:     1,
				PageSize: defaultPageSize,
			},
		},
		{
			name: "page = pos and pageSize = neg",
			data: fields{
				Page:     3,
				PageSize: -1,
			},
			want: fields{
				Page:     3,
				PageSize: defaultPageSize,
			},
		},
		{
			name: "page = neg and pageSize = zero",
			data: fields{
				Page:     -1,
				PageSize: 0,
			},
			want: fields{
				Page:     1,
				PageSize: defaultPageSize,
			},
		},
		{
			name: "page = 0 and pageSize = zero",
			data: fields{
				Page:     0,
				PageSize: 0,
			},
			want: fields{
				Page:     1,
				PageSize: defaultPageSize,
			},
		},
		{
			name: "page = pos and pageSize = zero",
			data: fields{
				Page:     3,
				PageSize: 0,
			},
			want: fields{
				Page:     3,
				PageSize: defaultPageSize,
			},
		},
		{
			name: "page = neg and pageSize = pos",
			data: fields{
				Page:     -1,
				PageSize: 4,
			},
			want: fields{
				Page:     1,
				PageSize: 4,
			},
		},
		{
			name: "page = 0 and pageSize = pos",
			data: fields{
				Page:     0,
				PageSize: 4,
			},
			want: fields{
				Page:     1,
				PageSize: 4,
			},
		},
		{
			name: "page = pos and pageSize = pos",
			data: fields{
				Page:     3,
				PageSize: 4,
			},
			want: fields{
				Page:     3,
				PageSize: 4,
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				Page:     tt.data.Page,
				PageSize: tt.data.PageSize,
			}
			p.SetDefault()
			if p.Page != tt.want.Page || p.PageSize != tt.want.PageSize {
				t.Errorf("PAGE: get  %d want %d; PAGESIZE: get %d want %d",
					p.Page, tt.want.Page, p.PageSize, tt.want.PageSize)
			}
		})
	}
}

func TestPagiation_IsStringBlank(t *testing.T) {
	// map
	type field struct {
		name string
		data *string
		want bool
	}
	emptystring := ""
	notemptystring := "not empty"
	tests := []field{
		{
			name: "nil",
			want: true,
		},
		{
			name: "empty",
			data: &emptystring,
			want: true,
		},
		{
			name: "not empty",
			data: &notemptystring,
			want: false,
		},
	}

	for _, val := range tests {
		t.Run(val.name, func(t *testing.T) {
			if got := IsStringBlank(val.data); got != val.want {
				t.Errorf("get: %v, want %v, data %v", got, val.want, val.data)
			}
		})
	}
}

func TestPagination_GetSortBy(t *testing.T) {
	// combination of sortBy and orderBy field
	// each can be nil, empty string, and ordinary string

	// var nullString *string 
	emptyString := ""
	normalString := "normalString"
	randomString := "randomString"
	asc := "asc"
	desc := "desc"

	type field struct {
		name    string
		sortBy  *string
		orderBy *string
		want    string
	}

	tests := []field{
		{
			name:    "",
			want:    defaultSortBy + " " + desc,
			orderBy: &randomString,
			sortBy:  nil,
		},
		{
			name:    "",
			want:    "modified_at" + " " + asc,
			orderBy: &asc,
			sortBy:  nil,
		},
		{
			name:    "",
			want:    "modified_at" + " " + asc,
			orderBy: &asc,
			sortBy:  &emptyString,
		},
		{
			name:    "",
			want:    normalString + " " + asc,
			orderBy: &asc,
			sortBy:  &normalString,
		},

		{
			name:    "",
			want:    "modified_at" + " " + desc,
			orderBy: &emptyString,
			sortBy:  nil,
		},
		{
			name:    "",
			want:    "modified_at" + " " + desc,
			orderBy: &emptyString,
			sortBy:  &emptyString,
		},
		{
			name:    "",
			want:    normalString + " " + desc,
			orderBy: &emptyString,
			sortBy:  &normalString,
		},

		{
			name:    "",
			want:    "modified_at" + " " + desc,
			orderBy: nil,
			sortBy:  nil,
		},
		{
			name:    "",
			want:    "modified_at" + " " + desc,
			orderBy: nil,
			sortBy:  &emptyString,
		},
		{
			name:    "",
			want:    normalString + " " + desc,
			orderBy: nil,
			sortBy:  &normalString,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := new(Pagination)
			p.OrderBy = test.orderBy
			p.SortBy = test.sortBy
			p.SetDefault()
			got := p.GetSortBy()
			if got != test.want {
				t.Errorf("get %s, want %s", got, test.want)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}


func TestChangeDefaultSortingClause_Updated(t *testing.T) {
	tests := []struct {
		name         string
		initialSort  *string
		initialOrder *string
		sortBy       string
		orderBy      *string
		wantSort     *string
		wantOrder    *string
	}{
		{
			name:         "Sets to defaultSortBy then overrides with sortBy",
			initialSort:  nil,
			initialOrder: nil,
			sortBy:       "created_at",
			orderBy:      strPtr("asc"),
			wantSort:     strPtr("created_at"),
			wantOrder:    strPtr("asc"),
		},
		{
			name:         "Change SortBy when it equals defaultSortBy",
			initialSort:  strPtr(defaultSortBy),
			initialOrder: nil,
			sortBy:       "created_at",
			orderBy:      nil,
			wantSort:     strPtr("created_at"),
			wantOrder:    nil,
		},
		{
			name:         "Does not change sortBy if already set and not default",
			initialSort:  strPtr("email"),
			initialOrder: nil,
			sortBy:       "created_at",
			orderBy:      strPtr("desc"),
			wantSort:     strPtr("email"),
			wantOrder:    strPtr("desc"),
		},
		{
			name:         "Sets default first, then applies sortBy if default",
			initialSort:  nil,
			initialOrder: strPtr("asc"),
			sortBy:       "name",
			orderBy:      nil,
			wantSort:     strPtr("name"),
			wantOrder:    strPtr("asc"),
		},
		{
			name:         "Skips everything when sortBy is empty",
			initialSort:  nil,
			initialOrder: nil,
			sortBy:       "",
			orderBy:      strPtr("desc"),
			wantSort:     nil,
			wantOrder:    nil,
		},
		{
			name:         "Does not override OrderBy if already set",
			initialSort:  strPtr(defaultSortBy),
			initialOrder: strPtr("asc"),
			sortBy:       "updated_at",
			orderBy:      strPtr("desc"),
			wantSort:     strPtr("updated_at"),
			wantOrder:    strPtr("asc"),
		},
		{
			name:         "Skips OrderBy if orderBy is blank",
			initialSort:  strPtr(defaultSortBy),
			initialOrder: nil,
			sortBy:       "email",
			orderBy:      strPtr(" "),
			wantSort:     strPtr("email"),
			wantOrder:    nil,
		},
		{
			name:         "Only changes OrderBy if it was blank",
			initialSort:  strPtr("name"),
			initialOrder: nil,
			sortBy:       "email",
			orderBy:      strPtr("asc"),
			wantSort:     strPtr("name"), // unchanged
			wantOrder:    strPtr("asc"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				SortBy:  tt.initialSort,
				OrderBy: tt.initialOrder,
			}

			p.ChangeDefaultSortingClause(tt.sortBy, tt.orderBy)

			if tt.wantSort == nil {
				assert.Nil(t, p.SortBy)
			} else {
				assert.NotNil(t, p.SortBy)
				assert.Equal(t, *tt.wantSort, *p.SortBy)
			}

			if tt.wantOrder == nil {
				assert.Nil(t, p.OrderBy)
			} else {
				assert.NotNil(t, p.OrderBy)
				assert.Equal(t, *tt.wantOrder, *p.OrderBy)
			}
		})
	}
}