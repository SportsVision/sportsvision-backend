package event

import (
	"time"
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	EventDao *model.Event
	VideoDao *model.Video
}

func NewService() *Service {
	return &Service{
		EventDao: &model.Event{},
	}
}

type EventRepos struct {
	model.Event
	Repos []string
}

func (s *Service) CreateEvent(userOpenID, courtID int32, date, startTime, endTime int32) (*model.Event, error) {
	// create event
	event, err := s.EventDao.Create(&model.Event{
		OpenID:      userOpenID,
		CourtID:     courtID,
		Date:        date,
		StartTime:   startTime,
		EndTime:     endTime,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return event, err
}

func (s *Service) GetEventsByUser(userOpenID int32) ([]model.Event, error) {
	events, err := s.EventDao.GetsByDesc(&model.Event{OpenID: userOpenID})
	if err != nil {
		return nil, err
	}
	return events, nil
}

//func (s *Service) GetEventRepos(openID int32)(&EventRepo,error)
