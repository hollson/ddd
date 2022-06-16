//Package po generated by 'freedom new-po'
package po

import (
	"gorm.io/gorm"
	"time"
)

// Order .
type Order struct {
	changes map[string]interface{}
	ID      int       `gorm:"primaryKey;column:id"`
	UserID  int       `gorm:"column:user_id"`  // 用户id
	GoodsID int       `gorm:"column:goods_id"` // 商品id
	Num     int       `gorm:"column:num"`      // 数量
	Created time.Time `gorm:"column:created"`
	Updated time.Time `gorm:"column:updated"`
}

// TableName .
func (obj *Order) TableName() string {
	return "order"
}

// Location .
func (obj *Order) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.ID}
}

// GetChanges .
func (obj *Order) GetChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// Update .
func (obj *Order) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}

// SetUserID .
func (obj *Order) SetUserID(userID int) {
	obj.UserID = userID
	obj.Update("user_id", userID)
}

// SetGoodsID .
func (obj *Order) SetGoodsID(goodsID int) {
	obj.GoodsID = goodsID
	obj.Update("goods_id", goodsID)
}

// SetNum .
func (obj *Order) SetNum(num int) {
	obj.Num = num
	obj.Update("num", num)
}

// SetCreated .
func (obj *Order) SetCreated(created time.Time) {
	obj.Created = created
	obj.Update("created", created)
}

// SetUpdated .
func (obj *Order) SetUpdated(updated time.Time) {
	obj.Updated = updated
	obj.Update("updated", updated)
}

// AddUserID .
func (obj *Order) AddUserID(userID int) {
	obj.UserID += userID
	obj.Update("user_id", gorm.Expr("user_id + ?", userID))
}

// AddGoodsID .
func (obj *Order) AddGoodsID(goodsID int) {
	obj.GoodsID += goodsID
	obj.Update("goods_id", gorm.Expr("goods_id + ?", goodsID))
}

// AddNum .
func (obj *Order) AddNum(num int) {
	obj.Num += num
	obj.Update("num", gorm.Expr("num + ?", num))
}
