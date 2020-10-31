package metrics

import "k8s.io/component-base/version"

var (
	buildInfo = NewGaugeVec(
		&GaugeOpts{
			Name:           "platform_build_info",
			Help:           "A metric with a constant '1' value labeled by major, minor, git version, git commit, git tree state, build date, " +
				"Go version, and compiler from which this library was built, and platform on which it is running.",
			StabilityLevel: ALPHA,
		},
		[]string{"major", "minor", "gitVersion", "gitCommit", "gitTreeState", "buildDate", "goVersion", "compiler", "platform"},
	)
)

// RegisterBuildInfo registers the build and version info in a metadata metric in prometheus
func RegisterBuildInfo(r PlatformRegistry) {
	info := version.Get()
	r.MustRegister(buildInfo)
	buildInfo.WithLabelValues(info.Major, info.Minor, info.GitVersion, info.GitCommit, info.GitTreeState, info.BuildDate, info.GoVersion, info.Compiler, info.Platform).Set(1)
}
