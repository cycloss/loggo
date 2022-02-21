package loggo

import (
	"log"
	"os"
)

const (
	log_emerg  = "<0>" // system is unusable
	log_alert  = "<1>" // system must have action taken immediately
	log_crit   = "<2>" // software is in a critical condition
	log_err    = "<3>" // software experienced an error
	log_warn   = "<4>" // software experienced a potential problem
	log_notice = "<5>" // software experienced a noteworthy event
	log_info   = "<6>" // software running information
	log_debug  = "<7>" // software debug information
)

var (
	emergencyLogger *log.Logger
	alertLogger     *log.Logger
	criticalLogger  *log.Logger
	errorLogger     *log.Logger
	warningLogger   *log.Logger
	noticeLogger    *log.Logger
	infoLogger      *log.Logger
	debugLogger     *log.Logger
)

const (
	errorFlags    = log.Ldate | log.Ltime | log.Lshortfile
	standardFlags = log.LstdFlags
)

// Must be called before using any of the logging functions.
func Init() {
	emergencyLogger = log.New(os.Stderr, log_emerg, errorFlags)
	alertLogger = log.New(os.Stderr, log_alert, errorFlags)
	criticalLogger = log.New(os.Stderr, log_crit, errorFlags)
	errorLogger = log.New(os.Stderr, log_err, errorFlags)
	warningLogger = log.New(os.Stderr, log_warn, errorFlags)
	noticeLogger = log.New(os.Stderr, log_notice, errorFlags)
	infoLogger = log.New(os.Stderr, log_info, errorFlags)
	debugLogger = log.New(os.Stderr, log_debug, errorFlags)
}

// Logs and then closes the program.
// Level 2 priority in `journald`.
// Appears in red in `journalctl`.
// Included for backward compatibility.
func Fatal(v ...interface{}) {
	log.SetPrefix(log_crit)
	log.Fatal(v...)
}

// Logs and then closes the program.
// Level 2 priority in `journald`.
// Appears in red in `journalctl`.
// Included for backward compatibility.
func Fatalf(format string, v ...interface{}) {
	log.SetPrefix(log_crit)
	log.Fatalf(format, v...)
}

// Indicates that the system is unusable.
// Level 0 priority in `journald`.
// A broadcast message will be sent to all users with a terminal attached to the machine.
// Appears in red in `journalctl`.
func Emergency(v ...interface{}) {
	emergencyLogger.Print(v...)
}

// Indicates that the system is unusable.
// Level 0 priority in `journald`.
// A broadcast message will be sent to all users with a terminal attached to the machine.
// Appears in red in `journalctl`.
func Emergencyf(format string, v ...interface{}) {
	emergencyLogger.Printf(format, v...)
}

// Indicates that action must be taken on the system immediately.
// Level 1 priority in `journald`.
// Appears in red in `journalctl`.
func Alert(v ...interface{}) {
	alertLogger.Print(v...)
}

// Indicates that action must be taken on the system immediately.
// Level 1 priority in `journald`.
// Appears in red in `journalctl`.
func Alertf(format string, v ...interface{}) {
	alertLogger.Printf(format, v...)
}

// Indicates that software is in a critical condition.
// Level 2 priority in `journald`.
// Appears in red in `journalctl`.
func Critical(v ...interface{}) {
	criticalLogger.Print(v...)
}

// Indicates that software is in a critical condition.
// Level 2 priority in `journald`.
// Appears in red in `journalctl`.
func Criticalf(format string, v ...interface{}) {
	criticalLogger.Printf(format, v...)
}

// Indicates that software experienced an error.
// Level 3 priority in `journald`.
// Appears in red in `journalctl`.
func Error(v ...interface{}) {
	errorLogger.Print(v...)
}

// Indicates that software experienced an error.
// Level 3 priority in `journald`.
// Appears in red in `journalctl`.
func Errorf(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}

// Indicates that software experienced a potential problem.
// Level 4 priority in `journald`.
// Appears in yellow in `journalctl`.
func Warn(v ...interface{}) {
	warningLogger.Print(v...)
}

// Indicates that software experienced a potential problem.
// Level 4 priority in `journald`.
// Appears in yellow in `journalctl`.
func Warnf(format string, v ...interface{}) {
	warningLogger.Printf(format, v...)
}

// Indicates that software experienced a noteworthy event.
// Level 5 priority in `journald`.
// Appears in bold white in `journalctl`.
func Notice(v ...interface{}) {
	noticeLogger.Print(v...)
}

// Indicates that software experienced a noteworthy event.
// Level 5 priority in `journald`.
// Appears in bold white in `journalctl`.
func Noticef(format string, v ...interface{}) {
	noticeLogger.Printf(format, v...)
}

// Indicates software running information.
// Level 6 priority in `journald`.
// Appears in white in `journalctl`.
func Info(v ...interface{}) {
	infoLogger.Print(v...)
}

// Indicates software running information.
// Level 6 priority in `journald`.
// Appears in white in `journalctl`.
func Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

// Indicates software debugging information.
// Level 7 priority in `journald`.
// Appears in grey in `journalctl`.
func Debug(v ...interface{}) {
	debugLogger.Print(v...)
}

// Indicates software debugging information.
// Level 7 priority in `journald`.
// Appears in grey in `journalctl`.
func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

// Logs as level 7 priority in `journald`.
// Appears in grey in `journalctl`.
// Included for backward compatibility.
func Print(v ...interface{}) {
	debugLogger.Print(v...)
}

// Logs as level 7 priority in `journald`.
// Appears in grey in `journalctl`.
// Included for backward compatibility.
func Printf(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}
