package f1

import "fmt"

// generate id by combining 4, race number (2 digits starting at 1), year (2 digits) and the number of qualifying (2 digits starting at 1)
// for example 4012201
type Race struct {
	Id     int
	Track  Tracks
	Laps   int
	First  Driver
	Second Driver
	Third  Driver
	Season Season
}

func (r Race) Str() string {
	return fmt.Sprintf(
		"%d, %s, %d, %d, %d, %d, %d",
		r.Id,
		r.Track.Str(),
		r.Laps,
		r.First.Id,
		r.Second.Id,
		r.Third.Id,
		r.Season.Id,
	)
}

func (r Race) SqlStr() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d, %d, %d, %d",
		r.Id,
		r.Track,
		r.Laps,
		r.First.Id,
		r.Second.Id,
		r.Third.Id,
		r.Season.Id,
	)
}

func (r Race) Out() string {
	return fmt.Sprintf(
		"`` Race: %s\n Laps: %d\n Winner: %s\n Second: %s\n Third: %s\n Season: %s ``",
		r.Track.Str(),
		r.Laps,
		r.First.Name,
		r.Second.Name,
		r.Third.Name,
		r.Season.Season.Str(),
	)
}
