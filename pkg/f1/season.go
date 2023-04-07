package f1

import "fmt"

type Season struct {
	Id       string
	Name     string
	Champion Driver
	Races    []Race
}

func (s Season) Str() string {
	return fmt.Sprintf(
		"%s, %s",
		s.Name,
		s.Champion.Name,
	)
}

func (s Season) RacesStr() string {
	output := ""

	for _, race := range s.Races {
		output += fmt.Sprintf(
			"%s \n",
			race.Name,
		)
	}

	return output
}
