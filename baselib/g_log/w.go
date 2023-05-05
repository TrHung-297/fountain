package g_log

import (
	"context"
	"time"
)

// WithField allocates a new entry and adds a field to it.
// Debug, Print, Info, Warn, Error, Fatal or Panic must be then applied to
// this new returned entry.
// If you want multiple fields, use `WithFields`.
func (kg *g_log) WithField(key string, value interface{}) *g_entry {
	return &g_entry{Entry: kg.logger.WithField(key, value), level: kg.level}
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (kg *g_log) WithFields(fields Fields) *g_entry {
	return &g_entry{Entry: kg.logger.WithFields(fields.ToLrFields()), level: kg.level}
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (kg *g_log) WithError(err error) *g_entry {
	return &g_entry{Entry: kg.logger.WithError(err), level: kg.level}
}

// Add a context to the log entry.
func (kg *g_log) WithContext(ctx context.Context) *g_entry {
	return &g_entry{Entry: kg.logger.WithContext(ctx), level: kg.level}
}

// Overrides the time of the log entry.
func (kg *g_log) WithTime(t time.Time) *g_entry {
	return &g_entry{Entry: kg.logger.WithTime(t), level: kg.level}
}
