package envelopes

import (
	"context"
	"time"

	"github.com/acrossmounation/redpack/services"
	"github.com/acrossmounation/redpack/utils"

	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
)

type RedEnvelopeGoodsDomain struct {
	RedEnvelopeGoods
	Db *gorm.DB `autowire:""`
}

// 生成红包编号
func (domain *RedEnvelopeGoodsDomain) createEnvelopeNo() {
	domain.EnvelopeNo = ksuid.New().Next().String()
}

// 创建红包对象
func (domain *RedEnvelopeGoodsDomain) generate(dto services.RedEnvelopeGoodsDTO) {
	domain.RedEnvelopeGoods.FromDTO(&dto)
	domain.RemainQuantity = dto.Quantity
	domain.Username.Valid = true
	domain.Blessing.Valid = true

	// 普通红包
	if domain.EnvelopeType == services.GeneralEnvelopeType {
		domain.Amount = dto.Amount.Mul(decimal.NewFromFloat(float64(domain.Quantity)))
	}

	// 碰运气红包
	if domain.EnvelopeType == services.LuckyEnvelopeType {
		domain.AmountOne = decimal.NewFromFloat(0)
	}
	domain.RemainAmount = dto.Amount
	// 过期时间
	domain.ExpiredAt = time.Now().Add(24 * time.Hour)
	domain.Status = services.OrderCreate
	domain.createEnvelopeNo()
}

// 保存到红包商品表
func (domain *RedEnvelopeGoodsDomain) save(ctx context.Context) (id int64, err error) {
	err = utils.ExecuteContext(ctx, func(tx *gorm.DB) error {
		dao := RedEnvelopeGoodsDao{runner: tx}
		id, err = dao.Insert(&domain.RedEnvelopeGoods)
		return err
	})
	return id, err
}

// 创建并保存
func (domain *RedEnvelopeGoodsDomain) Create(ctx context.Context, dto services.RedEnvelopeGoodsDTO) (id int64, err error) {
	// 创建红包商品
	domain.generate(dto)
	return domain.save(ctx)
}
