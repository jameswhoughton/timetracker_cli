package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
	"timetracker_cli/internal"

	"github.com/rodaine/table"
)

func TestReturnMessageIfNoSessions(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewListCmd(tracker)
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

func TestListSessionsInTable(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Add(internal.Session{
		Start:       0,
		End:         100,
		Description: "test",
	})

	tracker.Add(internal.Session{
		Start:       30,
		End:         40,
		Description: "test 2",
	})

	tracker.Add(internal.Session{
		Start:       60,
		End:         65,
		Description: "test 3",
	})

	cmd := NewListCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	expectedTable := table.New("ID", "Description", "Start", "End", "Length")
	expectedTable.AddRow("1", "test", "01/01/70 00:00:00", "01/01/70 00:01:40", "1m 40s")
	expectedTable.AddRow("2", "test 2", "01/01/70 00:00:30", "01/01/70 00:00:40", "10s")
	expectedTable.AddRow("3", "test 3", "01/01/70 00:01:00", "01/01/70 00:01:05", "5s")

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
