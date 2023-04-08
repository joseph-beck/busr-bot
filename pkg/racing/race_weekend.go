package racing

import "fmt"

type RaceWeekendPrimitive struct {
	Id         int
	Track      int
	Qualifying int
	Sprint     int
	Race       int
	Points     float32
}

type RaceWeekend struct {
	Id         int
	Track      Circuit
	Qualifying QualifyingResult
	Sprint     SprintResult
	Race       RaceResult
	Points     float32
}

func (r RaceWeekendPrimitive) Str() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d, %d, %.2f",
		r.Id,
		r.Track,
		r.Qualifying,
		r.Sprint,
		r.Race,
		r.Points,
	)
}
