package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Event struct {
	ID          int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	OpenID      string    `json:"open_id" gorm:"column:open_id;type:int(11);not null;default:0;comment:'用户id'"`
	CourtID     int32     `json:"court_id" gorm:"column:court_id;type:int(11);not null;default:0;comment:'场馆id'"`
	Date        int32     `json:"date" gorm:"column:date;type:int(11);not null;default:0;comment:'日期'"`
	StartTime   int32     `json:"start_time" gorm:"column:start_time;type:int(11);not null;default:0;comment:'开始时间'"`
	EndTime     int32     `json:"end_time" gorm:"column:end_time;type:int(11);not null;default:0;comment:'结束时间'"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

// TableName get sql table name.获取数据库名字
func (obj *Event) TableName() string {
	return "t_event"
}

// Create 创建记录
func (obj *Event) Create(event *Event) (*Event, error) {
	err := db.Get().Create(event).Error
	return event, err
}

// Get 获取
func (obj *Event) Get(event *Event) (*Event, error) {
	result := new(Event)
	err := db.Get().Table(obj.TableName()).Where(event).First(result).Error
	return result, err
}

// GetsByDesc 获取批量结果
func (obj *Event) GetsByDesc(event *Event) ([]Event, error) {
	results := make([]Event, 0)
	err := db.Get().Table(obj.TableName()).Where(event).Find(&results).Order("start_time desc").Error
	return results, err
}

// Update 更新
func (obj *Event) Update(event *Event) (*Event, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", event.ID).Updates(event).Error
	return event, err
}

// Delete 删除
func (obj *Event) Delete(event *Event) error {
	return db.Get().Delete(event, "id = ?", event.ID).Error
}
