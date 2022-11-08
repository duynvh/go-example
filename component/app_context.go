package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
}

type appCtx struct {
	db     *gorm.DB
	secret string
}

func NewAppContext(db *gorm.DB, secretKey string) *appCtx {
	return &appCtx{db: db, secret: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db.Debug()
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secret
}
