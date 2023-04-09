package racing

import "fmt"

type RaceWeekendPrimitive struct {
	Id         int
	Track      int
	Qualifying int
	Sprint     int
	Race       int
}

type RaceWeekend struct {
	Id         int
	Track      Circuit
	Qualifying QualifyingResult
	Sprint     SprintResult
	Race       RaceResult
}

func (r *RaceWeekendPrimitive) Str() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d, %d",
		r.Id,
		r.Track,
		r.Qualifying,
		r.Sprint,
		r.Race,
	)
}
