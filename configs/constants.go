package configs

//XAuthTokenKey header key
const XAuthTokenKey = "X-Auth-Token"

//XSubjectTokenKey header key
const XSubjectTokenKey = "X-Subject-Token"

//DefaultAvailabilityZone declare default zone
const DefaultAvailabilityZone = "nova"

// Endpoint service type
type Endpoint int

var (
	// Nova endpoint item
	Nova Endpoint = 0
	// Cinder endpoint item
	Cinder Endpoint = 1
	// Keystone endpoint item
	Keystone Endpoint = 2
	// Neutron endpoint item
	Neutron Endpoint = 3
	// NovaLegacy endpoint item
	NovaLegacy Endpoint = 4
	// Glance endpoint item
	Glance Endpoint = 5
	// CinderV3 endpoint item
	CinderV3 Endpoint = 6
	// CinderV2 endpoint item
	CinderV2 Endpoint = 7
	// Placement endpoint item
	Placement Endpoint = 8
)

// GetEntrypoint return string
func GetEntrypoint(e Endpoint) string {
	endpointMap := map[Endpoint]string{
		0: "nova",
		1: "cinder",
		2: "keystone",
		3: "neutron",
		4: "nova_legacy",
		5: "glance",
		6: "cinderv3",
		7: "cinderv2",
		8: "placement",
	}
	return endpointMap[e]
}
