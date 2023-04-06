package f1

type Year struct {
	Year   int
	Winter Season
	Spring Season
	Summer Season
}

func (y Year) ToString() string {
	return ""
}

func (y Year) PrettyString() string {
	return ""
}

