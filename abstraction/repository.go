package abstraction

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type IRepository interface {
	CheckTrx(ctx *Context) *gorm.DB
}

type Repository struct {
	Connection *gorm.DB
	Db         *gorm.DB
	Tx         *gorm.DB
	Redis      *redis.Pool
}

func (r *Repository) CheckTrx(ctx *Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db.WithContext(ctx)
	}
	return r.Db.WithContext(ctx)
}
