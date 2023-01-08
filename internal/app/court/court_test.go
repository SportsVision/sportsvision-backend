package court

import (
	"os"
	"testing"
	"wxcloudrun-golang/internal/pkg/db"
)

func TestMain(m *testing.M) {
	db.Init()
	os.Exit(m.Run())
}

func TestService_GetCourts(t *testing.T) {
	type args struct {
		latitude  string
		longitude string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				latitude:  "50.915",
				longitude: "116.404",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService()
			_, err := s.GetCourts(tt.args.latitude, tt.args.longitude)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCourts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
