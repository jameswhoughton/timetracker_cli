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

func TestCanAddDescriptionWhenStarting(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	expected := "test description"

	cmd := NewStartCmd(tracker)
	cmd.SetArgs([]string{expected})
	cmd.Execute()

	if tracker.Current.Description == "" {
		t.Errorf("Expected description '%s' got ''", expected)
	}
}

func TestStartAcceptsOnlyOneArgument(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewStartCmd(tracker)
	cmd.SetArgs([]string{"a", "b"})
	b := bytes.NewBufferString("")
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) == "" {
		t.Error("Expected error, got ''")
	}

	if tracker.Current.Description == "a" {
		t.Error("Description shouldn't be set on error")
	}
}
