package racing

import "fmt"

type Time struct {
	Minutes      int
	Seconds      int
	Milliseconds int
}

func (t Time) Str() string {
	return fmt.Sprintf(
		"%d, %d, %d",
		t.Minutes,
		t.Seconds,
		t.Milliseconds,
	)
}

func (t Time) Out() string {
	return fmt.Sprintf(
		"%d:%d.%d",
		t.Minutes,
		t.Seconds,
		t.Milliseconds,
	)
}
