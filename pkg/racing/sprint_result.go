package racing

type SprintResultPrimitive struct {
	Id           int
	Driver       int
	Position     int
	Minutes      int
	Seconds      int
	Milliseconds int
}

type SprintResult struct {
	Id       int
	Driver   RaceWeekendPrimitve
	Position int
	Time     Time
}
