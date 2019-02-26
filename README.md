# Clock (including Tardis)

Provides a `Clock` that returns a `time`.
This is useful for testing.
Instead of using `time` directly, use a `Clock` instead.
There are various clocks available:
- `SystemDefaultZone().Now()` which is equivalent to `time.Now()`
- `System(z *time.Location).Now()` which is equivalent to `time.Now().In(z)`
- `SystemUTC().Now()` which is equivalent to `time.Now().UTC()`

The most powerful thing for testing is a clock called `Tardis`.
As you could expect from the name, this clock allows you to change the time and location at will.

Happy Mocking of Time!
