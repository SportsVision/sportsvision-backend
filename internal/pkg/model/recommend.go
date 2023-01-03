package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

// Recommend recommend model
type Recommend struct {
	ID          int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key;auto_increment"`
	VideoURL    string    `json:"video_url" gorm:"column:video_url;type:varchar(255);not null"`
	Desc        string    `json:"desc" gorm:"column:desc;type:varchar(255);not null"`
	LikedCount  int32     `json:"liked_count" gorm:"liked_count"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

// TableName get sql table name.获取数据库名字
func (obj *Recommend) TableName() string {
	return "t_recommend"
}

// Create 创建记录
func (obj *Recommend) Create(r *Recommend) (*Recommend, error) {
	err := db.Get().Create(r).Error
	return r, err
}

// Get 获取
func (obj *Recommend) Get(r *Recommend) (*Recommend, error) {
	result := new(Recommend)
	err := db.Get().Table(obj.TableName()).Where(r).First(result).Error
	return result, err
}

// Gets 获取批量结果
func (obj *Recommend) Gets(r *Recommend) ([]Recommend, error) {
	results := make([]Recommend, 0)
	err := db.Get().Table(obj.TableName()).Where(r).Find(&results).Error
	return results, err
}

// Update 更新
func (obj *Recommend) Update(r *Recommend) (*Recommend, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", r.ID).Updates(r).Error
	return r, err
}

// Delete 删除
func (obj *Recommend) Delete(r *Recommend) error {
	return db.Get().Delete(r, "id = ?", r.ID).Error
}
