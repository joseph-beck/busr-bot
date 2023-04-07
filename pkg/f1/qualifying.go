package f1

import "fmt"

// generate id by combining 9, race number (2 digits), year (2 digits) and the number (2 digits) of the times this qualifying appears, starting at 1
// for example 9012201
type Qualifying struct {
	Id     int
	Race   Race
	First  Driver
	Second Driver
	Third  Driver
}

func (q Qualifying) Str() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d, %d",
		q.Id,
		q.Race.Id,
		q.First.Id,
		q.Second.Id,
		q.Third.Id,
	)
}

func (q Qualifying) SqlStr() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d, %d",
		q.Id,
		q.Race.Id,
		q.First.Id,
		q.Second.Id,
		q.Third.Id,
	)
}

func (q Qualifying) Out() string {
	return fmt.Sprintf(
		"``` Track: %s\n Pole: %s\n Second: %s\n Third: %s\n ```",
		q.Race.Track.Str(),
		q.First.Name,
		q.Second.Name,
		q.Third.Name,
	)
}
