package f1

import (
	"bot/pkg/university"
)

type Driver struct {
	name       string
	university university.University
	wins       int
	points     float32
}
