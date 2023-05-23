package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	path = "../../configs/database.json"

	id  = 262333533480026112
	qid = 29110123
	rid = 24110123
	sid = 25110123
	wid = 21110123
)

func TestInteractions(t *testing.T) {

}

func TestDriver(t *testing.T) {
	c := MakeConn(path)
	d := c.Driver(id)

	assert.NotNil(t, d)
}

func TestCheckDriver(t *testing.T) {
	c := MakeConn(path)
	e := c.CheckDriver(id)

	assert.True(t, e)

	e = c.CheckDriver(123)

	assert.False(t, e)
}

func TestQualifying(t *testing.T) {
	c := MakeConn(path)
	q := c.QualifyingResult(qid, id)

	assert.NotNil(t, q)
}

func TestCheckQualifying(t *testing.T) {
	c := MakeConn(path)
	e := c.CheckQualifyingResult(qid, id)

	assert.True(t, e)

	e = c.CheckQualifyingResult(123, id)

	assert.False(t, e)
}

func TestSprint(t *testing.T) {
	c := MakeConn(path)
	s := c.SprintResult(sid, id)

	assert.NotNil(t, s)
}

func TestCheckSprint(t *testing.T) {
	c := MakeConn(path)
	e := c.CheckSprintResult(sid, id)

	assert.True(t, e)

	e = c.CheckSprintResult(123, id)

	assert.False(t, e)
}

func TestRace(t *testing.T) {
	c := MakeConn(path)
	r := c.RaceResult(rid, id)

	assert.NotNil(t, r)
}

func TestCheckRace(t *testing.T) {
	c := MakeConn(path)
	e := c.CheckRaceResult(rid, id)

	assert.True(t, e)

	e = c.CheckRaceResult(123, id)

	assert.False(t, e)
}

func TestWeekend(t *testing.T) {
	c := MakeConn(path)
	w := c.WeekendPrimitive(wid)

	assert.NotNil(t, w)
}

func TestCheckWeekend(t *testing.T) {
	c := MakeConn(path)
	e := c.CheckWeekend(wid)

	assert.True(t, e)

	e = c.CheckWeekend(123)

	assert.False(t, e)
}
