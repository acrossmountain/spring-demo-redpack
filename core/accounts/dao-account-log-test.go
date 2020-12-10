package accounts

import (
	"testing"

	"github.com/acrossmounation/redpack/services"
	"github.com/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

type TestAccountLogDao struct {
	_  SpringBoot.JUnitSuite `export:""`
	Db *gorm.DB              `autowire:""`
}

func (am *TestAccountLogDao) Test(t *testing.T) {
	err := am.Db.Transaction(func(tx *gorm.DB) error {

		dao := AccountLogDao{runner: tx}

		Convey("查询用户流水日志", t, func() {
			a := &AccountLog{
				LogNo:      ksuid.New().Next().String(),
				TradeNo:    ksuid.New().Next().String(),
				AccountNo:  ksuid.New().Next().String(),
				UserId:     ksuid.New().Next().String(),
				Status:     1,
				Username:   "测试用户",
				Amount:     decimal.NewFromFloat(1),
				Balance:    decimal.NewFromFloat(100),
				ChangeFlag: services.ChangeFlagAccountCreated,
				ChangeType: services.ChangeTypeAccountCreated,
			}
			insert, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(insert, ShouldBeGreaterThan, 0)

			Convey("通过日志编号查询", func() {
				na := dao.GetOne(a.LogNo)
				So(na, ShouldNotBeNil)
				So(na.LogNo, ShouldEqual, a.LogNo)
				So(na.Balance.String(), ShouldEqual, a.Balance.String())
				So(na.Amount.String(), ShouldEqual, a.Amount.String())
				So(na.CreatedAt, ShouldNotBeNil)
			})

			Convey("通过交易编号查询", func() {
				na := dao.GetByTradeNo(a.TradeNo)
				So(na, ShouldNotBeNil)
				So(na.LogNo, ShouldEqual, a.LogNo)
				So(na.Balance.String(), ShouldEqual, a.Balance.String())
				So(na.Amount.String(), ShouldEqual, a.Amount.String())
				So(na.CreatedAt, ShouldNotBeNil)
			})

		})

		return nil
	})

	if err != nil {
		t.Log(err.Error())
	}
}
