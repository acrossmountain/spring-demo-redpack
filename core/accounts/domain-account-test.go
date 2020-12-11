package accounts

import (
	"testing"

	"github.com/acrossmounation/redpack/services"

	"github.com/go-spring/spring-boot"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

type TestAccountDomain struct {
	_      SpringBoot.JUnitSuite `export:""`
	Domain *AccountDomain        `autowire:""`
}

func (am *TestAccountDomain) Test(t *testing.T) {
	dto := services.AccountDTO{
		UserId:   ksuid.New().Next().String(),
		Username: "测试用户",
		Balance:  decimal.NewFromFloat(0),
		Status:   1,
	}

	Convey("创建账户", t, func() {
		ndto, err := am.Domain.Create(dto)
		So(err, ShouldBeNil)
		So(ndto, ShouldNotBeNil)
		So(ndto.Balance.String(), ShouldEqual, dto.Balance.String())
		So(ndto.Username, ShouldEqual, dto.Username)
		So(ndto.UserId, ShouldEqual, dto.UserId)
		So(ndto.Status, ShouldEqual, dto.Status)
	})
}

type TestAccountDomainTransfer struct {
	_      SpringBoot.JUnitSuite `export:""`
	Domain *AccountDomain        `autowire:""`
}

func (am *TestAccountDomainTransfer) Test(t *testing.T) {
	dto1 := services.AccountDTO{
		UserId:   ksuid.New().Next().String(),
		Username: "测试用户1",
		Balance:  decimal.NewFromFloat(100),
		Status:   1,
	}

	dto2 := services.AccountDTO{
		UserId:   ksuid.New().Next().String(),
		Username: "测试用户2",
		Balance:  decimal.NewFromFloat(100),
		Status:   1,
	}

	Convey("转账测试", t, func() {

		// 账户1创建
		ndto1, err := am.Domain.Create(dto1)
		So(err, ShouldBeNil)
		So(ndto1, ShouldNotBeNil)
		So(ndto1.Balance.String(), ShouldEqual, dto1.Balance.String())
		So(ndto1.Username, ShouldEqual, dto1.Username)
		So(ndto1.UserId, ShouldEqual, dto1.UserId)
		So(ndto1.Status, ShouldEqual, dto1.Status)

		// 账户2创建
		ndto2, err := am.Domain.Create(dto2)
		So(err, ShouldBeNil)
		So(ndto2, ShouldNotBeNil)
		So(ndto2.Balance.String(), ShouldEqual, dto2.Balance.String())
		So(ndto2.Username, ShouldEqual, dto2.Username)
		So(ndto2.UserId, ShouldEqual, dto2.UserId)
		So(ndto2.Status, ShouldEqual, dto2.Status)

		// 转账（余额充足）
		Convey("转账（余额充足）", func() {

			amount := decimal.NewFromFloat(1)

			body := services.TradeParticipator{
				AccountNo: ndto1.AccountNo,
				UserId:    ndto1.UserId,
				Username:  ndto1.Username,
			}
			target := services.TradeParticipator{
				AccountNo: ndto2.AccountNo,
				UserId:    ndto2.UserId,
				Username:  ndto2.Username,
			}

			dto := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   body,
				TradeTarget: target,
				AmountStr:   amount.String(),
				Amount:      amount,
				ChangeType:  services.ChangeType(-1),
				ChangeFlag:  services.ChangeFlagTransferOut,
				Desc:        "转账",
			}

			status, err := am.Domain.Transfer(dto)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, services.TransferStatusSuccess)

			// 验证资金
			no := am.Domain.GetAccountByNo(ndto1.AccountNo)
			So(no, ShouldNotBeNil)
			So(no.Balance.String(), ShouldEqual, ndto1.Balance.Sub(amount).String())
		})

		// 转账（余额充足）
		Convey("转账（余额不足）", func() {

			amount := decimal.NewFromFloat(110)

			body := services.TradeParticipator{
				AccountNo: ndto1.AccountNo,
				UserId:    ndto1.UserId,
				Username:  ndto1.Username,
			}
			target := services.TradeParticipator{
				AccountNo: ndto2.AccountNo,
				UserId:    ndto2.UserId,
				Username:  ndto2.Username,
			}

			dto := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   body,
				TradeTarget: target,
				AmountStr:   amount.String(),
				Amount:      amount,
				ChangeType:  services.ChangeType(-1),
				ChangeFlag:  services.ChangeFlagTransferOut,
				Desc:        "转账",
			}

			status, err := am.Domain.Transfer(dto)
			So(err, ShouldNotBeNil)
			So(status, ShouldEqual, services.TransferStatusSufficientFunds)

			// 验证资金
			no := am.Domain.GetAccountByNo(ndto1.AccountNo)
			So(no, ShouldNotBeNil)
			So(no.Balance.String(), ShouldEqual, ndto1.Balance.String())
		})

		// 储值
		Convey("储值", func() {

			amount := decimal.NewFromFloat(100)

			body := services.TradeParticipator{
				AccountNo: ndto1.AccountNo,
				UserId:    ndto1.UserId,
				Username:  ndto1.Username,
			}
			target := services.TradeParticipator{
				AccountNo: ndto2.AccountNo,
				UserId:    ndto2.UserId,
				Username:  ndto2.Username,
			}

			dto := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   body,
				TradeTarget: target,
				Amount:      amount,
				ChangeType:  services.ChangeTypeAccountStoreValue,
				ChangeFlag:  services.ChangeFlagTransferIn,
				Desc:        "储值",
			}

			status, err := am.Domain.Transfer(dto)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, services.TransferStatusSuccess)

			// 验证资金
			no := am.Domain.GetAccountByNo(ndto1.AccountNo)
			So(no, ShouldNotBeNil)
			So(no.Balance.String(), ShouldEqual, ndto1.Balance.Add(amount).String())
		})
	})
}
