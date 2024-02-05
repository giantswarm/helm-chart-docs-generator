package output

import (
	"flag"
	"io"
	"os"
	"testing"

	"github.com/giantswarm/microerror"
	"github.com/google/go-cmp/cmp"

	"github.com/giantswarm/helm-chart-docs-generator/pkg/chart"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestWritePage(t *testing.T) {
	type args struct {
		content      string
		metadata     chart.Metadata
		repoURL      string
		repoRef      string
		templatePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		golden  string
	}{
		{
			name: "Test 01",
			args: args{
				content:      "This is the content",
				metadata:     chart.Metadata{Name: "Test Chart", Description: "This is a test chart", Version: "1.0.0"},
				repoURL:      "https://github.com/giantswarm/my-repo",
				repoRef:      "main",
				templatePath: "testdata/chart.template",
			},
			golden:  "test_01",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "TestWritePage")
			if err != nil {
				t.Fatalf("Could not create temp dir: %s", err)
			}
			defer os.RemoveAll(tempDir)

			resultPath, err := WritePage(tt.args.metadata, tt.args.content, tempDir, tt.args.repoURL, tt.args.repoRef, tt.args.templatePath)
			if err != tt.wantErr {
				t.Errorf("WritePage() error = %v, wantErr %v", err, tt.wantErr)
				t.Logf("%s", microerror.Pretty(err, true))
			}

			gotBytes, err := os.ReadFile(resultPath)
			if err != nil {
				t.Errorf("Could not open result file %s: %s", resultPath, err)
			}
			got := string(gotBytes)
			want := goldenValue(t, tt.golden, got, *update)

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("WritePage() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func goldenValue(t *testing.T, goldenFile string, actual string, update bool) string {
	t.Helper()
	goldenPath := "testdata/" + goldenFile + ".golden"

	f, err := os.OpenFile(goldenPath, os.O_RDWR, 0644)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	defer f.Close()

	if update {
		_, err := f.WriteString(actual)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Error reading content of file %s: %s", goldenPath, err)
	}
	return string(content)
}
