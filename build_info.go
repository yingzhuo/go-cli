package cli

// BuildInfo stores app build info
type BuildInfo struct {
	Timestamp   string
	GitBranch   string
	GitCommit   string
	GitRevCount string
	BuiltBy     string
}
