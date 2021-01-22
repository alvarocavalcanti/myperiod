package main

import "errors"

// EntryType Enum
type EntryType int

// EntryType Enum Values
const (
	HighFlow = iota
	MediumFlow
	LowFlow
	UnprotectedIntercourse
	ProtectedIntercourse
)

// Entry ...
type Entry struct {
	ID 		int 		`json: "id"`
	Type 	EntryType 	`json: "type"`

}

func (p *Entry) getEntry() error {
  return errors.New("Not implemented")
}

func (p *Entry) updateEntry() error {
  return errors.New("Not implemented")
}

func (p *Entry) deleteEntry() error {
  return errors.New("Not implemented")
}

func (p *Entry) createEntry() error {
  return errors.New("Not implemented")
}

func getEntries(start, count int) ([]Entry, error) {
  return nil, errors.New("Not implemented")
}