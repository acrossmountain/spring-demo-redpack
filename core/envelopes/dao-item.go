package envelopes

import (
	"github.com/go-spring/spring-logger"
	"github.com/jinzhu/gorm"
)

type RedEnvelopeGoodsItemDao struct {
	runner *gorm.DB
}

// 查询
func (dao *RedEnvelopeGoodsItemDao) GetOne(itemNo string) *RedEnvelopeItem {
	redEnvelopeItem := &RedEnvelopeItem{}
	err := dao.runner.Where("item_no = ?", itemNo).First(redEnvelopeItem).Error
	if err != nil {
		SpringLogger.Error(err)
		return nil
	}
	return redEnvelopeItem
}

// 插入
func (dao *RedEnvelopeGoodsItemDao) Insert(po *RedEnvelopeItem) (int64, error) {
	result := dao.runner.Model(&RedEnvelopeItem{}).Create(po)
	if result.Error != nil {
		SpringLogger.Error(result.Error)
		return 0, nil
	}
	return result.RowsAffected, nil
}
