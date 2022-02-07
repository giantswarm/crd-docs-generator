package main

import "testing"

func Test_generateCrdDocs(t *testing.T) {
	type args struct {
		configFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				configFilePath: "testdata/case1/config.yaml",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := generateCrdDocs(tt.args.configFilePath); (err != nil) != tt.wantErr {
				t.Errorf("generateCrdDocs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
