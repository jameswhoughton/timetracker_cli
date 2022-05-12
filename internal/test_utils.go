package internal

type StubClock struct {
	time int64
	inc  int64
}

func (sc *StubClock) Now() int64 {
	sc.time += sc.inc
	return sc.time
}

func NewStubClock(inc int64) *StubClock {
	return &StubClock{
		inc: inc,
	}
}

func NewTestTracker(clock trackerClock, config TrackerConfig) *Tracker {
	return &Tracker{
		clock:  clock,
		config: config,
	}
}
