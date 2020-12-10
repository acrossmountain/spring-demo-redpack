package accounts

import (
	"database/sql"
	"testing"

	"github.com/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

// 测试 根据账户编号获取账号信息
type TestAccountDaoInsert struct {
	_  SpringBoot.JUnitSuite `export:""`
	Db *gorm.DB              `autowire:""`
}

func (am *TestAccountDaoInsert) Test(t *testing.T) {

	err := am.Db.Transaction(func(tx *gorm.DB) error {

		dao := AccountDao{runner: tx}

		Convey("查询账户", t, func() {
			a := &Account{
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户",
				Username: sql.NullString{
					String: "测试用户",
					Valid:  true,
				},
			}
			insert, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(insert, ShouldBeGreaterThan, 0)

			na := dao.GetOne(a.AccountNo)
			So(na, ShouldNotBeNil)
			So(na.AccountNo, ShouldEqual, a.AccountNo)
			So(na.CreatedAt, ShouldNotBeNil)
			So(na.UpdatedAt, ShouldNotBeNil)
		})

		return nil
	})

	if err != nil {
		t.Log(err.Error())
	}
}

// 测试 根据用户ID获取账户信息
type TestAccountDaoGetUserById struct {
	_  SpringBoot.JUnitSuite `export:""`
	Db *gorm.DB              `autowire:""`
}

func (am *TestAccountDaoGetUserById) Test(t *testing.T) {

	err := am.Db.Transaction(func(tx *gorm.DB) error {

		dao := AccountDao{runner: tx}

		Convey("通过ID查询账户", t, func() {
			a := &Account{
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户",
				UserId:      ksuid.New().Next().String(),
				Username: sql.NullString{
					String: "测试用户",
					Valid:  true,
				},
				AccountType: 2,
			}
			insert, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(insert, ShouldBeGreaterThan, 0)

			na := dao.GetUserById(a.UserId, 2)
			So(na, ShouldNotBeNil)
			So(na.UserId, ShouldEqual, a.UserId)
			So(na.CreatedAt, ShouldNotBeNil)
			So(na.UpdatedAt, ShouldNotBeNil)
		})

		return nil
	})

	if err != nil {
		t.Log(err.Error())
	}
}

// 测试 用户余额
type TestAccountDaoUpdateBalance struct {
	_  SpringBoot.JUnitSuite `export:""`
	Db *gorm.DB              `autowire:""`
}

func (am *TestAccountDaoUpdateBalance) Test(t *testing.T) {
	err := am.Db.Transaction(func(tx *gorm.DB) error {

		dao := AccountDao{runner: tx}
		balance := decimal.NewFromFloat(100)

		Convey("更新账户余额", t, func() {

			a := &Account{
				Balance:     balance,
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户",
				UserId:      ksuid.New().Next().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
			}
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			// 1. 增加余额
			Convey("增加余额", func() {
				amount := decimal.NewFromFloat(10)
				rows, err := dao.UpdateBalance(a.AccountNo, amount)
				So(err, ShouldBeNil)
				So(rows, ShouldEqual, 1)
				na := dao.GetOne(a.AccountNo)
				newBalance := balance.Add(amount)
				So(na, ShouldNotBeNil)
				So(na.Balance.String(), ShouldEqual, newBalance.String())
			})

			// 2. 扣减余额，余额足够
			Convey("扣减余额，余额足够", func() {
				amount := decimal.NewFromFloat(-10)
				rows, err := dao.UpdateBalance(a.AccountNo, amount)
				So(err, ShouldBeNil)
				So(rows, ShouldEqual, 1)
				na := dao.GetOne(a.AccountNo)
				newBalance := balance.Add(amount)
				So(na, ShouldNotBeNil)
				So(na.Balance.String(), ShouldEqual, newBalance.String())
			})

			// 3. 扣减余额，余额不够
			Convey("扣减余额，余额不够", func() {
				a1 := dao.GetOne(a.AccountNo)
				So(a1, ShouldNotBeNil)
				amount := decimal.NewFromFloat(-300)
				rows, err := dao.UpdateBalance(a.AccountNo, amount)
				So(err, ShouldBeNil)
				So(rows, ShouldEqual, 0)
				a2 := dao.GetOne(a.AccountNo)
				So(a2, ShouldNotBeNil)
				So(a1.Balance.String(), ShouldEqual, a2.Balance.String())
			})
		})

		return nil
	})

	if err != nil {
		t.Log(err.Error())
	}
}
