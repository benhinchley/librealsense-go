package librealsense

/*
#cgo linux darwin LDFLAGS: -L/usr/local/lib/ -lrealsense -I/usr/local/include/librealsense
#include <librealsense/rs.h>
*/
import "C"

import "fmt"

// Error represents an error from librealsense
type Error struct {
	msg string
	// funcName string
	// args     string
}

func (e *Error) Error() string {
	return fmt.Sprintf("librealsense: %s", e.msg)
}

// errorFrom creates an Error from an rs_error
func errorFrom(err *C.rs_error) *Error {
	return &Error{
		msg: C.GoString(C.rs_get_error_message(err)),
		// funcName: C.GoString(C.rs_get_failed_function(err)),
		// args:     C.GoString(C.rs_get_failed_args(err)),
	}
}
