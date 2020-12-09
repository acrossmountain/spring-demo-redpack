package test

import (
	"github.com/acrossmounation/redpack/core/accounts"

	"github.com/go-spring/spring-boot"
)

// account dao test
func AccountDaoTest() {
	SpringBoot.RegisterBean(new(accounts.TestAccountDaoInsert))
	SpringBoot.RegisterBean(new(accounts.TestAccountDaoGetUserById))
	SpringBoot.RegisterBean(new(accounts.TestAccountDaoUploadBalance))
}
