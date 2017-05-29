// Package librealsense is a wrapper around the Intel RealSense SDK
// https://github.com/IntelRealSense/librealsense
package librealsense

/*
#cgo linux darwin LDFLAGS: -L/usr/local/lib/ -lrealsense -I/usr/local/include/librealsense
#include <librealsense/rs.h>
*/
import "C"

// CreateContext creates a new realsense context
func CreateContext() (*Context, error) {
	var err *C.rs_error
	c := (*Context)(C.rs_create_context(C.RS_API_VERSION, &err))
	if err != nil {
		return nil, errorFrom(err)
	}
	return c, nil
}
