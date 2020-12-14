package main

import (
	_ "github.com/acrossmounation/redpack/apis"
	_ "github.com/acrossmounation/redpack/core/accounts"

	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
	_ "github.com/go-spring/starter-echo"
	_ "github.com/go-spring/starter-gorm/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	SpringBoot.Config(func(db *gorm.DB) {
		db.SingularTable(true)
	})
	SpringWeb.Validator = SpringWeb.NewDefaultValidator()
	SpringBoot.RunApplication()
}
