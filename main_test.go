package main

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCanStartASession(t *testing.T) {
	tracker := Tracker{}

	tracker.Start()

	if tracker.Current.Start == 0 {
		t.Error("Expected start to be non-zero timestamp")
	}
}

func TestCanEndASessionThatHasStarted(t *testing.T) {
	tracker := Tracker{}

	tracker.Start()

	tracker.SetDescription("test")

	tracker.End()

	if len(tracker.Sessions) != 1 {
		t.Errorf("Expected 1 recorded session, got %d", len(tracker.Sessions))
	}
}

func TestEndedSessionContainsTheCorrectPeriodLength(t *testing.T) {
	tracker := Tracker{}

	tracker.SetDescription("test")

	tracker.Start()

	expected := 5

	time.Sleep(time.Duration(expected) * time.Second)

	tracker.End()

	lastSession := tracker.Sessions[len(tracker.Sessions)-1]

	sessionLength := lastSession.End - lastSession.Start

	if sessionLength != 5 {
		t.Errorf("Expected session length of %d got %d", expected, sessionLength)
	}
}

func TestCurrentShouldResetAfterEndingSession(t *testing.T) {
	tracker := Tracker{}

	tracker.Start()

	tracker.SetDescription("test")

	tracker.End()

	if tracker.Current.Start != 0 {
		t.Error("Expected current session to reset after completion")
	}
}

func TestCanManuallyAddSession(t *testing.T) {
	tracker := Tracker{}

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
	tracker := Tracker{}

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
	tracker := Tracker{}

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
	tracker := Tracker{}

	tracker.Start()

	tracker.SetDescription("test")

	if tracker.Current.Description != "test" {
		t.Errorf("Expected description 'test', got '%s'", tracker.Current.Description)
	}
}

func TestCurrentSessionCannotEndIfDescriptionIsEmpty(t *testing.T) {
	tracker := Tracker{}

	tracker.Start()

	err := tracker.End()

	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestCanDeleteSessionByIndex(t *testing.T) {
	tracker := Tracker{}

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
	tracker := Tracker{}

	err := tracker.DeleteByIndex(100)

	if err == nil {
		t.Error("Expected error got none")
	}
}

func TestDeleteIndexCannotBeNegative(t *testing.T) {
	tracker := Tracker{}

	err := tracker.DeleteByIndex(-1)

	if err == nil {
		t.Error("Expected error got none")
	}
}
