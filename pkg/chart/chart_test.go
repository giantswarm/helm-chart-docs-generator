package chart

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRead(t *testing.T) {
	type args struct {
		basePath  string
		chartName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "file with one test chart with one version and some schema",
			args: args{
				basePath:  "testdata",
				chartName: "test-chart",
			},
			want:    "# Values schema documentation",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadChartMetadata(tt.args.basePath, tt.args.chartName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Read() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
