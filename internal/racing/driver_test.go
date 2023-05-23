package racing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDriver(t *testing.T) {
	d := Driver{
		Id:         123,
		Name:       "Driver",
		University: "University",
	}

	assert.NotNil(t, d)
}
