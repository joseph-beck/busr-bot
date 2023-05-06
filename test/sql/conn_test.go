package sql

import (
	"bot/internal/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	c := sql.MakeConn(path)

	assert.NotNil(t, c)
}
