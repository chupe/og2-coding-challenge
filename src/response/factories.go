package response

import "time"

type Factories struct {
	Type              string
	Level             int
	RatePerMinute     int
	UnderConstruction bool
	TimeToFinish      time.Time
}
