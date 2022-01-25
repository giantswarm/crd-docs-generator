package config

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *FromFile
		wantErr bool
	}{
		{
			name: "test 1 - read valid config file",
			args: args{
				path: "testdata/config1.yaml",
			},
			want: &FromFile{
				SourceRepositories: []SourceRepository{
					{
						URL:          "https://github.com/giantswarm/apiextensions",
						Organization: "giantswarm",
						ShortName:    "apiextensions",
						CRDPaths:     []string{"config/crd", "helm"},
						CRPaths:      []string{"docs/cr"},
						Metadata: map[string]CRDItem{
							"crd.with.full.info": {
								Owners:    []string{"owner"},
								Topics:    []string{"apps"},
								Providers: []string{"aws", "azure"},
								Hidden:    false,
							},
							"unpublished.crd": {
								Hidden: true,
							},
							"only.defaults": {
								Hidden: false,
							},
							"deprecated.crd": {
								Hidden: false,
								Deprecation: &Deprecation{
									ReplacedBy: &DeprecationReplacedBy{
										FullName:  "new.full.crd.name",
										ShortName: "New",
									},
								},
							},
							"simply.deprecated.crd": {
								Hidden: false,
								Deprecation: &Deprecation{
									Info: "This CRD is deprecated",
								},
							},
						},
						CommitReference: "v3.39.0",
					},
				},
				TemplatePath: "my/file",
			},
			wantErr: false,
		},
		{
			name: "test 2 - file does not exist",
			args: args{
				path: "testdata/foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test 3 - invalid file",
			args: args{
				path: "testdata/config2.yaml",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
