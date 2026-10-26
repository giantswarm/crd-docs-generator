// Package project carries the build identity of the crd-docs-generator binary:
// its version, git commit and build time. The generated Makefile.gen.go.mk and
// the architect CI orb stamp them through `-ldflags -X`; a plain `go build` or
// `go install` has no ldflags and falls back to the build info Go records from
// version control.
package project

import (
	"runtime/debug"
	"strings"
)

// Name is the binary name used in messages.
const Name = "crd-docs-generator"

// dev is what Version reports for a build that carries no version at all.
const dev = "dev"

// devel is the main module version runtime/debug reports for a build without
// a resolvable VCS tag.
const devel = "(devel)"

// Set at link time via `-X github.com/giantswarm/crd-docs-generator/pkg/project.<name>=…`.
// The CI orb stamps version only on tag builds, as the tag without its `v`.
var (
	version        string
	gitSHA         string
	buildTimestamp string
)

// Version returns the version in the form the tags use ("v0.11.4"): the
// linker-stamped one, else the main module version Go recorded, else "dev".
func Version() string {
	if stamped := strings.TrimSpace(version); stamped != "" {
		if stamped[0] >= '0' && stamped[0] <= '9' {
			stamped = "v" + stamped
		}
		return stamped
	}
	if recorded := readVCS().version; recorded != "" {
		return recorded
	}
	return dev
}

// GitSHA returns the commit the binary was built from: the linker-stamped
// one, else the revision Go recorded, else "".
func GitSHA() string {
	if gitSHA != "" {
		return gitSHA
	}
	return readVCS().revision
}

// ShortSHA is GitSHA cut to seven characters, or "" when unknown.
func ShortSHA() string {
	sha := GitSHA()
	if len(sha) > 7 {
		sha = sha[:7]
	}
	return sha
}

// BuildTimestamp returns the build time in RFC 3339: the linker-stamped one,
// else the commit time Go recorded, else "".
func BuildTimestamp() string {
	if buildTimestamp != "" {
		return buildTimestamp
	}
	return readVCS().time
}

// VersionLine is the version followed by the details the binary knows:
// "v0.11.4 (commit 361cdef, built 2026-10-11T09:12:00Z)".
func VersionLine() string {
	line := Version()
	var details []string
	if sha := ShortSHA(); sha != "" {
		details = append(details, "commit "+sha)
	}
	if timestamp := BuildTimestamp(); timestamp != "" {
		details = append(details, "built "+timestamp)
	}
	if len(details) > 0 {
		line += " (" + strings.Join(details, ", ") + ")"
	}
	return line
}

type vcs struct {
	version  string
	revision string
	time     string
}

// readVCS is a variable so tests can substitute the build info.
var readVCS = func() vcs {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return vcs{}
	}
	var recorded vcs
	if info.Main.Version != devel {
		recorded.version = info.Main.Version
	}
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			recorded.revision = setting.Value
		case "vcs.time":
			recorded.time = setting.Value
		}
	}
	return recorded
}
