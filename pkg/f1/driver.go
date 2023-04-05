package f1

import (
	"bot/pkg/university"
	"fmt"
)

type Driver struct {
	Id            int
	Name          string
	University    university.University
	Wins          int
	Poles         int
	Podiums       int
	Points        float32
	Starts        int
	Avg_quali     float32
	Championships []Season
}

func (d Driver) ToString() string {
	return fmt.Sprintf(
		"%d, %s, %s, %d, %d, %d, %f, %d, %f",
		d.Id,
		d.Name,
		d.University.ToString(),
		d.Wins,
		d.Poles,
		d.Podiums,
		d.Points,
		d.Starts,
		d.Avg_quali)
}
