package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
	"timetracker_cli/internal"
)

func TestCanDeleteSessionByIndex(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Add(internal.Session{0, 10, "test"})

	cmd := NewDeleteCmd(tracker)
	cmd.SetArgs([]string{"1"})
	cmd.Execute()

	if len(tracker.Sessions) != 0 {
		t.Errorf("Expected no sessions, got %d", len(tracker.Sessions))
	}
}

func TestShouldReturnErrorIfNoIndexProvided(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewDeleteCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) == "" {
		t.Error("Expected error, got none")
	}
}

func TestShouldReturnErrorIfIndexOutOfRange(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewDeleteCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetArgs([]string{"1"})
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) == "" {
		t.Error("Expected error, got none")
	}
}

func TestShouldReturnErrorIfIndexNonNumeric(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewDeleteCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetArgs([]string{"abc"})
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) == "" {
		t.Error("Expected error, got none")
	}
}
