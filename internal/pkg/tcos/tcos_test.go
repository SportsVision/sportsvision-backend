package tcos

import "testing"

func TestGetCosFileList(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				prefix: "highlight/court1/20230103/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCosFileList(tt.args.prefix)
		})
	}
}
