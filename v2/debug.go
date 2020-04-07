package iphhy

import (
	"fmt"
	"io"
	"os"
)

var pkgdebug bool
var debugout io.Writer

// SetDebug sets debugging on/off
func SetDebug(dbg bool) {
	pkgdebug = dbg
	debugout = os.Stdout
}

// SetDebugOut sets the target for debugging
func SetDebugOut(out io.Writer) {
	debugout = out
}

// debugf prints debug output
func debugf(format string, a ...interface{}) {
	if pkgdebug && debugout != nil {
		fmt.Fprintf(debugout, format, a...)
	}
}

// debugln printlns the strings
func debugln(a ...interface{}) {
	if pkgdebug && debugout != nil {
		fmt.Fprintln(debugout, a...)
	}
}
