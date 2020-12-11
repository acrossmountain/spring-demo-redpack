package accounts

import (
	"testing"

	"github.com/acrossmounation/redpack/services"

	"github.com/go-spring/spring-boot"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

type TestAccountServiceCreate struct {
	_       SpringBoot.JUnitSuite   `export:""`
	Service services.AccountService `autowire:""`
}

func (am *TestAccountServiceCreate) Test(t *testing.T) {
	dto := services.AccountCreatedDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户",
		Amount:       "100",
		AccountName:  "测试资金账户111",
		AccountType:  int(services.AccountTypeEnvelope),
		CurrencyCode: "CNY",
	}

	Convey("创建账户", t, func() {
		ndto, err := am.Service.CreateAccount(dto)
		So(err, ShouldBeNil)
		So(ndto, ShouldNotBeNil)
		So(ndto.Username, ShouldEqual, dto.Username)
		So(ndto.UserId, ShouldEqual, dto.UserId)
	})
}

type TestAccountServiceTransfer struct {
	_       SpringBoot.JUnitSuite   `export:""`
	Service services.AccountService `autowire:""`
}

func (am *TestAccountServiceTransfer) Test(t *testing.T) {

	dto1 := services.AccountCreatedDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户1",
		Amount:       "100",
		AccountName:  "测试资金账户1",
		AccountType:  int(services.AccountTypeEnvelope),
		CurrencyCode: "CNY",
	}

	dto2 := services.AccountCreatedDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户2",
		Amount:       "100",
		AccountName:  "测试资金账户2",
		AccountType:  int(services.AccountTypeEnvelope),
		CurrencyCode: "CNY",
	}

	Convey("创建测试账户", t, func() {
		account1, err := am.Service.CreateAccount(dto1)
		So(err, ShouldBeNil)
		So(account1, ShouldNotBeNil)
		So(account1.Balance.String(), ShouldEqual, dto1.Amount)
		So(account1.Username, ShouldEqual, dto1.Username)
		So(account1.UserId, ShouldEqual, dto1.UserId)

		account2, err := am.Service.CreateAccount(dto2)
		So(err, ShouldBeNil)
		So(account2, ShouldNotBeNil)
		So(account2.Balance.String(), ShouldEqual, dto2.Amount)
		So(account2.Username, ShouldEqual, dto2.Username)
		So(account2.UserId, ShouldEqual, dto2.UserId)

		// 转账（余额充足）
		Convey("转账（余额充足）", func() {

			amount := decimal.NewFromFloat(1)

			body := services.TradeParticipator{
				AccountNo: account1.AccountNo,
				UserId:    account1.UserId,
				Username:  account1.Username,
			}
			target := services.TradeParticipator{
				AccountNo: account2.AccountNo,
				UserId:    account2.UserId,
				Username:  account2.Username,
			}

			dto := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   body,
				TradeTarget: target,
				AmountStr:   "1",
				ChangeType:  services.ChangeType(-1),
				ChangeFlag:  services.ChangeFlagTransferOut,
				Desc:        "转账",
			}

			status, err := am.Service.Transfer(dto)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, services.TransferStatusSuccess)

			// 验证资金
			no1 := am.Service.GetAccountByNo(account1.AccountNo)
			So(no1, ShouldNotBeNil)
			So(no1.Balance.String(), ShouldEqual, account1.Balance.Sub(amount).String())
			no2 := am.Service.GetAccountByNo(account2.AccountNo)
			So(no2, ShouldNotBeNil)
			So(no2.Balance.String(), ShouldEqual, account2.Balance.Add(amount).String())
		})

		// 转账（余额不足）
		Convey("转账（余额不足）", func() {

			body := services.TradeParticipator{
				AccountNo: account1.AccountNo,
				UserId:    account1.UserId,
				Username:  account1.Username,
			}
			target := services.TradeParticipator{
				AccountNo: account2.AccountNo,
				UserId:    account2.UserId,
				Username:  account2.Username,
			}

			dto := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   body,
				TradeTarget: target,
				AmountStr:   "110",
				ChangeType:  services.ChangeType(-1),
				ChangeFlag:  services.ChangeFlagTransferOut,
				Desc:        "转账",
			}

			status, err := am.Service.Transfer(dto)
			So(err, ShouldNotBeNil)
			So(status, ShouldEqual, services.TransferStatusSufficientFunds)

			// 验证资金
			no1 := am.Service.GetAccountByNo(account1.AccountNo)
			So(no1, ShouldNotBeNil)
			So(no1.Balance.String(), ShouldEqual, account1.Balance.String())
			no2 := am.Service.GetAccountByNo(account2.AccountNo)
			So(no2, ShouldNotBeNil)
			So(no2.Balance.String(), ShouldEqual, account2.Balance.String())
		})

		// 储值
		Convey("储值", func() {

			amount := decimal.NewFromFloat(100)

			body := services.TradeParticipator{
				AccountNo: account1.AccountNo,
				UserId:    account1.UserId,
				Username:  account1.Username,
			}

			dto := services.AccountTransferDTO{
				TradeNo:    ksuid.New().Next().String(),
				TradeBody:  body,
				AmountStr:  "100",
				ChangeType: services.ChangeTypeAccountStoreValue,
				ChangeFlag: services.ChangeFlagTransferIn,
				Desc:       "储值",
			}

			status, err := am.Service.StoreValue(dto)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, services.TransferStatusSuccess)

			// 验证资金
			no := am.Service.GetAccountByNo(account1.AccountNo)
			So(no, ShouldNotBeNil)
			So(no.Balance.String(), ShouldEqual, account1.Balance.Add(amount).String())
		})
	})
}
