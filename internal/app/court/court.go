package court

import (
	"log"
	"strconv"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/pkg/location"
)

type Service struct {
	courtDao *model.Court
}

func NewService() *Service {
	return &Service{
		courtDao: &model.Court{},
	}
}

type CourtWithDistance struct {
	model.Court
	Distance float64 `json:"distance"`
}

// GetCourts 获取所有场地，按距离倒序排列
func (s *Service) GetCourts(latitude, longitude string) ([]model.Court, error) {
	results, err := s.courtDao.Gets(&model.Court{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// convert to float
	lat, _ := strconv.ParseFloat(latitude, 64)
	lng, _ := strconv.ParseFloat(longitude, 64)

	var courtsWithDistance []CourtWithDistance
	// calculate distance
	for i := 0; i < len(results); i++ {
		distance := location.GetDistance(lat, lng, results[i].Latitude, results[i].Longitude)
		courtsWithDistance = append(courtsWithDistance, CourtWithDistance{
			Court:    results[i],
			Distance: distance,
		})
	}
	// sort by distance
	sortByDistance(courtsWithDistance)
	return results, nil
}

func sortByDistance(courts []CourtWithDistance) {
	for i := 0; i < len(courts); i++ {
		for j := i + 1; j < len(courts); j++ {
			if courts[i].Distance > courts[j].Distance {
				courts[i], courts[j] = courts[j], courts[i]
			}
		}
	}
}

func (s *Service) GetAllCourts() ([]model.Court, error) {
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
