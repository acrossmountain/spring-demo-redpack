package test

import (
	"github.com/acrossmounation/redpack/core/accounts"

	"github.com/go-spring/spring-boot"
)

// account dao test
func AccountDaoTest() {
	SpringBoot.RegisterBean(new(accounts.TestAccountDaoInsert))
	SpringBoot.RegisterBean(new(accounts.TestAccountDaoGetUserById))
	SpringBoot.RegisterBean(new(accounts.TestAccountDaoUpdateBalance))
}

// account dao log test
func AccountLogDaoTest() {
	SpringBoot.RegisterBean(new(accounts.TestAccountLogDao))
}

// account domain test
func AccountLogDomainTest() {
	SpringBoot.RegisterBean(new(accounts.TestAccountDomain))
	SpringBoot.RegisterBean(new(accounts.TestAccountDomainTransfer))
}

// account service test
func AccountServiceTest() {
	SpringBoot.RegisterBean(new(accounts.TestAccountServiceCreate))
	SpringBoot.RegisterBean(new(accounts.TestAccountServiceTransfer))
}
