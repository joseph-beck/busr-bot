package f1

import "fmt"

type Time struct {
	Minutes     int
	Seconds     int
	Miliseconds int
}

func (t Time) Str() string {
	return fmt.Sprintf(
		"%d:%d.%d",
		t.Minutes,
		t.Seconds,
		t.Miliseconds,
	)
}