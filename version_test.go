package main

import (
	"bytes"
	"testing"

	"github.com/giantswarm/crd-docs-generator/pkg/config"
	"github.com/giantswarm/crd-docs-generator/pkg/project"
)

const (
	configFlag    = "--config"
	missingConfig = "does-not-exist.yaml"
)

func TestVersionNeedsNoConfig(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
	}{
		{name: "subcommand", args: []string{"version"}},
		{name: "flag", args: []string{"--version"}},
		{name: "subcommand with missing config", args: []string{"version", configFlag, missingConfig}},
		{name: "flag with missing config", args: []string{"--version", configFlag, missingConfig}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Chdir(t.TempDir())

			var configFilePath string
			rootCommand := newRootCommand(&configFilePath)
			var output bytes.Buffer
			rootCommand.SetOut(&output)
			rootCommand.SetArgs(tc.args)

			if err := rootCommand.ExecuteContext(t.Context()); err != nil {
				t.Fatalf("execute %v: %v", tc.args, err)
			}
			want := "crd-docs-generator version " + project.VersionLine() + "\n"
			if got := output.String(); got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestRootCommandStillReadsConfig(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
	}{
		{name: "default path", args: nil},
		{name: "explicit path", args: []string{configFlag, missingConfig}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Chdir(t.TempDir())

			var configFilePath string
			rootCommand := newRootCommand(&configFilePath)
			rootCommand.SetOut(&bytes.Buffer{})
			rootCommand.SetErr(&bytes.Buffer{})
			rootCommand.SetArgs(tc.args)

			err := rootCommand.ExecuteContext(t.Context())
			if !config.IsCouldNotReadConfigFile(err) {
				t.Fatalf("execute %v: got error %v, want CouldNotReadConfigFileError", tc.args, err)
			}
		})
	}
}
