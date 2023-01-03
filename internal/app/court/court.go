package court

import (
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	courtDao *model.Court
}

func NewService() *Service {
	return &Service{
		courtDao: &model.Court{},
	}
}

func (s *Service) GetCourts() ([]model.Court, error) {
	results, err := s.courtDao.Gets(&model.Court{})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *Service) GetCountInfo(id int32) (*model.Court, error) {
	result, err := s.courtDao.Get(&model.Court{ID: id})
	if err != nil {
		return nil, err
	}
	return result, nil
}
