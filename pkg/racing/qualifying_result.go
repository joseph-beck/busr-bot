package racing

type QualifyingResultPrimitive struct {
	Id           int
	Driver       int
	Position     int
	Minutes      int
	Seconds      int
	Milliseconds int
}

type QualifyingResult struct {
	Id       int
	Driver   RaceWeekendPrimitve
	Position int
	Time     Time
}
