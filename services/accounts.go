package services

import "time"

type AccountService interface {
	// 创建账户
	CreateAccount(dto AccountCreatedDTO) (*AccountDTO, error)
	// 转账
	Transfer(dto AccountTransferDTO) (TransferStatus, error)
	// 充值
	StoreValue(dto AccountTransferDTO) (TransferStatus, error)
	// 获取账户信息
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
}

// 账户交易参与者
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  string
}

// 账户转移
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	AmountStr   string
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Desc        string
}

// 账户创建
type AccountCreatedDTO struct {
	UserId       string
	Username     string
	AccountName  string
	AccountType  int
	CurrencyCode string
	Amount       string
}

// 账户信息
type AccountDTO struct {
	AccountCreatedDTO
	AccountNo string
	CreatedAt time.Time
}
