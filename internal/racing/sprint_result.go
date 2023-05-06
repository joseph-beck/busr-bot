package racing

import "fmt"

type SprintResultPrimitive struct {
	Id           int
	Driver       int
	Position     int
	Minutes      int
	Seconds      int
	Milliseconds int
	Points       float64
}

type SprintResult struct {
	Id       int
	Driver   Driver
	Position int
	Time     Time
	Points   float64
}

func (s *SprintResult) Str() string {
	return fmt.Sprintf(
		"Sprint Result\n Driver: %s\n Position: %d\n, Time: %s\n Points: %.2f \n",
		s.Driver.Name,
		s.Position,
		s.Time.Str(),
		s.Points,
	)
}

func (s *SprintResult) Out() string {
	return fmt.Sprintf(
		"\n Position: %d\n Time: %s\n Points: %.2f\n",
		s.Position,
		s.Time.Out(),
		s.Points,
	)
}
