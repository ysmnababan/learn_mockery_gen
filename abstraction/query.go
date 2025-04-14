package abstraction

import "gorm.io/gorm"

type Anyfilter interface {
	Applys(db *gorm.DB)
}

type AnyJoin interface {
	Join(db *gorm.DB)
}

type Query struct {
	Associations         string `query:"associations"         json:"associations"`
	UnscopedAssociations string `query:"unscopedassociations" json:"unscopedassociations"`
	Filter               Anyfilter
	InternalFilter       string
	Countfilter          string
	Join                 AnyJoin
	SkipCount            bool
	DisableCountWhere    bool
	Pagination
	Cursor *PaginationCursor
	Hook   func(db *gorm.DB) *gorm.DB
}
