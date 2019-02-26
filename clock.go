package clock

import "time"

type Clock interface {
	Now() time.Time
}

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
		return time.Now()
	}
	return t.clock.Now()
}

func Fixed(t time.Time, z *time.Location) Clock {
	return &fixedClock{z: z, t: t}
}

func System(z *time.Location) Clock {
	return &wallClock{z: z}
}

func SystemUTC() Clock {
	return &wallClock{z: time.UTC}
}

func SystemDefaultZone() Clock {
	return &wallClock{}
}

func NewTardis(initialClock Clock) Tardis {
	return Tardis{clock: initialClock}
}
