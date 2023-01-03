package recommend

import "wxcloudrun-golang/internal/pkg/model"

type Service struct {
	RecommendDao *model.Recommend
}

func NewService() *Service {
	return &Service{
		RecommendDao: &model.Recommend{},
	}
}

func (s *Service) GetRecommend() ([]model.Recommend, error) {
	recommends, err := s.RecommendDao.Gets(&model.Recommend{})
	if err != nil {
		return nil, err
	}
	return recommends, nil
}
