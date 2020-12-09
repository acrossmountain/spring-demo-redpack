package test

import (
	"github.com/acrossmounation/redpack/core/accounts"
	"github.com/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
)

// config
func config() {
	SpringBoot.Config(func(db *gorm.DB) {
		db.SingularTable(true)
		db.LogMode(true)
		db.AutoMigrate(&accounts.Account{})
	})
}
