package envelopes

import (
	"time"

	"github.com/acrossmounation/redpack/services"

	"github.com/go-spring/spring-logger"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type RedEnvelopeGoodsDao struct {
	runner *gorm.DB
}

// 插入
func (dao *RedEnvelopeGoodsDao) Insert(po *RedEnvelopeGoods) (int64, error) {
	err := dao.runner.Create(po).Error
	if err != nil {
		SpringLogger.Error(err)
		return 0, err
	}
	return po.Id, nil
}

// 查询，根据红包编号来
func (dao *RedEnvelopeGoodsDao) GetOne(envelopeNo string) *RedEnvelopeGoods {
	po := &RedEnvelopeGoods{}
	err := dao.runner.Where("envelope_no = ?", envelopeNo).First(po).Error
	if err != nil {
		SpringLogger.Error(err)
		return nil
	}
	return po
}

// 更新红包余额和数量
func (dao *RedEnvelopeGoodsDao) UpdateBalance(
	envelopeNo string, amount decimal.Decimal) (int64, error) {

	result := dao.runner.
		Model(&RedEnvelopeGoods{}).
		Where("envelope_no = ? and remain_quantity > 0 and remain_amount >=  CAST(? AS DECIMAL(30,6))", envelopeNo, amount.String()).
		Update(map[string]interface{}{
			"remain_amount":   gorm.Expr("remain_amount - CAST(? AS DECIMAL(30,6))", amount.String()),
			"remain_quantity": gorm.Expr("remain_quantity - ?", 1),
		})

	if result.Error != nil {
		SpringLogger.Error(result.Error)
		return 0, nil
	}
	return result.RowsAffected, nil
}

// 更新订单状态
func (dao *RedEnvelopeGoodsDao) UpdateOrderStatus(
	envelopeNo string, status services.OrderStatus) (int64, error) {

	result := dao.runner.
		Model(&RedEnvelopeGoods{}).
		Where("envelope_no = ?", envelopeNo).
		Update("status", status)

	if result.Error != nil {
		SpringLogger.Error(result.Error)
		return 0, nil
	}
	return result.RowsAffected, nil
}

// 过期，把过期的红包查询出来（分页）
func (dao *RedEnvelopeGoodsDao) FindExpired(offset, size int) []RedEnvelopeGoods {
	var goods []RedEnvelopeGoods
	now := time.Now()

	err := dao.runner.Model(&RedEnvelopeGoods{}).
		Where("expired_at < ?", now).
		Offset(offset).
		Limit(size).
		Find(&goods).Error

	if err != nil {
		SpringLogger.Error(err)
	}
	return goods
}
