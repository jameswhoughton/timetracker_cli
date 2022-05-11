package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
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

func newTestTracker(clock trackerClock, config TrackerConfig) *Tracker {
	return &Tracker{
		clock:  clock,
		config: config,
	}
}

func TestCanStartASession(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	tracker.Start()

	if tracker.Current.Start == 0 {
		t.Error("Expected start to be non-zero timestamp")
	}
}

func TestCanEndASessionThatHasStarted(t *testing.T) {
	clock := &StubClock{
		inc: 5,
	}

	tracker := newTestTracker(clock, TrackerConfig{})

	tracker.Start()

	tracker.SetDescription("test")

	err := tracker.End()

	if err != nil {
		t.Fatal(err)
	}

	if len(tracker.Sessions) != 1 {
		t.Errorf("Expected 1 recorded Session, got %d", len(tracker.Sessions))
	}
}

func TestEndedSessionContainsTheCorrectPeriodLength(t *testing.T) {
	clock := &StubClock{
		inc: 5,
	}

	tracker := newTestTracker(clock, TrackerConfig{})

	tracker.SetDescription("test")

	tracker.Start()

	expected := 5

	tracker.End()

	lastSession := tracker.Sessions[len(tracker.Sessions)-1]

	SessionLength := lastSession.End - lastSession.Start

	if SessionLength != 5 {
		t.Errorf("Expected Session length of %d got %d", expected, SessionLength)
	}
}

func TestCurrentShouldResetAfterEndingSession(t *testing.T) {
	clock := &StubClock{
		inc: 5,
	}

	tracker := newTestTracker(clock, TrackerConfig{})

	tracker.Start()

	tracker.SetDescription("test")

	err := tracker.End()

	if err != nil {
		t.Fatal(err)
	}

	if tracker.Current.Start != 0 {
		t.Error("Expected current Session to reset after completion")
	}
}

func TestCanManuallyAddSession(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	Session := Session{
		Start:       10,
		End:         100,
		Description: "test",
	}

	tracker.Add(Session)

	if len(tracker.Sessions) != 1 {
		t.Errorf("Expected 1 recorded Session, got %d", len(tracker.Sessions))
	}
}

func TestEndTimeMustBeAfterStartTime(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	Session := Session{
		Start:       1000,
		End:         100,
		Description: "test",
	}

	err := tracker.Add(Session)

	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestManualSessionMustHaveDescription(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	Session := Session{
		Start: 0,
		End:   100,
	}

	err := tracker.Add(Session)

	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestCanSetDescriptionForCurrentSession(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	tracker.Start()

	tracker.SetDescription("test")

	if tracker.Current.Description != "test" {
		t.Errorf("Expected description 'test', got '%s'", tracker.Current.Description)
	}
}

func TestCurrentSessionCannotEndIfDescriptionIsEmpty(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	tracker.Start()

	err := tracker.End()

	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestCanDeleteSessionByIndex(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

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
		t.Errorf("Expected 1 remaining Session, got %d", len(tracker.Sessions))
	}

	if !cmp.Equal(tracker.Sessions[0], remainingSession) {
		t.Errorf("Expected remaining Session: %+v, got: %+v", remainingSession, tracker.Sessions[0])
	}
}

func TestCannotDeleteIfIndexIsOutOfRange(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	err := tracker.DeleteByIndex(100)

	if err == nil {
		t.Error("Expected error got none")
	}
}

func TestDeleteIndexCannotBeNegative(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	err := tracker.DeleteByIndex(-1)

	if err == nil {
		t.Error("Expected error got none")
	}
}

func TestCanDeleteAllSessions(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

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
		t.Errorf("Expected no Sessions, got %d", len(tracker.Sessions))
	}
}

func TestCanWriteToFile(t *testing.T) {
	fileName := "test.json"

	config := TrackerConfig{
		File: fileName,
	}

	tracker := NewTracker(config)

	tracker.Add(Session{
		Start:       10,
		End:         20,
		Description: "test 1",
	})

	tracker.Start()

	tracker.SetDescription("Test")

	tracker.Save()

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected file test.json to exist")
	}

	file, _ := ioutil.ReadFile(fileName)

	defer os.Remove(fileName)

	expected := trackerData{
		Current:  tracker.Current,
		Sessions: tracker.Sessions,
		Totals:   tracker.Totals,
	}

	got := trackerData{}

	json.Unmarshal([]byte(file), &got)

	if !cmp.Equal(expected, got) {
		t.Errorf("Expected %+v got %+v", expected, got)
	}
}

func TestCanSetPathOfSaveFile(t *testing.T) {
	fileName := "test_file_name.xyz"

	config := TrackerConfig{
		File: fileName,
	}

	tracker := NewTracker(config)

	tracker.Save()

	defer os.Remove(fileName)

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected file %s to exist", fileName)
	}
}

func TestCanRestoreFromFile(t *testing.T) {
	currentSession := Session{
		Start:       123,
		End:         5676,
		Description: "dfgebsddsfsdf",
	}

	Sessions := []Session{
		{
			Start:       50,
			End:         60,
			Description: "TEST",
		},
		{
			Start:       80,
			End:         100,
			Description: "fff",
		},
	}

	Totals := []Total{
		{"TEST", 10},
		{"fff", 20},
	}

	data := trackerData{
		Current:  currentSession,
		Sessions: Sessions,
		Totals:   Totals,
	}

	fileContent, _ := json.Marshal(data)

	fileName := "test.json"

	ioutil.WriteFile(fileName, fileContent, 0644)

	defer os.Remove(fileName)

	config := TrackerConfig{
		File: fileName,
	}

	tracker := NewTracker(config)

	tracker.Restore()

	if !cmp.Equal(tracker.trackerData, data) {
		t.Errorf("Expected %+v, got %+v", data, tracker.trackerData)
	}
}

func TestSessionsAreOrderedByStartTime(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	tracker.Add(Session{
		Start:       10,
		End:         20,
		Description: "test",
	})

	tracker.Add(Session{
		Start:       5,
		End:         30,
		Description: "test 2",
	})

	tracker.Add(Session{
		Start:       100,
		End:         105,
		Description: "test 3",
	})

	expected := "test 2testtest 3"
	var got string

	for _, session := range tracker.Sessions {
		got += session.Description
	}

	if expected != got {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestSessionsShouldOrderByEndTimeIfStartTimeTheSame(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	tracker.Add(Session{
		Start:       5,
		End:         20,
		Description: "test",
	})

	tracker.Add(Session{
		Start:       100,
		End:         120,
		Description: "test 2",
	})

	tracker.Add(Session{
		Start:       100,
		End:         105,
		Description: "test 3",
	})

	expected := "testtest 3test 2"
	var got string

	for _, session := range tracker.Sessions {
		got += session.Description
	}

	if expected != got {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestTotalsShouldMergeSessionsWithSameDescription(t *testing.T) {
	tracker := NewTracker(TrackerConfig{})

	tracker.Add(Session{
		Start:       0,
		End:         10,
		Description: "TEST",
	})

	tracker.Add(Session{
		Start:       30,
		End:         45,
		Description: "TEST",
	})

	if tracker.Totals[0].Total != 25 {
		t.Errorf("Expected 25, got %d", tracker.Totals[0].Total)
	}
}
