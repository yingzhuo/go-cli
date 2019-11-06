package cli

import (
	"testing"
)

func TestParseBuildInfo(t *testing.T) {
	input := `time:"Sat May 13 19:53:08 UTC 2017" branch:master commit:320279c patches:1234`

	info := ParseBuildInfo(input)

	if info.Timestamp != "Sat May 13 19:53:08 UTC 2017" {
		t.Error("parsed time is wrong")
	}
	if info.GitBranch != "master" {
		t.Error("parsed branch is wrong")
	}
	if info.GitCommit != "320279c" {
		t.Error("parsed commit is wrong")
	}
	if info.GitRevCount != "1234" {
		t.Error("parsed patches is wrong")
	}
}
