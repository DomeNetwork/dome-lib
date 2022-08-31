package log

import (
	"fmt"

	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/common"
)

// D will log a debug entry to standard out.  Nothing
// will be logged in the log level is above debug.
func D(args ...interface{}) {
	if cfg.Str("log.level") != "debug" {
		return
	}

	fmt.Printf("%d [DEBUG]", common.Unix())
	fmt.Println(args...)
}

// E will log a error entry to standard out.
func E(args ...interface{}) {
	fmt.Printf("%d [ERROR]", common.Unix())
	fmt.Println(args...)
}

// I will log a info entry to standard out only if the
// log level is not error.
func I(args ...interface{}) {
	if cfg.Str("log.level") == "error" {
		return
	}

	fmt.Printf("%d [INFO]", common.Unix())
	fmt.Println(args...)
}
