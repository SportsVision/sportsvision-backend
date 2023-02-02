package event

import (
	"os"
	"testing"
	"wxcloudrun-golang/internal/pkg/db"
)

var s Service

func TestMain(m *testing.M) {
	db.Init()
	os.Exit(m.Run())
}

func TestService_GetEventVideos(t *testing.T) {
	type args struct {
		openID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{openID: "ogDJL5R996RQKkZm0VtFaK2-i3ik"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetEventVideos(tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
