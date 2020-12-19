package envelopes

import (
	"database/sql"
	"time"

	"github.com/acrossmounation/redpack/services"

	"github.com/shopspring/decimal"
)

type RedEnvelopeItem struct {
	Id           int64           `gorm:"column:id"`             // 自增ID
	ItemNo       string          `gorm:"column:item_no;unique"` // 红包订单详情编号
	EnvelopeNo   string          `gorm:"column:envelope_no"`    // 红包编号,红包唯一标识
	RecvUsername sql.NullString  `gorm:"column:recv_username"`  // 红包接收者用户名称
	RecvUserId   string          `gorm:"column:recv_user_id"`   // 红包接收者用户编号
	Amount       decimal.Decimal `gorm:"column:amount"`         // 收到金额
	Quantity     int             `gorm:"column:quantity"`       // 收到数量：对于收红包来说是1
	RemainAmount decimal.Decimal `gorm:"column:remain_amount"`  // 收到后红包剩余金额
	AccountNo    string          `gorm:"column:account_no"`     // 红包接收者账户ID
	PayStatus    int             `gorm:"column:pay_status"`     // 支付状态：未支付，支付中，已支付，支付失败
	CreatedAt    time.Time       `gorm:"column:created_at"`     // 创建时间
	UpdatedAt    time.Time       `gorm:"column:updated_at"`     // 更新时间
}

func (po *RedEnvelopeItem) ToDTO() *services.RedEnvelopeItemDTO {
	dto := &services.RedEnvelopeItemDTO{
		ItemNo:       po.ItemNo,
		EnvelopeNo:   po.EnvelopeNo,
		RecvUsername: po.RecvUsername.String,
		RecvUserId:   po.RecvUserId,
		Amount:       po.Amount,
		Quantity:     po.Quantity,
		RemainAmount: po.RemainAmount,
		AccountNo:    po.AccountNo,
		PayStatus:    po.PayStatus,
		CreatedAt:    po.CreatedAt,
		UpdatedAt:    po.UpdatedAt,
	}
	return dto
}

func (po *RedEnvelopeItem) FromDTO(dto *services.RedEnvelopeItemDTO) {
	po.ItemNo = dto.ItemNo
	po.EnvelopeNo = dto.EnvelopeNo
	po.RecvUsername = sql.NullString{Valid: true, String: dto.RecvUsername}
	po.RecvUserId = dto.RecvUserId
	po.Amount = dto.Amount
	po.Quantity = dto.Quantity
	po.RemainAmount = dto.RemainAmount
	po.AccountNo = dto.AccountNo
	po.PayStatus = dto.PayStatus
	po.CreatedAt = dto.CreatedAt
	po.UpdatedAt = dto.UpdatedAt
}
