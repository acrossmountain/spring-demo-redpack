package test

import (
	"testing"

	"github.com/go-spring/spring-boot"
	_ "github.com/go-spring/starter-gorm/mysql"
)

func TestEntry(t *testing.T) {
	AccountDaoTest()
	AccountLogDaoTest()
	AccountLogDomainTest()
	AccountServiceTest()
	config()
	SpringBoot.RunTestApplication(t, 0, "../config")
}
