package metrics

import (
	"sync"
	"sync/atomic"

	"github.com/blang/semver"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	apimachineryversion "k8s.io/apimachinery/pkg/version"
	"k8s.io/component-base/version"
)

var (
	showHiddenOnce sync.Once
	showHidden     atomic.Value
)

// SetShowHidden will enable showing hidden metrics. This will no-opt
// after the initial call
func SetShowHidden() {
	showHiddenOnce.Do(func() {
		showHidden.Store(true)
	})
}

// ShouldShowHidden returns whether showing hidden deprecated metrics
// is enabled. While the primary usecase for this is internal (to determine
// registration behavior) this can also be used to introspect
func ShouldShowHidden() bool {
	return showHidden.Load() != nil && showHidden.Load().(bool)
}

// Registerable is an interface for a collector metric which we
// will register with PlatformRegistry.
type Registerable interface {
	prometheus.Collector
	Create(version *semver.Version) bool
}

// PlatformRegistry is an interface which implements a subset of prometheus.Registerer and
// prometheus.Gatherer interfaces
type PlatformRegistry interface {
	// Deprecated
	RawRegister(prometheus.Collector) error
	// Deprecated
	RawMustRegister(...prometheus.Collector)
	Register(Registerable) error
	MustRegister(...Registerable)
	Unregister(Registerable) bool
	Gather() ([]*dto.MetricFamily, error)
}

// BlackspacePlatformRegistry is a wrapper around a prometheus registry-type object. Upon initialization
// the binary version information and metadata is loaded into the registry object, so that
// automatic behavior can be configured for metric versioning.
type BlackspacePlatformRegistry struct {
	PromRegistry
	version semver.Version
}

// Register registers a new Collector to be included in metrics
// collection. It returns an error if the descriptors provided by the
// Collector are invalid or if they — in combination with descriptors of
// already registered Collectors — do not fulfill the consistency and
// uniqueness criteria described in the documentation of metric.Desc.
func (kr *BlackspacePlatformRegistry) Register(c Registerable) error {
	if c.Create(&kr.version) {
		return kr.PromRegistry.Register(c)
	}
	return nil
}

// MustRegister works like Register but registers any number of
// Collectors and panics upon the first registration that causes an
// error.
func (kr *BlackspacePlatformRegistry) MustRegister(cs ...Registerable) {
	metrics := make([]prometheus.Collector, 0, len(cs))
	for _, c := range cs {
		if c.Create(&kr.version) {
			metrics = append(metrics, c)
		}
	}
	kr.PromRegistry.MustRegister(metrics...)
}

// RawRegister takes a native prometheus.Collector and registers the collector
// to the registry. This bypasses metrics safety checks, so should only be used
// to register custom prometheus collectors.
//
// Deprecated
func (kr *BlackspacePlatformRegistry) RawRegister(c prometheus.Collector) error {
	return kr.PromRegistry.Register(c)
}

// RawMustRegister takes a native prometheus.Collector and registers the collector
// to the registry. This bypasses metrics safety checks, so should only be used
// to register custom prometheus collectors.
//
// Deprecated
func (kr *BlackspacePlatformRegistry) RawMustRegister(cs ...prometheus.Collector) {
	kr.PromRegistry.MustRegister(cs...)
}

// Unregister unregisters the Collector that equals the Collector passed
// in as an argument.  (Two Collectors are considered equal if their
// Describe method yields the same set of descriptors.) The function
// returns whether a Collector was unregistered. Note that an unchecked
// Collector cannot be unregistered (as its Describe method does not
// yield any descriptor).
func (kr *BlackspacePlatformRegistry) Unregister(collector Registerable) bool {
	return kr.PromRegistry.Unregister(collector)
}

// Gather calls the Collect method of the registered Collectors and then
// gathers the collected metrics into a lexicographically sorted slice
// of uniquely named MetricFamily protobufs. Gather ensures that the
// returned slice is valid and self-consistent so that it can be used
// for valid exposition. As an exception to the strict consistency
// requirements described for metric.Desc, Gather will tolerate
// different sets of label names for metrics of the same metric family.
func (kr *BlackspacePlatformRegistry) Gather() ([]*dto.MetricFamily, error) {
	return kr.PromRegistry.Gather()
}

func newPlatformRegistry(v apimachineryversion.Info) *BlackspacePlatformRegistry {
	r := &BlackspacePlatformRegistry{
		PromRegistry: prometheus.NewRegistry(),
		version:      parseVersion(v),
	}
	return r
}

// NewPlatformRegistry creates a new vanilla Registry without any Collectors
// pre-registered.
func NewPlatformRegistry() PlatformRegistry {
	r := newPlatformRegistry(version.Get())
	return r
}
