package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
	"timetracker_cli/internal"

	"github.com/rodaine/table"
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

	if string(out) != "No sessions" {
		t.Errorf("Expected 'No sessions', got %s", out)
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

	expectedTable := table.New("Description", "Total")

	tableBuffer := bytes.NewBufferString("")
	expectedTable.WithWriter(tableBuffer)

	expectedTable.AddRow("TEST", "10s")
	expectedTable.AddRow("TEST 2", "15s")

	expectedTable.Print()

	expected, err := ioutil.ReadAll(tableBuffer)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(out, expected) {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, out)
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

	expectedTable := table.New("Description", "Total")

	tableBuffer := bytes.NewBufferString("")
	expectedTable.WithWriter(tableBuffer)

	expectedTable.AddRow("TEST", "25s")

	expectedTable.Print()

	expected, err := ioutil.ReadAll(tableBuffer)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(out, expected) {
		t.Errorf("Expected %s, got %s", expected, out)
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

	expectedTable := table.New("Description", "Total")
	expectedTable.AddRow("TEST", "10s")
	expectedTable.AddRow("TEST 2", "2m 8s")
	expectedTable.AddRow("TEST 3", "11h 6m 40s")

	tableBuffer := bytes.NewBufferString("")
	expectedTable.WithWriter(tableBuffer)
	expectedTable.Print()

	expected, err := ioutil.ReadAll(tableBuffer)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(out, expected) {
		t.Errorf("Expected %s, got %s", expected, out)
	}
}
