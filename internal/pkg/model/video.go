package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Video struct {
	ID          int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name        string    `gorm:"column:name;type:varchar(255);not null;default:'';comment:'视频名称'"`
	Url         string    `gorm:"column:url;type:varchar(255);not null;default:'';comment:'视频地址'"`
	Rank        int32     `gorm:"column:rank;type:int(11);not null;default:0;comment:'视频排名'"`
	CreatedTime time.Time `gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

func (obj *Video) TableName() string {
	return "t_video"
}

func (obj *Video) Create(video *Video) (*Video, error) {
	err := db.Get().Create(video).Error
	return video, err
}

func (obj *Video) Get(video *Video) (*Video, error) {
	result := new(Video)
	err := db.Get().Table(obj.TableName()).Where(video).First(result).Error
	return result, err
}

func (obj *Video) Gets(video *Video) ([]Video, error) {
	results := make([]Video, 0)
	err := db.Get().Table(obj.TableName()).Where(video).Find(&results).Error
	return results, err
}

func (obj *Video) Update(video *Video) (*Video, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", video.ID).Updates(video).Error
	return video, err
}

func (obj *Video) Delete(video *Video) error {
	return db.Get().Delete(video, "id = ?", video.ID).Error
}

func (obj *Video) GetByDescRank(limit int32) ([]Video, error) {
	results := make([]Video, 0)
	err := db.Get().Table(obj.TableName()).Order("rank desc").Limit(int(limit)).Find(&results).Error
	return results, err
}
