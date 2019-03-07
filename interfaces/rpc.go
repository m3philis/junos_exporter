package interfaces

type InterfaceRpc struct {
	Information struct {
		Interfaces []PhyInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type PhyInterface struct {
	Name              string         `xml:"name"`
	AdminStatus       string         `xml:"admin-status"`
	OperStatus        string         `xml:"oper-status"`
	Description       string         `xml:"description"`
	MacAddress        string         `xml:"current-physical-address"`
	Stats             TrafficStat    `xml:"traffic-statistics"`
	LogicalInterfaces []LogInterface `xml:"logical-interface"`
	InputErrors       struct {
		Drops  uint64 `xml:"input-drops"`
		Errors uint64 `xml:"input-errors"`
	} `xml:"input-error-list"`
	OutputErrors struct {
		Drops  uint64 `xml:"output-drops"`
		Errors uint64 `xml:"output-errors"`
	} `xml:"output-error-list"`
}

type LogInterface struct {
	Name        string         `xml:"name"`
	Description string         `xml:"description"`
	Stats       TrafficStat    `xml:"traffic-statistics"`
	LagStats    LagTrafficStat `xml:"lag-traffic-statistics"`
}

type TrafficStat struct {
	InputBytes    uint64   `xml:"input-bytes"`
	InputPackets  uint64   `xml:"input-packets"`
	OutputBytes   uint64   `xml:"output-bytes"`
	OutputPackets uint64   `xml:"output-packets"`
	IPv6Traffic   IPv6Stat `xml:"ipv6-transit-statistics"`
}

type IPv6Stat struct {
	InputBytes    uint64 `xml:"input-bytes"`
	InputPackets  uint64 `xml:"input-packets"`
	OutputBytes   uint64 `xml:"output-bytes"`
	OutputPackets uint64 `xml:"output-packets"`
}

type LagTrafficStat struct {
	Stats TrafficStat `xml:"lag-bundle"`
	Links []struct {
		Name string `xml:"name"`
	} `xml:"lag-link"`
}
