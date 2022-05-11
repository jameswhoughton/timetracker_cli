package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sort"
	"time"
)

type Session struct {
	Start       int64
	End         int64
	Description string
}

type Total struct {
	Description string
	Total       int
}

type trackerClock interface {
	Now() int64
}

type realClock struct{}

func (tc realClock) Now() int64 {
	return time.Now().Unix()
}

type trackerData struct {
	Current  Session   `json:"current"`
	Sessions []Session `json:"sessions"`
	Totals   []Total   `json:"totals"`
}

type TrackerConfig struct {
	File string
}

func NewTracker(config TrackerConfig) *Tracker {
	return &Tracker{
		config: config,
		clock:  realClock{},
	}
}

type Tracker struct {
	config TrackerConfig
	clock  trackerClock
	trackerData
}

func (t *Tracker) Start() {
	t.Current.Start = t.clock.Now()
}

func (t *Tracker) End() error {
	t.Current.End = t.clock.Now()

	err := t.Add(t.Current)

	if err != nil {
		return err
	}

	t.Current = Session{}

	return nil
}

func (t *Tracker) SetDescription(description string) {
	t.Current.Description = description
}

func (t *Tracker) Add(session Session) error {
	if session.End-session.Start <= 0 {
		return errors.New("end time must be greater than start")
	}

	if session.Description == "" {
		return errors.New("description cannot be empty")
	}

	t.Sessions = append(t.Sessions, session)

	t.addTotal(session.Description, int(session.End)-int(session.Start))

	sort.Slice(t.Sessions, func(i, j int) bool {
		if t.Sessions[i].Start == t.Sessions[j].Start {
			return t.Sessions[i].End < t.Sessions[j].End
		}

		return t.Sessions[i].Start < t.Sessions[j].Start
	})

	return nil
}

func (t *Tracker) addTotal(description string, inc int) {
	for i, total := range t.Totals {
		if total.Description == description {
			t.Totals[i].Total += inc
			return
		}
	}

	t.Totals = append(t.Totals, Total{
		Description: description,
		Total:       inc,
	})
}

func (t *Tracker) DeleteByIndex(index int) error {
	if index < 0 || index >= len(t.Sessions) {
		return errors.New("index is invalid")
	}

	t.Sessions = append(t.Sessions[:index], t.Sessions[index+1:]...)

	return nil
}

func (t *Tracker) DeleteAll() {
	t.Sessions = []Session{}
}

func (t *Tracker) Save() {
	fileContent, _ := json.Marshal(t)

	ioutil.WriteFile(t.config.File, fileContent, 0644)
}

func (t *Tracker) Restore() {
	fileContents, _ := ioutil.ReadFile(t.config.File)

	restoredData := trackerData{}

	json.Unmarshal([]byte(fileContents), &restoredData)

	t.trackerData = restoredData
}
