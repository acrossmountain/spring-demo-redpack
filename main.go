package main

import (
	_ "github.com/acrossmounation/redpack/controllers"

	"github.com/go-spring/spring-boot"
	_ "github.com/go-spring/starter-echo"
	_ "github.com/go-spring/starter-gorm/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	SpringBoot.Config(func(db *gorm.DB) {
		db.SingularTable(true)
	})

	SpringBoot.RunApplication()
}
