package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

type session struct {
	Start       int64
	End         int64
	Description string
}

type trackerClock interface {
	Now() int64
}

type realClock struct{}

func (tc realClock) Now() int64 {
	return time.Now().Unix()
}

type trackerData struct {
	Current  session   `json:"current"`
	Sessions []session `json:"sessions"`
}

type trackerConfig struct {
	file string
}

type tracker struct {
	config trackerConfig
	clock  trackerClock
	trackerData
}

func (t *tracker) Start() {
	t.Current.Start = t.clock.Now()
}

func (t *tracker) End() error {
	if t.Current.Description == "" {
		return errors.New("description cannot be empty")
	}

	t.Current.End = t.clock.Now()

	t.Sessions = append(t.Sessions, t.Current)

	t.Current = session{}

	return nil
}

func (t *tracker) SetDescription(description string) {
	t.Current.Description = description
}

func (t *tracker) Add(session session) error {
	if session.End-session.Start <= 0 {
		return errors.New("end time must be greater than start")
	}

	if session.Description == "" {
		return errors.New("description cannot be empty")
	}

	t.Sessions = append(t.Sessions, session)

	return nil
}

func (t *tracker) DeleteByIndex(index int) error {
	if index < 0 || index >= len(t.Sessions) {
		return errors.New("index is invalid")
	}

	t.Sessions = append(t.Sessions[:index], t.Sessions[index+1:]...)

	return nil
}

func (t *tracker) DeleteAll() {
	t.Sessions = []session{}
}

func (t *tracker) Save() {
	fileContent, _ := json.Marshal(t)

	ioutil.WriteFile(t.config.file, fileContent, 0644)
}

func (t *tracker) Restore() {
	fileContents, _ := ioutil.ReadFile(t.config.file)

	restoredData := trackerData{}

	json.Unmarshal([]byte(fileContents), &restoredData)

	t.trackerData = restoredData
}
