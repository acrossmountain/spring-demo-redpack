package accounts

import (
	"errors"
	"github.com/shopspring/decimal"

	"github.com/acrossmounation/redpack/services"

	"github.com/go-spring/spring-logger"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
)

type AccountDomain struct {
	account    Account
	accountLog AccountLog
	Db         *gorm.DB
}

// logNo
func (domain *AccountDomain) createAccountLogNo() {
	domain.accountLog.LogNo = ksuid.New().Next().String()
}

// accountNo
func (domain *AccountDomain) createAccountNo() {
	domain.account.AccountNo = ksuid.New().Next().String()
}

// 创建账户日志
func (domain *AccountDomain) createAccountLog() {
	domain.accountLog = AccountLog{}
	domain.createAccountLogNo()
	domain.accountLog.TradeNo = domain.accountLog.LogNo

	// 主体信息
	domain.accountLog.AccountNo = domain.account.AccountNo
	domain.accountLog.UserId = domain.account.UserId
	domain.accountLog.Username = domain.account.Username.String

	// 交易对象
	domain.accountLog.TargetAccountNo = domain.account.AccountNo
	domain.accountLog.TargetUserId = domain.account.UserId
	domain.accountLog.TargetUsername = domain.account.Username.String

	// 交易金额
	domain.accountLog.Amount = domain.account.Balance
	domain.accountLog.Balance = domain.account.Balance

	// 交易信息
	domain.accountLog.Decs = "账户创建"
	domain.accountLog.ChangeType = services.ChangeTypeAccountCreated
	domain.accountLog.ChangeFlag = services.ChangeFlagAccountCreated
}

// 账户创建
func (domain *AccountDomain) Create(dto services.AccountDTO) (
	*services.AccountDTO, error) {
	// 账户对象
	domain.account = Account{}
	domain.account.FromDTO(&dto)
	domain.createAccountNo()
	domain.account.Username.Valid = true
	// 账户流水
	domain.createAccountLog()

	var ado *services.AccountDTO

	err := domain.Db.Transaction(func(tx *gorm.DB) error {
		accountDao := &AccountDao{runner: tx}
		accountLogDao := &AccountLogDao{runner: tx}

		// 插入账户数据
		id, err := accountDao.Insert(&domain.account)
		if err != nil {
			SpringLogger.Error(err)
			return err
		}
		if id <= 0 {
			SpringLogger.Error("create account fail")
			return errors.New("create account fail")
		}

		// 账户流水日志
		id, err = accountLogDao.Insert(&domain.accountLog)
		if err != nil {
			SpringLogger.Error(err)
			return err
		}
		if id <= 0 {
			SpringLogger.Error("create account log fail")
			return errors.New("create account log fail")
		}

		// 查询账户信息
		domain.account = *accountDao.GetOne(domain.account.AccountNo)
		return nil
	})
	// 转 dto 对象
	ado = domain.account.ToDTO()
	return ado, err
}

// 账户转账（单方面转入或转出）
func (domain *AccountDomain) Transfer(dto services.AccountTransferDTO) (
	status services.TransferStatus, err error) {

	// 如果交易变化为支出，修正amount
	amount := dto.Amount
	if dto.ChangeFlag == services.ChangeFlagTransferOut {
		amount = dto.Amount.Mul(decimal.NewFromFloat(-1))
	}

	// 创建流水记录
	domain.accountLog = AccountLog{}
	domain.accountLog.FromTransferDTO(&dto)
	domain.createAccountLogNo()

	err = domain.Db.Transaction(func(tx *gorm.DB) error {
		accountDao := &AccountDao{runner: tx}
		accountLogDao := &AccountLogDao{runner: tx}

		rows, err := accountDao.UpdateBalance(dto.TradeBody.AccountNo, amount)

		if err != nil {
			status = services.TransferStatusFailure
			SpringLogger.Error(err)
			return err
		}

		if rows <= 0 && dto.ChangeFlag == services.ChangeFlagTransferOut {
			status = services.TransferStatusSufficientFunds
			SpringLogger.Error("balance not enough")
			return errors.New("balance not enough")
		}

		account := accountDao.GetOne(dto.TradeBody.AccountNo)
		if account == nil {
			//status = ??
			SpringLogger.Error("account error")
			return errors.New("account error")
		}

		domain.account = *account
		domain.accountLog.Balance = domain.account.Balance
		id, err := accountLogDao.Insert(&domain.accountLog)

		if err != nil || id <= 0 {
			status = services.TransferStatusFailure
			SpringLogger.Error(err)
			return err
		}

		return nil
	})

	if err == nil {
		status = services.TransferStatusSuccess
	}

	return
}
