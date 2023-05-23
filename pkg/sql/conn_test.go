package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const path = "../../configs/database.json"

func TestConnection(t *testing.T) {
	d := MakeDbx(path)
	
	assert.NotNil(t, d)
}
