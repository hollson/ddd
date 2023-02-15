//Package po generated by 'freedom new-po'
package po

import (
	"gorm.io/gorm"
	"time"
)

// User .
type User struct {
	changes  map[string]interface{}
	ID       int       `gorm:"primaryKey;column:id"` // 用户id
	Name     string    `gorm:"column:name"`          // 用户名称
	Money    int       `gorm:"column:money"`         // 金钱
	Password string    `gorm:"column:password"`      // 密码
	Created  time.Time `gorm:"column:created"`
	Updated  time.Time `gorm:"column:updated"`
}

// TableName .
func (obj *User) TableName() string {
	return "user"
}

// Location .
func (obj *User) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.ID}
}

// GetChanges .
func (obj *User) GetChanges() map[string]interface{} {
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
func (obj *User) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}

// SetName .
func (obj *User) SetName(name string) {
	obj.Name = name
	obj.Update("name", name)
}

// SetMoney .
func (obj *User) SetMoney(money int) {
	obj.Money = money
	obj.Update("money", money)
}

// SetPassword .
func (obj *User) SetPassword(password string) {
	obj.Password = password
	obj.Update("password", password)
}

// SetCreated .
func (obj *User) SetCreated(created time.Time) {
	obj.Created = created
	obj.Update("created", created)
}

// SetUpdated .
func (obj *User) SetUpdated(updated time.Time) {
	obj.Updated = updated
	obj.Update("updated", updated)
}

// AddMoney .
func (obj *User) AddMoney(money int) {
	obj.Money += money
	obj.Update("money", gorm.Expr("money + ?", money))
}