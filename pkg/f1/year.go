package f1

type Year struct {
	Year   int
	Winter Season
	Spring Season
	Summer Season
}

func (y Year) Str() string {
	return ""
}

func (y Year) Out() string {
	return ""
}
