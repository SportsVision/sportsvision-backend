package model

import "wxcloudrun-golang/internal/pkg/db"

type Court struct {
	ID        int32  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `json:"name" gorm:"column:name;type:varchar(255);not null;default:'';comment:'场馆名称'"`
	Latitude  int32  `json:"latitude" gorm:"column:latitude;type:int(11);not null;default:0;comment:'纬度'"`
	Longitude int32  `json:"longitude" gorm:"column:longitude;type:int(11);not null;default:0;comment:'经度'"`
	PicURL    string `json:"pic_url" gorm:"column:pic_url;type:varchar(255);not null;default:'';comment:'场馆图片'"`
	Info      string `json:"info" gorm:"column:info;type:varchar(255);not null;default:'';comment:'场馆简介'"`
}

// TableName get sql table name.获取数据库名字
func (obj *Court) TableName() string {
	return "t_court"
}

// Create 创建记录
func (obj *Court) Create(count *Court) (*Court, error) {
	err := db.Get().Create(count).Error
	return count, err
}

// Get 获取
func (obj *Court) Get(count *Court) (*Court, error) {
	result := new(Court)
	err := db.Get().Table(obj.TableName()).Where(count).First(result).Error
	return result, err
}

// Gets 获取批量结果
func (obj *Court) Gets(count *Court) ([]Court, error) {
	results := make([]Court, 0)
	err := db.Get().Table(obj.TableName()).Where(count).Find(&results).Error
	return results, err
}

// Update 更新
func (obj *Court) Update(count *Court) (*Court, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", count.ID).Updates(count).Error
	return count, err
}

// Delete 删除
func (obj *Court) Delete(count *Court) error {
	return db.Get().Delete(count, "id = ?", count.ID).Error
}

// GetsGetsWithLimit 获取批量结果
func (obj *Court) GetsWithLimit(count *Court, limit int32) ([]Court, error) {
	results := make([]Court, 0)
	err := db.Get().Table(obj.TableName()).Where(count).Limit(int(limit)).Find(&results).Error
	return results, err
}
