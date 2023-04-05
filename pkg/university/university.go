package university

type University struct {
	name string
}

func (u University) ToString() string {
	return u.name
}