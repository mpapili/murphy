package occ

// OccurrenceType catalog — expandable catalog of in-world events.
type OccurrenceType string

const (
	TypeWeather OccurrenceType = "weather"
	TypeAnomaly OccurrenceType = "anomaly"
	TypeSpawn   OccurrenceType = "spawn"
	TypeTimed   OccurrenceType = "timed"
)

// AllTypes is the starter set.
func AllTypes() []OccurrenceType {
	return []OccurrenceType{TypeWeather, TypeAnomaly, TypeSpawn, TypeTimed}
}
