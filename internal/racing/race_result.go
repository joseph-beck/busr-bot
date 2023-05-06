package racing

import "fmt"

type RaceResultPrimitive struct {
	Id           int
	Driver       int
	Position     int
	Minutes      int
	Seconds      int
	Milliseconds int
	Points       float64
}

type RaceResult struct {
	Id       int
	Driver   Driver
	Position int
	Time     Time
	Points   float64
}

func (r *RaceResult) Str() string {
	return fmt.Sprintf(
		"Sprint Result\n Driver: %s\n Position: %d\n, Time: %s\n Points: %.2f \n",
		r.Driver.Name,
		r.Position,
		r.Time.Str(),
		r.Points,
	)
}

func (r *RaceResult) Out() string {
	return fmt.Sprintf(
		"\n Position: %d\n Time: %s\n Points: %.2f\n",
		r.Position,
		r.Time.Out(),
		r.Points,
	)
}
