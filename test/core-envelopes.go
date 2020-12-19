package test

import (
	"github.com/acrossmounation/redpack/core/envelopes"
	"github.com/go-spring/spring-boot"
)

// goods dao test
func EnvelopeGoodsDaoTest() {
	SpringBoot.RegisterBean(new(envelopes.TestRedEnvelopeGoodsDao))
}

func EnvelopeGoodsItemDaoTest() {
	SpringBoot.RegisterBean(new(envelopes.TestRedEnvelopeGoodsItemDao))
}
