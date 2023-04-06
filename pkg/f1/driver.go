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

func (d Driver) String() string {
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
		d.Avg_quali)
}

func (d Driver) Output() string {
	return fmt.Sprintf(
		"Id: %d, Name: %s, University: %s, Wins: %d, Poles: %d, Podiums %d, Starts: %d, Points: %.2f, Average Qualifying: %.2f",
		d.Id,
		d.Name,
		d.University,
		d.Wins,
		d.Poles,
		d.Podiums,
		d.Starts,
		d.Points,
		d.Avg_quali)
}
