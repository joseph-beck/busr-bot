package racing

import "fmt"

type QualifyingResultPrimitive struct {
	Id           int
	Driver       int
	Position     int
	Minutes      int
	Seconds      int
	Milliseconds int
	Points       float64
}

type QualifyingResult struct {
	Id       int
	Driver   Driver
	Position int
	Time     Time
	Points   float64
}

func (q *QualifyingResult) Str() string {
	return fmt.Sprintf(
		"Qualifying Result\n Driver: %s\n Position: %d\n, Time: %s\n Points: %.2f \n",
		q.Driver.Name,
		q.Position,
		q.Time.Str(),
		q.Points,
	)
}

func (q *QualifyingResult) Out() string {
	return fmt.Sprintf(
		"\n Position: %d\n Time: %s\n Points: %.2f\n",
		q.Position,
		q.Time.Out(),
		q.Points,
	)
}
