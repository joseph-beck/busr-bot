package racing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRaceResult(t *testing.T) {
	r := RaceResult{
		Id: 123,
		Driver: Driver{
			Id:         456,
			Name:       "Driver",
			University: "University",
		},
		Position: 12,
		Time: Time{
			Minutes:      1,
			Seconds:      23,
			Milliseconds: 672,
		},
		Points: 23.25,
	}

	assert.NotNil(t, r)
}
