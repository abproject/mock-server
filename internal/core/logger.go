package core

import (
	"sort"
	"sync"
	"time"
)

var instanceLogger ILogger
var onceLogger sync.Once

type ILogger interface {
	Info(module string, message string)
	Warn(module string, error error)
	Error(module string, error error)
	Size() map[string]int
	GetInfo() []LoggerInfo
	GetWarn() []LoggerFails
	GetError() []LoggerFails
	GetAll() []LoggerEntry
}

type LoggerInfo struct {
	module    string
	message   string
	timestamp time.Time
}

type LoggerFails struct {
	module    string
	error     error
	timestamp time.Time
}

type LoggerEntry struct {
	level     string
	module    string
	message   string
	error     error
	timestamp time.Time
}

type Logger struct {
	logs   []LoggerInfo
	warns  []LoggerFails
	errors []LoggerFails
}

func GetLogger() ILogger {
	onceLogger.Do(func() {
		instanceLogger = newLogger()
	})
	return instanceLogger
}

func newLogger() ILogger {
	return &Logger{
		logs:   []LoggerInfo{},
		warns:  []LoggerFails{},
		errors: []LoggerFails{},
	}
}

func (logger *Logger) Info(module string, message string) {
	logger.logs = append(logger.logs, LoggerInfo{
		module:    module,
		message:   message,
		timestamp: time.Now(),
	})
}

func (logger *Logger) Warn(module string, error error) {
	logger.warns = append(logger.warns, LoggerFails{
		module:    module,
		error:     error,
		timestamp: time.Now(),
	})
}

func (logger *Logger) Error(module string, error error) {
	logger.errors = append(logger.errors, LoggerFails{
		module:    module,
		error:     error,
		timestamp: time.Now(),
	})
}

func (logger *Logger) Size() map[string]int {
	return map[string]int{
		"info":  len(logger.logs),
		"warn":  len(logger.warns),
		"error": len(logger.errors),
	}
}

func (logger *Logger) GetInfo() []LoggerInfo {
	return append(logger.logs[:0:0], logger.logs...)
}

func (logger *Logger) GetWarn() []LoggerFails {
	return append(logger.warns[:0:0], logger.warns...)
}

func (logger *Logger) GetError() []LoggerFails {
	return append(logger.errors[:0:0], logger.errors...)
}

func (logger *Logger) GetAll() []LoggerEntry {
	var entries []LoggerEntry
	entries = append(entries, mapLoggerInfo(logger.logs)...)
	entries = append(entries, mapLoggerFails(logger.warns, "warn")...)
	entries = append(entries, mapLoggerFails(logger.errors, "error")...)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].timestamp.Before(entries[j].timestamp)
	})
	return entries
}

func mapLoggerInfo(logs []LoggerInfo) []LoggerEntry {
	entries := make([]LoggerEntry, len(logs))
	for i, v := range logs {
		entries[i] = LoggerEntry{
			level:     "info",
			module:    v.module,
			message:   v.message,
			error:     nil,
			timestamp: v.timestamp,
		}
	}
	return entries
}

func mapLoggerFails(fails []LoggerFails, level string) []LoggerEntry {
	entries := make([]LoggerEntry, len(fails))
	for i, v := range fails {
		entries[i] = LoggerEntry{
			level:     level,
			module:    v.module,
			message:   "",
			error:     v.error,
			timestamp: v.timestamp,
		}
	}
	return entries
}
