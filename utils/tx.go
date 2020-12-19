package utils

import (
	"context"

	"github.com/go-spring/spring-logger"
	"github.com/jinzhu/gorm"
)

const TX = "tx"

func WithValueContext(parent context.Context, runner *gorm.DB) context.Context {
	return context.WithValue(parent, TX, runner)
}

func ExecuteContext(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx, ok := ctx.Value(TX).(*gorm.DB)
	if !ok || tx == nil {
		SpringLogger.Panic("是否在事务函数块中使用？")
	}
	return fn(tx)
}
