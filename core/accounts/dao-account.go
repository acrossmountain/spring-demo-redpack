package accounts

import (
	"github.com/go-spring/spring-logger"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type AccountDao struct {
	runner *gorm.DB
}

// 查询一行数据
func (dao *AccountDao) GetOne(accountNo string) *Account {
	a := &Account{}
	if err := dao.runner.Where("account_no = ?", accountNo).First(a).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			SpringLogger.Info("record not found")
			return nil
		}
		SpringLogger.Error(err)
		return nil
	}
	return a
}

// 通过用户ID和账户类型查询账户信息
func (dao *AccountDao) GetUserById(userId string, accountType int) *Account {
	a := &Account{}

	if err := dao.runner.Where("user_id = ? and account_type = ? ",
		userId, accountType).First(a).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			SpringLogger.Info("record not found")
			return nil
		}
		SpringLogger.Error(err)
		return nil
	}

	return a
}

// 账户数据的插入
func (dao *AccountDao) Insert(a *Account) (id int64, err error) {
	result := dao.runner.Create(a)

	if result.Error != nil {
		SpringLogger.Error(err)
		return 0, result.Error
	}

	return a.Id, nil
}

// 账户余额的更新
func (dao *AccountDao) UpdateBalance(accountNo string, amount decimal.Decimal) (
	rows int64, err error) {

	result := dao.runner.Model(&Account{}).
		Where("account_no = ? and balance >= -1 * CAST(? AS DECIMAL(30,6)) ", accountNo, amount.String()).
		Update("balance", gorm.Expr("balance + CAST(? AS DECIMAL(30,6))", amount.String()))

	if result.Error != nil {
		SpringLogger.Error(result.Error)
		return 0, nil
	}

	return result.RowsAffected, nil
}

// 账户状态更新
func (dao *AccountDao) UpdateStatus(accountNo string, status int) (rows int64, err error) {
	result := dao.runner.Model(&Account{}).
		Where("account_no = ?", accountNo).
		Update("status = ?", status)
	if result.Error != nil {
		SpringLogger.Error(result.Error)
		return 0, err
	}
	return result.RowsAffected, nil
}
