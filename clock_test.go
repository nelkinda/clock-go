package clock

import (
	"github.com/christianhujer/assert"
	"testing"
	"time"
)

func TestSystemUTC(t *testing.T) {
	before := time.Now().UTC()
	systemUTC := SystemUTC().Now()
	after := time.Now().UTC()
	_ = assert.False(t, before.After(systemUTC))
	_ = assert.False(t, after.Before(systemUTC))
}

func TestSystemDefaultZone(t *testing.T) {
	before := time.Now()
	system := SystemDefaultZone().Now()
	after := time.Now()
	_ = assert.False(t, before.After(system))
	_ = assert.False(t, after.Before(system))
}

func TestSystem(t *testing.T) {
	ist := time.FixedZone("IST", 19800)
	before := time.Now().In(ist)
	system := System(ist).Now()
	after := time.Now().In(ist)
	_ = assert.False(t, before.After(system))
	_ = assert.False(t, after.Before(system))
}

func TestFixedWithoutZone(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		panic(err)
	}
	fixed := Fixed(timestamp, nil)
	_ = assert.Equals(t, timestamp, fixed.Now())
}

func TestFixedInUTC(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		panic(err)
	}
	fixed := Fixed(timestamp, time.UTC)
	_ = assert.Equals(t, timestamp, fixed.Now())
}

func TestFixedInCustomZone(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+05:30")
	if err != nil {
		panic(err)
	}
	fixed := Fixed(timestamp, time.FixedZone("IST", 19800))
	_ = assert.True(t, timestamp.Equal(fixed.Now()))
}

func TestTardis(t *testing.T) {
	now := time.Now().UTC()
	tardis := NewTardis(Fixed(now, time.UTC))
	_ = assert.True(t, now.Equal(tardis.Now()))
	newNow, err := time.Parse(time.RFC3339,  "2006-01-02T15:04:05+05:30")
	if err != nil { panic(err) }
	tardis.clock = Fixed(newNow, time.FixedZone("IST", 19800))
	_ = assert.True(t, newNow.Equal(tardis.Now()))
}
