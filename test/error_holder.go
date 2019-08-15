package test

import (
	"errors"
	"strings"
	"testing"
)

type errorMessage struct {
	message string
	level   int
}

type errorGroup struct {
	name   string
	errors []errorMessage
}

// ErrorHolder Error Holder object
type ErrorHolder struct {
	data []*errorGroup
}

// NewErrorHolder Error Holder
func NewErrorHolder() ErrorHolder {
	return ErrorHolder{
		data: []*errorGroup{},
	}
}

// Group Error Holder creationg group
func (holder *ErrorHolder) Group(groupName string) func(message string, level int) *errorGroup {
	group, err := holder.getByGroupName(groupName)
	if err != nil {
		group = &errorGroup{
			name:   groupName,
			errors: []errorMessage{},
		}
		holder.data = append(holder.data, group)
	}

	return func(message string, level int) *errorGroup {
		group.errors = append(group.errors, errorMessage{
			message,
			level,
		})
		return group
	}
}

func (holder *ErrorHolder) getByGroupName(groupName string) (*errorGroup, error) {
	for _, g := range holder.data {
		if g.name == groupName {
			return g, nil
		}
	}
	return &errorGroup{}, errors.New("Not Found")
}

// HasErrors Error Holder has errors
func (holder *ErrorHolder) HasErrors() bool {
	for _, group := range holder.data {
		if len(group.errors) > 0 {
			return true
		}
	}
	return false
}

// Print Print all errors
func (holder *ErrorHolder) Print(t *testing.T) {
	message := "\n"
	for _, group := range holder.data {
		if len(group.errors) > 0 {
			message += group.name + "\n"
			for _, error := range group.errors {
				message += strings.Repeat("\t", error.level) + error.message + "\n"
			}
			message += "\n"
		}
	}
	t.Error(message)
}
