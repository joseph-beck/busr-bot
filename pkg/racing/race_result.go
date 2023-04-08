package racing

type RaceResultPrimitive struct {
	Id           int
	Driver       int
	Position     int
	Minutes      int
	Seconds      int
	Milliseconds int
}

type RaceResult struct {
	Id       int
	Driver   RaceWeekendPrimitve
	Position int
	Time     Time
}
