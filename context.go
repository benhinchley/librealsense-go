package librealsense

/*
#cgo linux darwin LDFLAGS: -L/usr/local/lib/ -lrealsense -I/usr/local/include/librealsense
#include <librealsense/rs.h>
*/
import "C"

// Context wraps rs_context
type Context C.rs_context

// GetDeviceCount returns the number of connected devices
func (c *Context) GetDeviceCount() (int, error) {
	var err *C.rs_error
	n := int(C.rs_get_device_count((*C.rs_context)(c), &err))
	if err != nil {
		return -1, errorFrom(err)
	}
	return n, nil
}

// GetDevice get the device at idx
func (c *Context) GetDevice(idx int) (*Device, error) {
	var err *C.rs_error
	d := (*Device)(C.rs_get_device((*C.rs_context)(c), C.int(idx), &err))
	if err != nil {
		return nil, errorFrom(err)
	}
	return d, nil
}

// Close the context
func (c *Context) Close() error {
	var err *C.rs_error

	C.rs_delete_context((*C.rs_context)(c), &err)
	if err != nil {
		return errorFrom(err)
	}

	return nil
}
