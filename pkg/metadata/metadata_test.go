package metadata

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
		want    *Root
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{"testdata/1.yaml"},
			want: &Root{
				CRDs: map[string]CRDItem{
					"crd.with.full.info": {
						Owner:     []string{"owner"},
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
				},
			},
			wantErr: false,
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
