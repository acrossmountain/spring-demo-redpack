package accounts

import (
	"github.com/go-spring/spring-logger"
	"github.com/jinzhu/gorm"
)

type AccountLogDao struct {
	runner *gorm.DB
}

// 通过流水编号
func (dao *AccountLogDao) GetOne(logNo string) *AccountLog {
	al := &AccountLog{}
	if err := dao.runner.Where("log_no = ?", logNo).First(al).Error; err != nil {
		SpringLogger.Error(err)
		return nil
	}
	return al
}

// 通过交易编号
func (dao *AccountLogDao) GetByTradeNo(tradeNo string) *AccountLog {
	al := &AccountLog{}
	if err := dao.runner.Where("trade_no = ?", tradeNo).First(al).Error; err != nil {
		SpringLogger.Error(err)
		return nil
	}
	return al
}

// 流水记录写入
func (dao *AccountLogDao) Insert(al *AccountLog) (id int64, err error) {
	result := dao.runner.Create(al)
	if result.Error != nil {
		SpringLogger.Error(err)
		return 0, nil
	}
	return al.Id, nil
}
