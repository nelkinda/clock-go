package clock

import "time"

// Clock provides access to the current instant, date and time, optionally using a time-zone.
type Clock interface {
	Now() time.Time
}

// Tardis is a special type of clock that can travel by telling which other clock it shows.
type Tardis struct {
	Clock
	clock Clock
}

type fixedClock struct {
	z *time.Location
	t time.Time
}

func (w fixedClock) Now() time.Time {
	t := w.t
	if w.z != nil {
		t = t.In(w.z)
	}
	return t
}

type wallClock struct {
	z *time.Location
}

func (s wallClock) Now() time.Time {
	now := time.Now()
	if s.z != nil {
		now = now.In(s.z)
	}
	return now
}

func (t Tardis) Now() time.Time {
	if t.clock == nil {
		return time.Now().UTC()
	}
	return t.clock.Now()
}

// Fixed creates a clock that will always return the same time at the same location.
func Fixed(t time.Time, z *time.Location) Clock {
	return &fixedClock{z: z, t: t}
}

// System creates a clock that will always return the system time at the specified location.
func System(z *time.Location) Clock {
	return &wallClock{z: z}
}

// SystemUTC creates a clock that will always return the system time in UTC.
func SystemUTC() Clock {
	return &wallClock{z: time.UTC}
}

// SystemDefaultZone creates a clock that will always return the system time in the system's default zone.
func SystemDefaultZone() Clock {
	return &wallClock{}
}

// NewTardis creates a clock that will return the time of whatever Clock is used as reference of the Tardis.
// The clock can be changed, and if it is nil, the Tardis will use time.Now().UTC().
func NewTardis(initialClock Clock) Tardis {
	return Tardis{clock: initialClock}
}
