package project

import (
	"testing"
)

const (
	testTag       = "v0.11.4"
	testCommit    = "361cdefdb8c6e2e0e940def05b4b026ee524cff3"
	testTimestamp = "2026-10-11T09:12:00Z"
)

func stub(t *testing.T, stampedVersion, stampedSHA, stampedTimestamp string, recorded vcs) {
	t.Helper()
	origVersion, origSHA, origTimestamp, origReadVCS := version, gitSHA, buildTimestamp, readVCS
	t.Cleanup(func() {
		version, gitSHA, buildTimestamp, readVCS = origVersion, origSHA, origTimestamp, origReadVCS
	})
	version, gitSHA, buildTimestamp = stampedVersion, stampedSHA, stampedTimestamp
	readVCS = func() vcs { return recorded }
}

func assertEqual(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestVersion(t *testing.T) {
	for _, tc := range []struct {
		name     string
		stamped  string
		recorded vcs
		want     string
	}{
		{name: "nothing known", want: dev},
		{name: "ldflag without v gets it", stamped: "0.11.4", want: testTag},
		{name: "ldflag with v is kept", stamped: testTag, want: testTag},
		{name: "ldflag is trimmed", stamped: " 0.12.0-rc.1 ", want: "v0.12.0-rc.1"},
		{name: "ldflag wins over build info", stamped: "0.12.0", recorded: vcs{version: testTag}, want: "v0.12.0"},
		{name: "build info supplies the tag", recorded: vcs{version: testTag}, want: testTag},
		{name: "commit alone is no version", recorded: vcs{revision: testCommit}, want: dev},
	} {
		t.Run(tc.name, func(t *testing.T) {
			stub(t, tc.stamped, "", "", tc.recorded)
			assertEqual(t, tc.want, Version())
		})
	}
}

func TestCommitAndTimestamp(t *testing.T) {
	t.Run("stamped values win", func(t *testing.T) {
		stub(t, "", "abc1234def", "2026-10-12T00:00:00Z", vcs{revision: testCommit, time: testTimestamp})
		assertEqual(t, "abc1234def", GitSHA())
		assertEqual(t, "abc1234", ShortSHA())
		assertEqual(t, "2026-10-12T00:00:00Z", BuildTimestamp())
	})

	t.Run("build info is the fallback", func(t *testing.T) {
		stub(t, "", "", "", vcs{revision: testCommit, time: testTimestamp})
		assertEqual(t, testCommit, GitSHA())
		assertEqual(t, "361cdef", ShortSHA())
		assertEqual(t, testTimestamp, BuildTimestamp())
	})

	t.Run("unknown is empty", func(t *testing.T) {
		stub(t, "", "", "", vcs{})
		assertEqual(t, "", GitSHA())
		assertEqual(t, "", ShortSHA())
		assertEqual(t, "", BuildTimestamp())
	})
}

func TestVersionLine(t *testing.T) {
	t.Run("all details", func(t *testing.T) {
		stub(t, "0.11.4", testCommit, testTimestamp, vcs{})
		assertEqual(t, "v0.11.4 (commit 361cdef, built 2026-10-11T09:12:00Z)", VersionLine())
	})

	t.Run("no details", func(t *testing.T) {
		stub(t, "", "", "", vcs{})
		assertEqual(t, "dev", VersionLine())
	})
}
