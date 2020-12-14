package accounts

import (
	"time"

	"github.com/acrossmounation/redpack/services"

	"github.com/shopspring/decimal"
)

type AccountLog struct {
	Id              int64               `gorm:"column:id;type:bigint(20) auto_increment;not null"`
	LogNo           string              `gorm:"column:log_no;type:varchar(32);not null;unique;comment:'流水编号'"`
	TradeNo         string              `gorm:"column:trade_no;type:varchar(32);not null;comment:'交易单号'"`
	AccountNo       string              `gorm:"column:account_no;type:varchar(32);not null;comment:'账户编号'"`
	UserId          string              `gorm:"column:user_id;type:varchar(40);not null;comment:'用户编号'"`
	Username        string              `gorm:"column:username;type:varchar(64);not null;comment:'用户名称'"`
	TargetAccountNo string              `gorm:"column:target_account_no;type:varchar(32);comment:'账户编号'"`
	TargetUserId    string              `gorm:"column:target_user_id;type:varchar(40);not null;comment:'目标用户编号'"`
	TargetUsername  string              `gorm:"column:target_username;type:varchar(64);not null;comment:'目标用户名称'"`
	Amount          decimal.Decimal     `gorm:"column:amount;type:decimal(30,6);not null;default:'0.000000';comment:'交易金额'"`
	Balance         decimal.Decimal     `gorm:"column:balance;type:decimal(30,6);not null;default:'0.000000';comment:'交易后余额'"`
	ChangeType      services.ChangeType `gorm:"column:change_type;type:tinyint(2);not null;default:'0';comment:'流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义'"`
	ChangeFlag      services.ChangeFlag `gorm:"column:change_flag;type:tinyint(2);not null;default:'0';comment:'交易变化标识：-1 出账 1为进账，枚举'"`
	Status          int                 `gorm:"column:status;type:tinyint(2);not null;default:'0';comment:'交易状态'"`
	Decs            string              `gorm:"column:decs;type:varchar(128);not null;comment:'交易描述'"`
	CreatedAt       time.Time           `gorm:"column:created_at;comment:'创建时间'"`
}

func (AccountLog) TableName() string {
	return "account_log"
}

func (po *AccountLog) FromTransferDTO(dto *services.AccountTransferDTO) {
	po.TradeNo = dto.TradeNo
	po.AccountNo = dto.TradeBody.AccountNo
	po.TargetAccountNo = dto.TradeTarget.AccountNo
	po.UserId = dto.TradeBody.UserId
	po.Username = dto.TradeBody.Username
	po.TargetUserId = dto.TradeTarget.UserId
	po.TargetUsername = dto.TradeTarget.Username
	po.Amount = dto.Amount
	po.ChangeType = dto.ChangeType
	po.ChangeFlag = dto.ChangeFlag
	po.Decs = dto.Desc
}

func (po *AccountLog) ToDTO() *services.AccountLogDTO {
	dto := &services.AccountLogDTO{
		TradeNo:         po.TradeNo,
		LogNo:           po.LogNo,
		AccountNo:       po.AccountNo,
		TargetAccountNo: po.TargetAccountNo,
		UserId:          po.UserId,
		Username:        po.Username,
		TargetUserId:    po.TargetUserId,
		TargetUsername:  po.TargetUsername,
		Amount:          po.Amount,
		Balance:         po.Balance,
		ChangeType:      po.ChangeType,
		ChangeFlag:      po.ChangeFlag,
		Status:          po.Status,
		Decs:            po.Decs,
		CreatedAt:       po.CreatedAt,
	}
	return dto
}

func (po *AccountLog) FromDTO(dto *services.AccountLogDTO) {

	po.TradeNo = dto.TradeNo
	po.LogNo = dto.LogNo
	po.AccountNo = dto.AccountNo
	po.TargetAccountNo = dto.TargetAccountNo
	po.UserId = dto.UserId
	po.Username = dto.Username
	po.TargetUserId = dto.TargetUserId
	po.TargetUsername = dto.TargetUsername
	po.Amount = dto.Amount
	po.Balance = dto.Balance
	po.ChangeType = dto.ChangeType
	po.ChangeFlag = dto.ChangeFlag
	po.Status = dto.Status
	po.Decs = dto.Decs
	po.CreatedAt = dto.CreatedAt
}
