package f1

import "fmt"

// season id generated by combining 5, the season number (2 digits, sprint is 1, summer 2, winter 3) and year
// for example 50123
type Season struct {
	Id       int
	Year     int
	Season   Seasons
	Champion Driver
}

func (s Season) Str() string {
	return fmt.Sprintf(
		"%d, %s, %s",
		s.Id,
		s.Season.Str(),
		s.Champion.Name,
	)
}

func (s Season) SqlStr() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d",
		s.Id,
		s.Year,
		s.Season,
		s.Champion.Id,
	)
}

func (s Season) Out() string {
	return fmt.Sprintf(
		"``` Season: %s %d\n Champion: %s ```",
		s.Season.Str(),
		s.Year,
		s.Champion.Name,
	)
}
