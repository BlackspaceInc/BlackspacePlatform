package base

// InstanceMetricType is used as metadata in metrics emitted. It details
// the type of client used to emit such metrics
type InstanceMetricType string

const (
	Service InstanceMetricType = "Service"
	Platform InstanceMetricType = "Platform"
	System InstanceMetricType = "System"
	Events InstanceMetricType = "Events"
	Database InstanceMetricType = "Database"
)

const DefaultInstanceName string = "BlackspacePlatform"
const DefaultInstanceMetricType InstanceMetricType = Platform
const DefaultInstanceVersion string = "0.0.1"

// InstanceType encapsulates an entities name and type. An entity can be defined
// as a service, a database, .... etc
type InstanceMetadata struct {
	// platform or service name
	InstanceName string
	// instance type
	InstanceType InstanceMetricType
	// instance version
	InstanceVersion string
}

// DefaultInstance provides a default instance initialization to client
func DefaultInstance() *InstanceMetadata {
	return &InstanceMetadata {
		InstanceName: DefaultInstanceName,
		InstanceType: DefaultInstanceMetricType,
		InstanceVersion: DefaultInstanceVersion,
	}
}

// NewInstance provides a customer instance initialization to client
func NewInstance(instanceName string, instanceType InstanceMetricType, instanceVersion string) *InstanceMetadata{
	if len(instanceType) == 0 {
		instanceType =  DefaultInstanceMetricType
	}

	return &InstanceMetadata{
		InstanceName: instanceName,
		InstanceType: instanceType,
		InstanceVersion: instanceVersion,
	}
}

