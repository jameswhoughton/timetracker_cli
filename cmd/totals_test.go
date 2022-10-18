package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
	"timetracker_cli/internal"

	"github.com/rodaine/table"
)

func TestTotalsCmdReturnsMessageIfNoSessions(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewTotalsCmd(tracker, TotalsConfig{})
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) != "No sessions" {
		t.Errorf("Expected 'No sessions' got '%s'", out)
	}
}

func TestTotalsCmdReturnsTableOfTotals(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Add(internal.Session{
		Start:       0,
		End:         100,
		Description: "test",
	})

	tracker.Add(internal.Session{
		Start:       30,
		End:         40,
		Description: "test",
	})

	tracker.Add(internal.Session{
		Start:       60,
		End:         65,
		Description: "test 2",
	})

	cmd := NewTotalsCmd(tracker, TotalsConfig{})
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	expectedTable := table.New("Description", "Total")
	expectedTable.AddRow("test", "1m 50s")
	expectedTable.AddRow("test 2", "5s")

	tableBuffer := bytes.NewBufferString("")
	expectedTable.WithWriter(tableBuffer)
	expectedTable.Print()

	expected, err := ioutil.ReadAll(tableBuffer)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, out) {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, out)
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

	cmd := NewTotalsCmd(tracker, TotalsConfig{})
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

	cmd := NewTotalsCmd(tracker, TotalsConfig{})
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

func TestTotalsShouldIncludeRoundedIfRoundByGreaterThan0(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Add(internal.Session{
		Start:       0,
		End:         100,
		Description: "test",
	})

	tracker.Add(internal.Session{
		Start:       30,
		End:         40,
		Description: "test",
	})

	tracker.Add(internal.Session{
		Start:       60,
		End:         65,
		Description: "test 2",
	})

	cmd := NewTotalsCmd(tracker, TotalsConfig{RoundBy: 15})
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	expectedTable := table.New("Description", "Total", "Rounded")
	expectedTable.AddRow("test", "1m 50s", "1m 45s")
	expectedTable.AddRow("test 2", "5s", "15s")

	tableBuffer := bytes.NewBufferString("")
	expectedTable.WithWriter(tableBuffer)
	expectedTable.Print()

	expected, err := ioutil.ReadAll(tableBuffer)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, out) {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, out)
	}
}
