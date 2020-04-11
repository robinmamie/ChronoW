package timer

import (
	"errors"
	"sync"
	"time"
)

// TickDuration designates the time a tick takes to pass
const TickDuration time.Duration = 10 * time.Millisecond

// TicksPerSecond contains the number of ticks executed per second
const TicksPerSecond uint32 = uint32(time.Second / TickDuration)

// Timer defines the timer datastructure
type Timer struct {
	sync.Mutex

	totalTicks     uint32
	remainingTicks uint32
	fixedAlert     uint32

	// Internal channel used to stop the ticking
	stop chan (bool)

	// AlertEnd sends true when the timer has reached 0, or false when it is
	// any other alert
	AlertEnd chan (bool)
}

// NewTimer creates a new timer
func NewTimer(seconds, numberOfPeriods uint32) (*Timer, error) {

	totalTicks := seconds * TicksPerSecond
	fixedAlert := totalTicks
	if numberOfPeriods != 0 {
		if seconds%numberOfPeriods != 0 {
			return nil, errors.New("the number of periods does not divide the total duration")
		}
		fixedAlert = totalTicks / numberOfPeriods
	}

	return &Timer{
		totalTicks:     totalTicks,
		remainingTicks: totalTicks,
		fixedAlert:     fixedAlert,
		stop:           make(chan bool),
		AlertEnd:       make(chan bool),
	}, nil
}

// Start (re)starts the timer
func (t *Timer) Start() {
	go t.tick()
}

// Stop interrupts the timer
func (t *Timer) Stop() {
	t.stop <- true
}

// Reset stops and then resets the timer
func (t *Timer) Reset() {
	t.stop <- true
	t.Lock()
	t.remainingTicks = t.totalTicks
	t.Unlock()
}

// TenthOfSecondsRemaining return the amount of tenth of seconds remaining in
// the timer
func (t *Timer) TenthOfSecondsRemaining() uint32 {
	return t.remainingTicks / (TicksPerSecond / 10)
}

// PeriodNumber returns the period number, aka how many alerts + 1 have already
// been called.
func (t *Timer) PeriodNumber() uint32 {
	rounding := uint32(1)
	if t.remainingTicks == 0 {
		rounding = 0
	}
	return ((t.totalTicks - t.remainingTicks) / t.fixedAlert) + rounding
}

// tick is the internal ticking logic for the timer
func (t *Timer) tick() {
	ticker := time.NewTicker(TickDuration)
	for {
		select {
		case <-ticker.C:
			t.Lock()
			t.remainingTicks--
			t.Unlock()
		case <-t.stop:
			ticker.Stop()
			return
		}
		if t.remainingTicks%t.fixedAlert == 0 {
			// Timer has to stop when an alert (or the end) has been reached
			t.AlertEnd <- (t.remainingTicks == 0)
			return
		}
	}
}
