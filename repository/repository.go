package repository

import (
	"mockery/abstraction"

	"gorm.io/gorm"
)

//go:generate mockery.bat --name=AnyRepository --output=mocks --with-expecter
type AnyRepository interface {
	GetName(ctx *abstraction.Context, id int) (string, error)
	Find(ctx *abstraction.Context, f *abstraction.Query, model interface{}) (*abstraction.PaginationInfo, error)
	Last(ctx *abstraction.Context, f *abstraction.Query, model any) (any, error)
	FindById(ctx *abstraction.Context, id int, f *abstraction.Query, model any) (any, error)
	Create(ctx *abstraction.Context, model any, opts ...func(*gorm.DB) *gorm.DB) (any, error)
	Update(ctx *abstraction.Context, id int, model any) (any, error)
	SetNil(ctx *abstraction.Context, model any, column string) (any, error)
	Delete(ctx *abstraction.Context, model any) (any, error)
	BatchDelete(ctx *abstraction.Context, filter abstraction.Anyfilter, model any) (any, error)
	ExecRawSQL(ctx *abstraction.Context, query string, args ...interface{}) (any, error)
	Count(ctx *abstraction.Context, f *abstraction.Query, model any, count *int64) (any, error)
	Association(payload string, query *gorm.DB) (*gorm.DB, error)
	UnscopedAssociation(payload string, query *gorm.DB) (*gorm.DB, error)
	Filter(payload string, query *gorm.DB) (*gorm.DB, error)
	LiterallyCount(ctx *abstraction.Context, f *abstraction.Query, m any, count *int64) (any, error)
	Save(ctx *abstraction.Context, model any) (any, error)
	HardDelete(ctx *abstraction.Context, f *abstraction.Query, model any) (any, error)
}
