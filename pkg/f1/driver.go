package f1

import (
	"fmt"
)

type Driver struct {
	Id         int
	Name       string
	University string
	Wins       int
	Poles      int
	Podiums    int
	Starts     int
	Points     float64
	Avg_quali  float64
}

func (d Driver) Str() string {
	return fmt.Sprintf(
		"%d, %s, %s, %d, %d, %d, %d, %f, %f",
		d.Id,
		d.Name,
		d.University,
		d.Wins,
		d.Poles,
		d.Podiums,
		d.Starts,
		d.Points,
		d.Avg_quali,
	)
}

func (d Driver) SqlStr() string {
	return fmt.Sprintf(
		"%d, '%s', '%s', %d, %d, %d, %d, %f, %f",
		d.Id,
		d.Name,
		d.University,
		d.Wins,
		d.Poles,
		d.Podiums,
		d.Starts,
		d.Points,
		d.Avg_quali,
	)
}

func (d Driver) UpdateStr() string {
	return fmt.Sprintf(
		"name='%s', university='%s', wins=%d, poles=%d, podiums=%d, starts=%d, points=%f, avg_quali=%f",
		d.Name,
		d.University,
		d.Wins,
		d.Poles,
		d.Podiums,
		d.Starts,
		d.Points,
		d.Avg_quali,
	)
}

func (d Driver) Out() string {
	return fmt.Sprintf(
		"``` Name: %s\n University: %s\n Wins: %d\n Poles: %d\n Podiums %d\n Starts: %d\n Points: %.2f\n Average Qualifying: %.2f```",
		d.Name,
		d.University,
		d.Wins,
		d.Poles,
		d.Podiums,
		d.Starts,
		d.Points,
		d.Avg_quali,
	)
}
