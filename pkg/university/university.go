package university

type University struct {
	name string
}

func (u University) String() string {
	return u.name
}
