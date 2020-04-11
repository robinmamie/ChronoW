package timer_test

import (
	"testing"
	"time"

	. "github.com/robinmamie/ChronoW/internal/timer"
)

const REF_TOTAL_SECONDS uint32 = 360

// Describes the time format the timer outputs (here, tenth of seconds)
const BASE_UNIT_MULTIPLIER uint32 = 10

func testTime(t *testing.T, total, duration uint32) *Timer {
	expected := total*BASE_UNIT_MULTIPLIER - duration
	myTimer, err := NewTimer(total, 0)
	if err != nil {
		t.Errorf(err.Error())
		return myTimer
	}
	myTimer.Start()
	refTimer := time.NewTicker(time.Millisecond * time.Duration(100*duration))
	select {
	case <-refTimer.C:
		myTimer.Stop()
		got := myTimer.TenthOfSecondsRemaining()
		if got != expected {
			t.Errorf("TenthOfSecondsRemaining() = %d; want %d", got, expected)
		}
	}
	return myTimer
}

func TestTwoTenthSecond(t *testing.T) {
	testTime(t, REF_TOTAL_SECONDS, 2)
}

func TestOneSecondAndFinishAlert(t *testing.T) {
	myTimer, _ := NewTimer(1, 0)
	myTimer.Start()
	time.Sleep(time.Millisecond * 1001)
	select {
	case val := <-myTimer.AlertEnd:
		if !val {
			t.Errorf("did not signal end through AlertEnd with a true value")
		}
	default:
		t.Errorf("did not signal end through AlertEnd channel")
	}
}

func TestAlmostOneSecondAndNoFinishAlert(t *testing.T) {
	myTimer, _ := NewTimer(1, 0)
	myTimer.Start()
	// Test precision: 1 millisecond
	time.Sleep(time.Millisecond * 999)
	select {
	case <-myTimer.AlertEnd:
		t.Errorf("signaled end even though not finished")
	default:
	}
}

func TestOneSecondAndPeriodAlert(t *testing.T) {
	myTimer, _ := NewTimer(2, 2)
	myTimer.Start()
	time.Sleep(time.Millisecond * 1001)
	select {
	case val := <-myTimer.AlertEnd:
		if val {
			t.Errorf("did not signal end of period through AlertEnd with a false value")
		}
	default:
		t.Errorf("did not signal end of period through AlertEnd channel")
	}
}

func TestAlmostOneSecondAndNoPeriodAlert(t *testing.T) {
	myTimer, _ := NewTimer(2, 2)
	myTimer.Start()
	// Test precision: 1 millisecond
	time.Sleep(time.Millisecond * 999)
	select {
	case <-myTimer.AlertEnd:
		t.Errorf("signaled end even though period not finished")
	default:
	}
}

func TestStopTimerWorksWithOneTenthSecond(t *testing.T) {
	myTimer := testTime(t, REF_TOTAL_SECONDS, 1)
	time.Sleep(time.Millisecond * 400)
	if myTimer.TenthOfSecondsRemaining() != REF_TOTAL_SECONDS*10-1 {
		t.Errorf("did not stop the timer even though Stop() was called")
	}
}

func TestTimerReset(t *testing.T) {
	myTimer, _ := NewTimer(REF_TOTAL_SECONDS, 1)
	myTimer.Start()
	time.Sleep(time.Second)
	myTimer.Reset()
	if myTimer.TenthOfSecondsRemaining() != REF_TOTAL_SECONDS*10 {
		t.Errorf("did not reset the timer correctly even though Reset() was called")
	}
}

func TestTimerRefusesIncoherentPeriods(t *testing.T) {
	for j := uint32(1); j <= 2*REF_TOTAL_SECONDS; j++ {
		for i := uint32(1); i <= j; i++ {
			_, err := NewTimer(j, i)
			if (err != nil) == (j%i == 0) {
				t.Errorf("timer accepts %d periods for %d seconds", i, j)
				return
			}
		}
	}
}

func TestPeriodNumber(t *testing.T) {
	myTimer, _ := NewTimer(2, 2)
	myTimer.Start()
	period := myTimer.PeriodNumber()
	if period != 1 {
		t.Errorf("timer does not show correct number for start of 1st period: %d", period)
	}
	time.Sleep(time.Millisecond * 500)
	period = myTimer.PeriodNumber()
	if period != 1 {
		t.Errorf("timer does not show correct number for 1st period: %d", period)
	}

	time.Sleep(time.Millisecond * 501)
	period = myTimer.PeriodNumber()
	if period != 2 {
		t.Errorf("timer does not show correct number before start of 2nd period: %d", period)
	}
	myTimer.Start()

	time.Sleep(time.Millisecond * 500)
	period = myTimer.PeriodNumber()
	if period != 2 {
		t.Errorf("timer does not show correct number for middle of 2nd period: %d", period)
	}
	time.Sleep(time.Millisecond * 501)
	period = myTimer.PeriodNumber()
	if period != 2 {
		t.Errorf("timer does not show correct number for end of 2nd period/bout: %d", period)
	}
}
