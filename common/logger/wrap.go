package logger

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// Alert alerts v in alert level, and the message is written to error log.
func Alert(v string) {
	logx.Alert(v)
}

// Close closes the logging.
func Close() error {
	return logx.Close()
}

// Disable disables the logging.
func Disable() {
	logx.Disable()
}

// DisableStat disables the stat logs.
func DisableStat() {
	logx.DisableStat()
}

// Error writes v into error log.
func Error(v ...interface{}) {
	logx.Error(v...)
}

// Errorf writes v with format into error log.
func Errorf(format string, v ...interface{}) {
	logx.Errorf(format, v...)
}

// ErrorStack writes v along with call stack into error log.
func ErrorStack(v ...interface{}) {
	// there is newline in stack string
	logx.ErrorStack(v...)
}

// ErrorStackf writes v along with call stack in format into error log.
func ErrorStackf(format string, v ...interface{}) {
	// there is newline in stack string
	logx.ErrorStackf(format, v...)
}

// Errorv writes v into error log with json content.
// No call stack attached, because not elegant to pack the messages.
func Errorv(v interface{}) {
	logx.Errorv(v)
}

// Errorw writes msg along with fields into error log.
func Errorw(msg string, fields ...logx.LogField) {
	logx.Errorw(msg, fields...)
}

// Field returns a LogField for the given key and value.
func Field(key string, value interface{}) logx.LogField {
	return logx.Field(key, value)
}

// Info writes v into access log.
func Info(v ...interface{}) {
	logx.Info(v...)
}

// Infof writes v with format into access log.
func Infof(format string, v ...interface{}) {
	logx.Infof(format, v...)
}

// Infov writes v into access log with json content.
func Infov(v interface{}) {
	logx.Infov(v)
}

// Infow writes msg along with fields into access log.
func Infow(msg string, fields ...logx.LogField) {
	logx.Infow(msg, fields...)
}

// Must checks if err is nil, otherwise logs the error and exits.
func Must(err error) {
	logx.Must(err)
}

// MustSetup sets up logging with given config c. It exits on error.
func MustSetup(c logx.LogConf) {
	logx.Must(SetUp(c))
}

// Reset clears the writer and resets the log level.
func Reset() logx.Writer {
	return logx.Reset()
}

// SetLevel sets the logging level. It can be used to suppress some logs.
func SetLevel(level uint32) {
	logx.SetLevel(level)
}

// SetWriter sets the logging writer. It can be used to customize the logging.
func SetWriter(w logx.Writer) {
	logx.SetWriter(w)
}

// SetUp sets up the logx. If already set up, just return nil.
// we allow SetUp to be called multiple times, because for example
// we need to allow different service frameworks to initialize logx respectively.
func SetUp(c logx.LogConf) (err error) {
	return logx.SetUp(c)
}

// Severe writes v into severe log.
func Severe(v ...interface{}) {
	logx.Severe(v...)
}

// Severef writes v with format into severe log.
func Severef(format string, v ...interface{}) {
	logx.Severef(format, v...)
}

// Slow writes v into slow log.
func Slow(v ...interface{}) {
	logx.Slow(v...)
}

// Slowf writes v with format into slow log.
func Slowf(format string, v ...interface{}) {
	logx.Slowf(format, v...)
}

// Slowv writes v into slow log with json content.
func Slowv(v interface{}) {
	logx.Slowv(v)
}

// Sloww writes msg along with fields into slow log.
func Sloww(msg string, fields ...logx.LogField) {
	logx.Sloww(msg, fields...)
}

// Stat writes v into stat log.
func Stat(v ...interface{}) {
	logx.Stat(v...)
}

// Statf writes v with format into stat log.
func Statf(format string, v ...interface{}) {
	logx.Statf(format, v...)
}

// WithCooldownMillis customizes logging on writing call stack interval.
func WithCooldownMillis(millis int) logx.LogOption {
	return logx.WithCooldownMillis(millis)
}

// WithKeepDays customizes logging to keep logs with days.
func WithKeepDays(days int) logx.LogOption {
	return logx.WithKeepDays(days)
}

// WithGzip customizes logging to automatically gzip the log files.
func WithGzip() logx.LogOption {
	return logx.WithGzip()
}

// WithMaxBackups customizes how many log files backups will be kept.
func WithMaxBackups(count int) logx.LogOption {
	return logx.WithMaxBackups(count)
}

// WithMaxSize customizes how much space the writing log file can take up.
func WithMaxSize(size int) logx.LogOption {
	return logx.WithMaxSize(size)
}

// WithRotation customizes which log rotation rule to use.
func WithRotation(r string) logx.LogOption {
	return logx.WithRotation(r)
}

// WithContext sets ctx to log, for keeping tracing information.
func WithContext(ctx context.Context) logx.Logger {
	return logx.WithContext(ctx)
}

// WithDuration returns a Logger which logs the given duration.
func WithDuration(d time.Duration) logx.Logger {
	return logx.WithDuration(d)
}
