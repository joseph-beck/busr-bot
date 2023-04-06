package f1

import "fmt"

type Race struct {
	Id     string
	Name   string
	Laps   int
	Winner Driver
}

func (r Race) String() string {
	return fmt.Sprintf(
		"%s, %d, %s",
		r.Name,
		r.Laps,
		r.Winner.Name)
}

func (r Race) Out() string {
	return fmt.Sprintf(
		"Race Name: %s, Laps: %d, Winner: %s",
		r.Name,
		r.Laps,
		r.Winner.Name)
}
