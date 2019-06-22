package core

import (
	"errors"
	"reflect"
	"testing"
)

func TestLoggerSingleton(t *testing.T) {
	logger1 := GetLogger()
	logger2 := GetLogger()
	if logger1 != logger2 {
		t.Errorf("Logger must be singletone")
	}
}

func TestLoggerNotSingleton(t *testing.T) {
	logger1 := newLogger()
	logger2 := newLogger()
	if logger1 == logger2 {
		t.Errorf("Logger must not be singletone")
	}
}

func TestLoggerInfo(t *testing.T) {
	logger := newLogger()

	logger.Info("test", "message")

	for k, v := range logger.Size() {
		if k == "info" && v != 1 {
			t.Errorf("Must be added 1 log entry")
		}
		if k != "info" && v != 0 {
			t.Errorf("Must be only 1 log entry")
		}
	}
}

func TestLoggerWarn(t *testing.T) {
	logger := newLogger()

	logger.Warn("test", errors.New("warning"))

	for k, v := range logger.Size() {
		if k == "warn" && v != 1 {
			t.Errorf("Must be added 1 warn entry")
		}
		if k != "warn" && v != 0 {
			t.Errorf("Must be only 1 warn entry")
		}
	}
}

func TestLoggerError(t *testing.T) {
	logger := newLogger()

	logger.Error("test", errors.New("error"))

	for k, v := range logger.Size() {
		if k == "error" && v != 1 {
			t.Errorf("Must be added 1 error entry")
		}
		if k != "error" && v != 0 {
			t.Errorf("Must be only 1 error entry")
		}
	}
}

func TestLoggerSize(t *testing.T) {
	logger := newLogger()
	logger.Info("test", "message")
	logger.Info("test", "message")
	logger.Warn("test", errors.New("warning"))
	logger.Error("test", errors.New("error"))
	logger.Error("test", errors.New("error"))
	logger.Error("test", errors.New("error"))

	size := logger.Size()

	if size["info"] != 2 {
		t.Errorf("Must be added 2 info entries")
	}
	if size["warn"] != 1 {
		t.Errorf("Must be added 1 warn entries")
	}
	if size["error"] != 3 {
		t.Errorf("Must be added 3 error entries")
	}
}

func TestLoggerGetInfo(t *testing.T) {
	logger := newLogger()
	logger.Info("test", "message")

	logs := logger.GetInfo()

	if len(logs) != 1 {
		t.Errorf("Must be added 1 info entry")
	}
	expected := LoggerInfo{
		module:    "test",
		message:   "message",
		timestamp: logs[0].timestamp,
	}
	if !reflect.DeepEqual(expected, logs[0]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			expected, logs[0])
	}
}

func TestLoggerGetInfoNoEffectWhenChangeReturnValues(t *testing.T) {
	logger := newLogger()
	logger.Info("test", "message")
	logs1 := logger.GetInfo()

	logs1[0].module = "new-module"

	logs2 := logger.GetInfo()
	if reflect.DeepEqual(logs1, logs2) {
		t.Errorf(`
Values must not be equal
Expected: %+v
Stored:	  %+v`,
			logs1, logs2)
	}
}

func TestLoggerGetWarn(t *testing.T) {
	logger := newLogger()
	logger.Warn("test", errors.New("warning"))

	logs := logger.GetWarn()

	if len(logs) != 1 {
		t.Errorf("Must be added 1 warn entry")
	}
	expected := LoggerFails{
		module:    "test",
		error:     errors.New("warning"),
		timestamp: logs[0].timestamp,
	}
	if !reflect.DeepEqual(expected, logs[0]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			expected, logs[0])
	}
}

func TestLoggerGetWarnNoEffectWhenChangeReturnValues(t *testing.T) {
	logger := newLogger()
	logger.Warn("test", errors.New("warning"))
	logs1 := logger.GetWarn()

	logs1[0].module = "new-module"

	logs2 := logger.GetWarn()
	if reflect.DeepEqual(logs1, logs2) {
		t.Errorf(`
Values must not be equal
Expected: %+v
Stored:	  %+v`,
			logs1, logs2)
	}
}

func TestLoggerGetError(t *testing.T) {
	logger := newLogger()
	logger.Error("test", errors.New("error"))

	logs := logger.GetError()

	if len(logs) != 1 {
		t.Errorf("Must be added 1 error entry")
	}
	expected := LoggerFails{
		module:    "test",
		error:     errors.New("error"),
		timestamp: logs[0].timestamp,
	}
	if !reflect.DeepEqual(expected, logs[0]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			expected, logs[0])
	}
}

func TestLoggerGetErrorNoEffectWhenChangeReturnValues(t *testing.T) {
	logger := newLogger()
	logger.Error("test", errors.New("error"))
	logs1 := logger.GetError()

	logs1[0].module = "new-module"

	logs2 := logger.GetError()
	if reflect.DeepEqual(logs1, logs2) {
		t.Errorf(`
Values must not be equal
Expected: %+v
Stored:	  %+v`,
			logs1, logs2)
	}
}

func TestLoggerGetAll(t *testing.T) {
	logger := newLogger()
	logger.Info("test", "message")
	logger.Warn("test", errors.New("warning"))
	logger.Error("test", errors.New("error"))
	logger.Info("test", "message2")

	entries := logger.GetAll()

	if len(entries) != 4 {
		t.Errorf("Must be added 4 entries")
	}
	expected1 := LoggerEntry{
		level:     "info",
		module:    "test",
		message:   "message",
		error:     nil,
		timestamp: entries[0].timestamp,
	}
	if !reflect.DeepEqual(expected1, entries[0]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			expected1, entries[0])
	}
	expected2 := LoggerEntry{
		level:     "warn",
		module:    "test",
		message:   "",
		error:     errors.New("warning"),
		timestamp: entries[1].timestamp,
	}
	if !reflect.DeepEqual(expected2, entries[1]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			expected2, entries[1])
	}
	expected3 := LoggerEntry{
		level:     "error",
		module:    "test",
		message:   "",
		error:     errors.New("error"),
		timestamp: entries[2].timestamp,
	}
	if !reflect.DeepEqual(expected3, entries[2]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			expected3, entries[2])
	}
	expected4 := LoggerEntry{
		level:     "info",
		module:    "test",
		message:   "message2",
		error:     nil,
		timestamp: entries[3].timestamp,
	}
	if !reflect.DeepEqual(expected4, entries[3]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			expected4, entries[3])
	}
}
