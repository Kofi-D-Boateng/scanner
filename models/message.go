package models

type CriticalLevel int

const (
	Low CriticalLevel = iota
	Medium
	High
	Critical
)

type DiscoveredItem struct {
	Number      int32
	Name        string
	Level       CriticalLevel
	Message     string
	DateFound   string
	DateScanned string
	Mitigation  string
}

type Vulnerability struct {
	DiscoveredItem
	CVEAvailable bool
}

type OpenPort struct {
	DiscoveredItem
}