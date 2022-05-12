package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
	"timetracker_cli/internal"
)

func TestCanStartSession(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewStartCmd(tracker)
	cmd.Execute()

	if tracker.Current.Start == 0 {
		t.Error("New Sesson hasn't started")
	}
}

func TestReturnMessageIfCurrentSessionAlreadyStarted(t *testing.T) {
	clock := internal.NewStubClock(5)

	tracker := internal.NewTestTracker(clock, internal.TrackerConfig{})

	tracker.Start()

	currentStartValue := tracker.Current.Start

	cmd := NewStartCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	expected := "Session has already started"

	if string(out) != expected {
		t.Errorf("expected '%s', got '%s'", expected, out)
	}

	if currentStartValue != tracker.Current.Start {
		t.Errorf("Expected start time to be %d, got %d", currentStartValue, tracker.Current.Start)
	}
}
