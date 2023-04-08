package f1

import "fmt"

// generate id by combining 9, race number (2 digits), year (2 digits) and the season number (01 for spring etc) then followed by 1
// for example 90122011
type Qualifying struct {
	Id      int
	Track   Tracks
	Results QualifyingResults
	Timings QualifyingTimings
}

// generate id by combining 9, race number (2 digits), year (2 digits) and the season number (01 for spring etc) then followed by 2
// for example 90122013
type QualifyingResults struct {
	Id     int
	First  Driver
	Second Driver
	Third  Driver
}

// generate id by combining 9, race number (2 digits), year (2 digits) and the season number (01 for spring etc) then followed by 3
// for example 90122012
type QualifyingTimings struct {
	Id     int
	First  Time
	Second Time
	Third  Time
}

func (q Qualifying) Str() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d, %d",
		q.Id,
		q.Track,
		q.Results.First.Id,
		q.Results.Second.Id,
		q.Results.Third.Id,
	)
}

func (q Qualifying) SqlStr() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d, %d",
		q.Id,
		q.Track,
		q.Results.First.Id,
		q.Results.Second.Id,
		q.Results.Third.Id,
	)
}

func (q Qualifying) Out() string {
	return fmt.Sprintf(
		"``` **Track**: %s\n **Pole**: %s (%s)\n **Second**: %s (%s)\n **Third**: %s (%s)\n ```",
		q.Track.Str(),
		q.Results.First.Name, q.Timings.First.Str(),
		q.Results.Second.Name, q.Timings.Second.Str(),
		q.Results.Third.Name, q.Timings.Third.Str(),
	)
}
