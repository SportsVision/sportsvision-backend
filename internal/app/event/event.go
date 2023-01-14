package event

import (
	"fmt"
	"log"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/tcos"
)

type Service struct {
	EventDao *model.Event
	VideoDao *model.Video
	CourtDao *model.Court
}

func NewService() *Service {
	return &Service{
		EventDao: &model.Event{},
	}
}

type EventRepos struct {
	model.Event
	CourtName string   `json:"court_name"`
	Videos    []string `json:"videos"`
}

func (s *Service) CreateEvent(userOpenID string, courtID int32, date, startTime, endTime int32) (*model.Event, error) {
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

func (s *Service) GetEventsByUser(userOpenID string) ([]model.Event, error) {
	events, err := s.EventDao.GetsByDesc(&model.Event{OpenID: userOpenID})
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Service) GetEventVideos(openID string) ([]EventRepos, error) {
	events := make([]model.Event, 0)
	var err error
	events, err = s.GetEventsByUser(openID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	results := make([]EventRepos, 0)
	for _, e := range events {
		videos := make([]string, 0)
		startTime := e.StartTime
		for startTime < e.EndTime {
			videoIDs, err := tcos.GetCosFileList(fmt.Sprintf("highlight/court%d/%d/%d/", e.CourtID, e.Date,
				e.StartTime))
			if err != nil {
				log.Println(err)
				return nil, err
			}
			videos = append(videos, videoIDs...)
			if startTime%100 != 0 {
				startTime += 100
				startTime -= 30
			} else {
				startTime += 30
			}
		}
		court, err := s.CourtDao.Get(&model.Court{ID: e.CourtID})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, EventRepos{Event: e, CourtName: court.Name, Videos: videos})
	}
	return results, nil
}
