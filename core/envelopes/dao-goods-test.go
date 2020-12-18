package envelopes

import (
	"database/sql"
	"testing"
	"time"

	"github.com/acrossmounation/redpack/services"

	"github.com/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

type TestGoodsDao struct {
	_  SpringBoot.JUnitSuite `export:""`
	Db *gorm.DB              `autowire:""`
}

func (am *TestGoodsDao) Test(t *testing.T) {
	err := am.Db.Transaction(func(tx *gorm.DB) error {

		dao := RedEnvelopeGoodsDao{runner: tx}
		amount := decimal.NewFromFloat(100)
		quantity := 10
		Convey("红包数据写入", t, func() {
			po := &RedEnvelopeGoods{
				EnvelopeNo:   ksuid.New().Next().String(),
				EnvelopeType: services.GeneralEnvelopeType,
				Username: sql.NullString{
					String: "Qyq",
				},
				UserId: ksuid.New().Next().String(),
				Blessing: sql.NullString{
					String: services.DefaultBlessing,
				},
				Amount:         amount,
				AmountOne:      decimal.NewFromFloat(100),
				Quantity:       quantity,
				RemainAmount:   amount,
				RemainQuantity: quantity,
				ExpiredAt:      time.Now().Add(time.Hour),
				Status:         services.OrderCreate,
				OrderType:      services.OrderTypeSending,
				PayStatus:      services.Payed,
			}

			id, err := dao.Insert(po)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			envelope := dao.GetOne(po.EnvelopeNo)
			So(envelope, ShouldNotBeNil)
			So(envelope.EnvelopeNo, ShouldEqual, po.EnvelopeNo)

			Convey("扣减红包余额", func() {
				sub := decimal.NewFromFloat(10)
				rows, err := dao.UpdateBalance(envelope.EnvelopeNo, sub)
				So(err, ShouldBeNil)
				So(rows, ShouldBeGreaterThan, 0)

				envelope := dao.GetOne(po.EnvelopeNo)
				So(envelope, ShouldNotBeNil)
				So(envelope.EnvelopeNo, ShouldEqual, po.EnvelopeNo)
				So(envelope.RemainAmount.String(), ShouldEqual, po.RemainAmount.Sub(sub).String())
				So(envelope.RemainQuantity, ShouldEqual, po.RemainQuantity-1)
			})

			Convey("修改红包状态", func() {
				rows, err := dao.UpdateOrderStatus(envelope.EnvelopeNo, services.OrderDisabled)
				So(err, ShouldBeNil)
				So(rows, ShouldBeGreaterThan, 0)

				envelope := dao.GetOne(po.EnvelopeNo)
				So(envelope, ShouldNotBeNil)
				So(envelope.EnvelopeNo, ShouldEqual, po.EnvelopeNo)
				So(envelope.Status, ShouldEqual, services.OrderDisabled)
			})
		})

		return nil
	})

	if err != nil {
		t.Log(err.Error())
	}
}
