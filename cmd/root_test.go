package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
	"timetracker_cli/internal"
)

func TestRootCmdShouldReturnMessageIfNoSessions(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewRootCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) != "No Sessions" {
		t.Errorf("Expected 'No Sessions', got %s", out)
	}
}

func TestRootCmdShouldListSessions(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Add(internal.Session{
		Start:       0,
		End:         10,
		Description: "TEST",
	})

	tracker.Add(internal.Session{
		Start:       30,
		End:         45,
		Description: "TEST 2",
	})

	cmd := NewRootCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	expected := `TEST - 10s
TEST 2 - 15s`

	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}

func TestSessionsWithTheSameDescriptionShouldBeMerged(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Add(internal.Session{
		Start:       0,
		End:         10,
		Description: "TEST",
	})

	tracker.Add(internal.Session{
		Start:       30,
		End:         45,
		Description: "TEST",
	})

	cmd := NewRootCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	expected := "TEST - 25s"

	if string(out) != expected {
		t.Errorf("Expected %s, got %s", expected, string(out))
	}
}

func TestTotalsShouldBeHumanReadable(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Add(internal.Session{
		Start:       0,
		End:         10,
		Description: "TEST",
	})

	tracker.Add(internal.Session{
		Start:       0,
		End:         128,
		Description: "TEST 2",
	})

	tracker.Add(internal.Session{
		Start:       0,
		End:         40000,
		Description: "TEST 3",
	})

	cmd := NewRootCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	expected := `TEST - 10s
TEST 2 - 2m 8s
TEST 3 - 11h 6m 40s`

	if string(out) != expected {
		t.Errorf("Expected %s, got %s", expected, string(out))
	}
}
