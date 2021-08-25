package output

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/giantswarm/microerror"
	"github.com/google/go-cmp/cmp"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/crd-docs-generator/pkg/metadata"
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
		crd          apiextensionsv1.CustomResourceDefinition
		annotations  []CRDAnnotationSupport
		md           metadata.CRDItem
		crFolder     string
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
				crd: apiextensionsv1.CustomResourceDefinition{
					TypeMeta: metav1.TypeMeta{
						Kind:       "CustomResourceDefinition",
						APIVersion: "apiextensions.k8s.io/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name: "demos.demo.giantswarm.io",
					},
					Spec: apiextensionsv1.CustomResourceDefinitionSpec{
						Group: "demo.giantswarm.io",
						Names: apiextensionsv1.CustomResourceDefinitionNames{
							Plural:     "demos",
							Singular:   "demo",
							ShortNames: []string{"dmo"},
							Kind:       "Demo",
							ListKind:   "DemoList",
							Categories: []string{"first", "second"},
						},
						Scope: "Namespaced",
						Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
							{
								Name:    "v1alpha1",
								Served:  true,
								Storage: true,
								Schema: &apiextensionsv1.CustomResourceValidation{
									OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
										Description: "Main description",
										Type:        "object",
										Format:      "",
										Title:       "Object title",
										Required:    []string{"first_property"},
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"first_property": {
												Description: "First property",
												Type:        "string",
												Format:      "",
												Title:       "The title",
											},
										},
									},
								},
							},
						},
					},
				},
				md: metadata.CRDItem{
					Owners:    []string{"first-owner", "second-owner"},
					Topics:    []string{"first-topic", "second-topic"},
					Providers: []string{"aws", "azure"},
					Deprecation: &metadata.Deprecation{
						Info: "This is some deprecation info",
						ReplacedBy: &metadata.DeprecationReplacedBy{
							FullName:  "another.demo.giantswarm.io",
							ShortName: "Another",
						},
					},
				},
				crFolder:     "testdata",
				repoURL:      "https://github.com/giantswarm/my-repo",
				repoRef:      "main",
				templatePath: "testdata/crd.template",
			},
			golden:  "test_01",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := ioutil.TempDir("", "TestWritePage")
			if err != nil {
				t.Fatalf("Could not create temp dir: %s", err)
			}
			defer os.RemoveAll(tempDir)

			resultPath, err := WritePage(tt.args.crd, tt.args.annotations, tt.args.md, tt.args.crFolder, tempDir, tt.args.repoURL, tt.args.repoRef, tt.args.templatePath)
			if err != tt.wantErr {
				t.Errorf("WritePage() error = %v, wantErr %v", err, tt.wantErr)
				t.Logf("%s", microerror.Pretty(err, true))
			}

			gotBytes, err := ioutil.ReadFile(resultPath)
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

	content, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Error reading content of file %s: %s", goldenPath, err)
	}
	return string(content)
}
