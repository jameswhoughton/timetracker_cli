package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type StubClock struct {
	time int64
	inc  int64
}

func (sc *StubClock) Now() int64 {
	sc.time += sc.inc
	return sc.time
}

func TestCanStartASession(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	tracker.Start()

	if tracker.Current.Start == 0 {
		t.Error("Expected start to be non-zero timestamp")
	}
}

func TestCanEndASessionThatHasStarted(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	tracker.Start()

	tracker.SetDescription("test")

	tracker.End()

	if len(tracker.Sessions) != 1 {
		t.Errorf("Expected 1 recorded session, got %d", len(tracker.Sessions))
	}
}

func TestEndedSessionContainsTheCorrectPeriodLength(t *testing.T) {
	tracker := Tracker{
		Clock: &StubClock{
			inc: 5,
		},
	}

	tracker.SetDescription("test")

	tracker.Start()

	expected := 5

	tracker.End()

	lastSession := tracker.Sessions[len(tracker.Sessions)-1]

	sessionLength := lastSession.End - lastSession.Start

	if sessionLength != 5 {
		t.Errorf("Expected session length of %d got %d", expected, sessionLength)
	}
}

func TestCurrentShouldResetAfterEndingSession(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	tracker.Start()

	tracker.SetDescription("test")

	tracker.End()

	if tracker.Current.Start != 0 {
		t.Error("Expected current session to reset after completion")
	}
}

func TestCanManuallyAddSession(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	session := Session{
		Start:       10,
		End:         100,
		Description: "test",
	}

	tracker.Add(session)

	if len(tracker.Sessions) != 1 {
		t.Errorf("Expected 1 recorded session, got %d", len(tracker.Sessions))
	}
}

func TestEndTimeMustBeAfterStartTime(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	session := Session{
		Start:       1000,
		End:         100,
		Description: "test",
	}

	err := tracker.Add(session)

	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestManualSessionMustHaveDescription(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	session := Session{
		Start: 0,
		End:   100,
	}

	err := tracker.Add(session)

	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestCanSetDescriptionForCurrentSession(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	tracker.Start()

	tracker.SetDescription("test")

	if tracker.Current.Description != "test" {
		t.Errorf("Expected description 'test', got '%s'", tracker.Current.Description)
	}
}

func TestCurrentSessionCannotEndIfDescriptionIsEmpty(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	tracker.Start()

	err := tracker.End()

	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestCanDeleteSessionByIndex(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	remainingSession := Session{
		Start:       0,
		End:         10,
		Description: "test",
	}

	tracker.Add(remainingSession)

	tracker.Add(Session{
		Start:       10,
		End:         20,
		Description: "test 2",
	})

	tracker.DeleteByIndex(1)

	if len(tracker.Sessions) != 1 {
		t.Errorf("Expected 1 remaining session, got %d", len(tracker.Sessions))
	}

	if !cmp.Equal(tracker.Sessions[0], remainingSession) {
		t.Errorf("Expected remaining session: %+v, got: %+v", remainingSession, tracker.Sessions[0])
	}
}

func TestCannotDeleteIfIndexIsOutOfRange(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	err := tracker.DeleteByIndex(100)

	if err == nil {
		t.Error("Expected error got none")
	}
}

func TestDeleteIndexCannotBeNegative(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	err := tracker.DeleteByIndex(-1)

	if err == nil {
		t.Error("Expected error got none")
	}
}

func TestCanDeleteAllSessions(t *testing.T) {
	tracker := Tracker{
		Clock: RealClock{},
	}

	tracker.Add(Session{
		Start:       10,
		End:         20,
		Description: "test 1",
	})

	tracker.Add(Session{
		Start:       10,
		End:         20,
		Description: "test 2",
	})

	tracker.Add(Session{
		Start:       10,
		End:         20,
		Description: "test 3",
	})

	tracker.DeleteAll()

	if len(tracker.Sessions) != 0 {
		t.Errorf("Expected no sessions, got %d", len(tracker.Sessions))
	}
}
