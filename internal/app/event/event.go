package event

import (
	"fmt"
	"log"
	"strings"
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

// type EventRepos struct {
// 	model.Event
// 	CourtName     string         `json:"court_name"`
// 	Videos        []string       `json:"videos"`
// 	VideosWithGif []VideoWithGif `json:"videos_with_gif"`
// }

type EventInfos struct {
	model.Event
	CourtName     string          `json:"court_name"`
	Videos        []string        `json:"videos"`
	VideosWithGif []VideoGifBatch `json:"videos_with_gif"`
}

type VideoWithGif struct {
	Gif             string `json:"gif"`
	Video           string `json:"video"`
	LowQualityVideo string `json:"low_quality_video"`
}

type VideoGifBatch struct {
	StartTime string         `json:"start_time"`
	EndTime   string         `json:"end_time"`
	VideoGifs []VideoWithGif `json:"video_gif_batch"`
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

func (s *Service) DeleteEvent(openID string, eventID int32) error {
	return s.EventDao.Delete(&model.Event{OpenID: openID, ID: eventID})
}

func (s *Service) GetEventsByUser(userOpenID string) ([]model.Event, error) {
	events, err := s.EventDao.GetsByDesc(&model.Event{OpenID: userOpenID})
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Service) GetEventInfo(eventID int32) (EventInfos, error) {
	event, err := s.EventDao.Get(&model.Event{ID: eventID})
	if err != nil {
		return EventInfos{}, err
	}
	court, err := s.CourtDao.Get(&model.Court{ID: event.CourtID})
	if err != nil {
		return EventInfos{}, err
	}
	startTime := event.StartTime
	videos := make([]string, 0)
	videoGifBatchs := make([]VideoGifBatch, 0)
	videosWithGif := make([]VideoWithGif, 0)
	for startTime < event.EndTime {
		allLinks, err := tcos.GetCosFileList(fmt.Sprintf("highlight/court%d/%d/%d/", event.CourtID, event.Date,
			startTime))
		if err != nil {
			log.Println(err)
			return EventInfos{}, err
		}
		videoLinks := filterVideos(allLinks)
		videos = append(videos, videoLinks...)
		videosWithGif = append(videosWithGif, getVideosWithGif(videoLinks)...)
		lastStartTime := ``
		lastEndTime := ``
		videoGifs := make([]VideoWithGif, 0)
		for _, video_gif := range videosWithGif {
			urlLastPart := strings.Split(video_gif.Gif, "/n")[1]
			_start_time_str := strings.Split(urlLastPart, "_")[0]
			_end_time_str := strings.Split(urlLastPart, "_")[1]
			if strings.Compare(lastStartTime, ``) == 0 {
				lastStartTime = _start_time_str
				lastEndTime = _end_time_str
			}
			if strings.Compare(_start_time_str, lastStartTime) != 0 {
				videoGifBatchs = append(videoGifBatchs, VideoGifBatch{
					StartTime: lastStartTime,
					EndTime:   lastEndTime,
					VideoGifs: videoGifs,
				})
				lastStartTime = _start_time_str
				lastEndTime = _end_time_str
				videoGifs = make([]VideoWithGif, 0)
			}
			videoGifs = append(videoGifs, video_gif)
		}
		if strings.Compare(lastStartTime, ``) != 0 {
			videoGifBatchs = append(videoGifBatchs, VideoGifBatch{
				StartTime: lastStartTime,
				EndTime:   lastEndTime,
				VideoGifs: videoGifs,
			})
		}

		if startTime%100 != 0 {
			startTime += 100
			startTime -= 30
		} else {
			startTime += 30
		}
	}
	return EventInfos{Event: *event, CourtName: court.Name, Videos: videos, VideosWithGif: videoGifBatchs}, nil
}

func (s *Service) GetEventVideos(openID string) ([]EventInfos, error) {
	events := make([]model.Event, 0)
	var err error
	events, err = s.GetEventsByUser(openID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	results := make([]EventInfos, 0)
	for _, e := range events {
		videos := make([]string, 0)
		videoGifBatchs := make([]VideoGifBatch, 0)
		videosWithGif := make([]VideoWithGif, 0)
		startTime := e.StartTime
		for startTime < e.EndTime {
			allLinks, err := tcos.GetCosFileList(fmt.Sprintf("highlight/court%d/%d/%d/", e.CourtID, e.Date,
				startTime))
			if err != nil {
				log.Println(err)
				return nil, err
			}
			videoLinks := filterVideos(allLinks)
			videos = append(videos, videoLinks...)
			videosWithGif = append(videosWithGif, getVideosWithGif(videoLinks)...)
			lastStartTime := ``
			lastEndTime := ``
			videoGifs := make([]VideoWithGif, 0)
			for _, video_gif := range videosWithGif {
				urlLastPart := strings.Split(video_gif.Gif, "/n")[1]
				_start_time_str := strings.Split(urlLastPart, "_")[0]
				_end_time_str := strings.Split(urlLastPart, "_")[1]
				if strings.Compare(lastStartTime, ``) == 0 {
					lastStartTime = _start_time_str
					lastEndTime = _end_time_str
				}
				if strings.Compare(_start_time_str, lastStartTime) != 0 {
					videoGifBatchs = append(videoGifBatchs, VideoGifBatch{
						StartTime: lastStartTime,
						EndTime:   lastEndTime,
						VideoGifs: videoGifs,
					})
					lastStartTime = _start_time_str
					lastEndTime = _end_time_str
					videoGifs = make([]VideoWithGif, 0)
				}
				videoGifs = append(videoGifs, video_gif)
			}
			if strings.Compare(lastStartTime, ``) != 0 {
				videoGifBatchs = append(videoGifBatchs, VideoGifBatch{
					StartTime: lastStartTime,
					EndTime:   lastEndTime,
					VideoGifs: videoGifs,
				})
			}
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
		results = append(results, EventInfos{Event: e, CourtName: court.Name, Videos: videos,
			VideosWithGif: videoGifBatchs})
	}
	return results, nil
}

func getVideosWithGif(videos []string) []VideoWithGif {
	videosWithGif := make([]VideoWithGif, 0)
	for _, video := range videos {
		// replace mp4 to gif
		links := strings.Split(video, ".")
		links[len(links)-1] = "gif"
		gif := strings.Join(links, ".")
		lowQualityVideos := getLowQualityVideo(video)
		videosWithGif = append(videosWithGif, VideoWithGif{Gif: gif, Video: video, LowQualityVideo: lowQualityVideos})
	}
	return videosWithGif
}

func filterVideos(links []string) []string {
	// filter link end with .mp4
	videos := make([]string, 0)
	for _, link := range links {
		if strings.HasSuffix(link, ".mp4") {
			videos = append(videos, link)
		}
	}
	return videos
}

func getLowQualityVideo(video string) string {
	domains := strings.Split(video, "/")
	for index, domain := range domains {
		if domain == "highlight" {
			domains[index] = "preview"
			break
		}
	}
	return strings.Join(domains, "/")
}
