package accounts

import (
	"database/sql"
	"time"

	"github.com/acrossmounation/redpack/services"

	"github.com/shopspring/decimal"
)

type Account struct {
	Id           int64           `gorm:"column:id;type:bigint(20) AUTO_INCREMENT;not null;comment:'账户ID'"`
	AccountNo    string          `gorm:"column:account_no;type:varchar(32);not null;unique;comment:'账户编号'"`
	AccountName  string          `gorm:"column:account_name;type:varchar(64);not null;comment:'账户名称'"`
	AccountType  int             `gorm:"column:account_type;type:tinyint(2);not null;comment:'账户类型'"`
	CurrencyCode string          `gorm:"column:currency_code;type:char(10);not null;default:'CNY';comment:'货币类型'"`
	UserId       string          `gorm:"column:user_id;type:varchar(40);not null;comment:'用户编号'"`
	Username     sql.NullString  `gorm:"column:username;type:varchar(64);default null;comment:'用户名称'"`
	Balance      decimal.Decimal `gorm:"column:balance;type:decimal(30, 6) unsigned;DEFAULT:'0.000000';comment:'账户可用余额'"`
	Status       int             `gorm:"column:status;type:tinyint(2);not null;comment:'账户状态: 0账户初始化，1启用，2停用'"`
	CreatedAt    time.Time       `gorm:"column:created_at;comment:'创建时间'"`
	UpdatedAt    time.Time       `gorm:"column:updated_at;comment:'更新时间'"`
}

func (po *Account) ToDTO() *services.AccountDTO {
	dto := &services.AccountDTO{}
	dto.AccountNo = po.AccountNo
	dto.AccountName = po.AccountName
	dto.AccountType = po.AccountType
	dto.CurrencyCode = po.CurrencyCode
	dto.UserId = po.UserId
	dto.Username = po.Username.String
	//dto.Balance = po.Balance
	//dto.Status = po.Status
	dto.CreatedAt = po.CreatedAt
	//dto.UpdatedAt = po.UpdatedAt
	return dto
}

func (po *Account) FromDTO(dto *services.AccountDTO) {
	po.AccountNo = dto.AccountNo
	po.AccountName = dto.AccountName
	po.AccountType = dto.AccountType
	po.CurrencyCode = dto.CurrencyCode
	po.UserId = dto.UserId
	po.Username = sql.NullString{Valid: true, String: dto.Username}
	//po.Balance = dto.Balance
	//po.Status = dto.Status
	po.CreatedAt = dto.CreatedAt
	//po.UpdatedAt = dto.UpdatedAt
}
