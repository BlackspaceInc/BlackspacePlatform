package metrics

import (
	"fmt"
	"github.com/blang/semver"
	apimachineryversion "k8s.io/apimachinery/pkg/version"
	"regexp"
)

const (
	versionRegexpString = `^v(\d+\.\d+\.\d+)`
)

var (
	versionRe = regexp.MustCompile(versionRegexpString)
)

func parseSemver(s string) *semver.Version {
	if s != "" {
		sv := semver.MustParse(s)
		return &sv
	}

	return nil
}

func parseVersion(ver apimachineryversion.Info) semver.Version {
	matches := versionRe.FindAllStringSubmatch(ver.String(), -1)

	if len(matches) != 1 {
panic(fmt.Sprintf("version string \"%v\" doesn't match expected regular expression: \"%v\"", ver.String(), versionRe.String()))
}
	return semver.MustParse(matches[0][1])
}
