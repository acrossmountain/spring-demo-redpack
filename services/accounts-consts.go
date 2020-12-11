package services

// 转账状态
type TransferStatus int8

const (
	// 转账失败
	TransferStatusFailure TransferStatus = -1

	// 余额不足
	TransferStatusSufficientFunds TransferStatus = 0

	// 转账成功
	TransferStatusSuccess TransferStatus = 1
)

// 转账类型 0 = 创建账户 >= 1 进账 <= -1 支出
type ChangeType int8

const (
	// 账户创建
	ChangeTypeAccountCreated ChangeType = 0
	// 储值
	ChangeTypeAccountStoreValue ChangeType = 1
	// 支出（红包）
	ChangeTypeEnvelopeOutGoing ChangeType = -2
	// 收入（红包）
	ChangeTypeEnvelopeIncoming ChangeType = 2
	// 过期（红包）
	ChangeTypeEnvelopeExpiredRefund ChangeType = 3
)

// 转账标识
type ChangeFlag int8

const (
	// 创建账户
	ChangeFlagAccountCreated ChangeFlag = 0
	// 支出
	ChangeFlagTransferOut ChangeFlag = -1
	// 收入
	ChangeFlagTransferIn ChangeFlag = 1
)

// 账户类型
type AccountType int8

const (
	AccountTypeSystemEnvelope AccountType = 1
	AccountTypeEnvelope       AccountType = 2
)
