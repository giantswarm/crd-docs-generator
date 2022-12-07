package main

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func Test_generateCrdDocs(t *testing.T) {
	type args struct {
		configFilePath string
	}
	tests := []struct {
		name       string
		args       args
		golden     string
		outputFile string
		wantErr    bool
	}{
		{
			name: "case1",
			args: args{
				configFilePath: "testdata/case1/config.yaml",
			},
			golden:     "testdata/case1/output.golden",
			outputFile: "testdata/case1/output/examples.example.giantswarm.io.md",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := generateCrdDocs(tt.args.configFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateCrdDocs() error = %v, wantErr %v", err, tt.wantErr)
			}

			content, err := os.ReadFile(tt.outputFile)
			if err != nil {
				t.Fatalf("Error loading output file %s: %s", tt.outputFile, err)
			}
			got := string(content)

			want := goldenValue(t, tt.golden, got, *update)
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("generateCrdDocs() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func goldenValue(t *testing.T, goldenPath string, actual string, update bool) string {
	t.Helper()

	f, _ := os.OpenFile(goldenPath, os.O_RDWR, 0644)
	defer f.Close()

	if update {
		_, err := f.WriteString(actual)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	return string(content)
}
