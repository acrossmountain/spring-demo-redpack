package test

import (
	"github.com/acrossmounation/redpack/core/envelopes"
	"github.com/go-spring/spring-boot"
)

// goods dao test
func GoodsDaoTest() {
	SpringBoot.RegisterBean(new(envelopes.TestGoodsDao))
}
