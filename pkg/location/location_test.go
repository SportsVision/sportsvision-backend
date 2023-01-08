package location

import "testing"

func TestGetDistanceByLocation(t *testing.T) {
	type args struct {
		lat1 float64
		lng1 float64
		lat2 float64
		lng2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				lat1: 39.915,
				lng1: 116.404,
				lat2: 39.98,
				lng2: 116.404,
			},
			want: 7235.7669,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDistance(tt.args.lat1, tt.args.lng1, tt.args.lat2, tt.args.lng2); got != tt.want {
				t.Errorf("GetDistanceByLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
