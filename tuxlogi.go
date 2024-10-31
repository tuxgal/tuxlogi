// Package tuxlogi exposes a generic level logging interface decoupled from any single logging implementation.
package tuxlogi

// Logger represents a generic level logging interface decoupled from any single logging implementation.
type Logger interface {
	// Logs the specified arguments at Trace level using the default format
	// (i.e. separated by space).
	Trace(args ...interface{})
	// Logs the specified arguments at Trace level using the specified format.
	Tracef(format string, args ...interface{})
	// Logs just a new empty line only if the logging level is at least Trace.
	TraceEmpty()
	// Logs the specified arguments at Debug level using the default format
	// (i.e. separated by space).
	Debug(args ...interface{})
	// Logs the specified arguments at Debug level using the specified format.
	Debugf(format string, args ...interface{})
	// Logs just a new empty line only if the logging level is at least Debug.
	DebugEmpty()
	// Logs the specified arguments at Info level using the default format
	// (i.e. separated by space).
	Info(args ...interface{})
	// Logs the specified arguments at Info level using the specified format.
	Infof(format string, args ...interface{})
	// Logs just a new empty line only if the logging level is at least Info.
	InfoEmpty()
	// Logs the specified arguments at Warn level using the default format
	// (i.e. separated by space).
	Warn(args ...interface{})
	// Logs the specified arguments at Warn level using the specified format.
	Warnf(format string, args ...interface{})
	// Logs just a new empty line only if the logging level is at least Warn.
	WarnEmpty()
	// Logs the specified arguments at Error level using the default format
	// (i.e. separated by space).
	Error(args ...interface{})
	// Logs the specified arguments at Error level using the specified format.
	Errorf(format string, args ...interface{})
	// Logs just a new empty line only if the logging level is at least Error.
	ErrorEmpty()
	// Logs the specified arguments at Fatal level using the default format
	// (i.e. separated by space).
	Fatal(args ...interface{})
	// Logs the specified arguments at Fatal level using the specified format.
	Fatalf(format string, args ...interface{})
	// Logs the specified arguments without any levels using the default format
	// (i.e. separated by space). Note that timestamp and log level will not
	// be printed when logging using this function.
	Print(args ...interface{})
	// Logs the specified arguments without any levels using the specified
	// format. Note that timestamp and log level will not be printed when
	// logging using this function.
	Printf(format string, args ...interface{})
}
