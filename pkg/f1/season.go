package f1

import "fmt"

type Season struct {
	Name     string
	Champion Driver
	Races    []Race
}

func (s Season) ToString() string {
	return fmt.Sprintf(
		"%s, %s",
		s.Name,
		s.Champion.Name)
}
