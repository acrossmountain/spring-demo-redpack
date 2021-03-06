package test

import (
	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
	"github.com/jinzhu/gorm"
)

// config
func config() {
	SpringBoot.Config(func(db *gorm.DB) {
		db.SingularTable(true)
		db.LogMode(true)
	})
	SpringWeb.Validator = SpringWeb.NewDefaultValidator()
}
