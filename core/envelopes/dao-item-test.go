package envelopes

import (
	"database/sql"
	"testing"

	"github.com/acrossmounation/redpack/services"

	"github.com/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

type TestRedEnvelopeGoodsItemDao struct {
	_  SpringBoot.JUnitSuite `export:""`
	Db *gorm.DB              `autowire:""`
}

func (am *TestRedEnvelopeGoodsItemDao) Test(t *testing.T) {
	err := am.Db.Transaction(func(tx *gorm.DB) error {

		dao := RedEnvelopeGoodsItemDao{runner: tx}

		Convey("收红包测试", t, func() {

			item := RedEnvelopeItem{
				ItemNo:       ksuid.New().Next().String(),
				EnvelopeNo:   ksuid.New().Next().String(),
				RecvUsername: sql.NullString{String: "测试"},
				RecvUserId:   ksuid.New().Next().String(),
				Amount:       decimal.NewFromFloat(10),
				Quantity:     1,
				RemainAmount: decimal.NewFromFloat(10),
				AccountNo:    ksuid.New().Next().String(),
				PayStatus:    int(services.Payed),
			}

			id, err := dao.Insert(&item)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			one := dao.GetOne(item.ItemNo)
			So(one, ShouldNotBeNil)
			So(one.ItemNo, ShouldEqual, item.ItemNo)
		})

		return nil
	})

	if err != nil {
		t.Log(err)
	}
}
