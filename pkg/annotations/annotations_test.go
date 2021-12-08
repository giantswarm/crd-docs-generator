package annotations

import (
	"reflect"
	"testing"
)

func Test_findFiles(t *testing.T) {
	type args struct {
		startPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{startPath: "testdata"},
			want:    []string{"testdata/aws.go"},
			wantErr: false,
		},
		{
			name:    "Non-existing path",
			args:    args{startPath: "dfsgdfggh"},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findFiles(tt.args.startPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("findFiles() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollect(t *testing.T) {
	type args struct {
		startPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []CRDAnnotationSupport
		wantErr bool
	}{
		{
			name: "Successful",
			args: args{
				startPath: "testdata",
			},
			want: []CRDAnnotationSupport{
				{
					Annotation:    "alpha.cni.aws.giantswarm.io/minimum-ip-target",
					CRDName:       "awsclusters.infrastructure.giantswarm.io",
					CRDVersion:    "v1alpha2",
					Release:       "Since 14.0.0",
					Documentation: "This annotation allows configuration of the MINIMUM_IP_TARGET parameter for AWS CNI.",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Collect(tt.args.startPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collect() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect() = %#v,\nwant %#v", got, tt.want)
			}
		})
	}
}
