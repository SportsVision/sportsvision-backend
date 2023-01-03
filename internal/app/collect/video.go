package collect

import (
	"fmt"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	CollectDao *model.Collect
}

func NewService() *Service {
	return &Service{
		CollectDao: &model.Collect{},
	}
}

func (s *Service) ToggleCollectVideo(openID string, fileID string) (*model.Collect, error) {
	// 查询是否已经收藏过
	collects, err := s.CollectDao.Gets(&model.Collect{OpenID: openID, FileID: fileID})
	fmt.Println(collects)
	if err != nil {
		return nil, err
	}
	if len(collects) > 0 {
		collect, err := s.CollectDao.Update(&model.Collect{ID: collects[0].ID, Status: collects[0].Status * (-1)})
		if err != nil {
			return nil, err
		}
		return collect, nil
	}
	// 创建收藏
	collect, err := s.CollectDao.Create(&model.Collect{
		OpenID:      openID,
		FileID:      fileID,
		Status:      1,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return collect, nil
}

func (s *Service) GetCollectByUser(userOpenID string) ([]model.Collect, error) {
	collects, err := s.CollectDao.Gets(&model.Collect{OpenID: userOpenID, Status: 1})
	if err != nil {
		return nil, err
	}
	return collects, nil
}
