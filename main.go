package main

import (
	"errors"
	"time"
)

func main() {
	// start recording

	// stop recording

	// list times

	// add time

	// remove time

	// clear all
}

type Session struct {
	Start       int64
	End         int64
	Description string
}

type Tracker struct {
	Current  Session
	Sessions []Session
}

func (t *Tracker) Start() {
	t.Current.Start = time.Now().Unix()
}

func (t *Tracker) End() error {
	if t.Current.Description == "" {
		return errors.New("description cannot be empty")
	}

	t.Current.End = time.Now().Unix()

	t.Sessions = append(t.Sessions, t.Current)

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

	return nil
}

func (t *Tracker) DeleteByIndex(index int) error {
	if index < 0 || index >= len(t.Sessions) {
		return errors.New("index is invalid")
	}

	t.Sessions = append(t.Sessions[:index], t.Sessions[index+1:]...)

	return nil
}
