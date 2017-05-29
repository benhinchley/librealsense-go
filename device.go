package librealsense

/*
#cgo linux darwin LDFLAGS: -L/usr/local/lib/ -lrealsense -I/usr/local/include/librealsense
#include <librealsense/rs.h>
*/
import "C"
import (
	"context"
	"unsafe"
)

// Device wraps rs_device
type Device C.rs_device

// Name retrieves human-readable device model string
func (d *Device) Name() (string, error) {
	var err *C.rs_error
	n := C.GoString(C.rs_get_device_name((*C.rs_device)(d), &err))
	if err != nil {
		return "", errorFrom(err)
	}
	return n, nil
}

// Serial retrieves unique serial number of the device
func (d *Device) Serial() (string, error) {
	var err *C.rs_error
	n := C.GoString(C.rs_get_device_serial((*C.rs_device)(d), &err))
	if err != nil {
		return "", errorFrom(err)
	}
	return n, nil
}

// FirmwareVersion retrieves the version of the firmware currently installed on the device
func (d *Device) FirmwareVersion() (string, error) {
	var err *C.rs_error
	n := C.GoString(C.rs_get_device_firmware_version((*C.rs_device)(d), &err))
	if err != nil {
		return "", errorFrom(err)
	}
	return n, nil
}

// Start begins streaming on all enabled streams for this device
func (d *Device) Start() error {
	var err *C.rs_error
	C.rs_start_device((*C.rs_device)(d), &err)
	if err != nil {
		return errorFrom(err)
	}
	return nil
}

// Stop ends data acquisition for the specified source providers
func (d *Device) Stop() error {
	var err *C.rs_error
	C.rs_stop_device((*C.rs_device)(d), &err)
	if err != nil {
		return errorFrom(err)
	}
	return nil
}

// EnableStreamPreset enables a specific stream and requests properties using a preset
func (d *Device) EnableStreamPreset(s Stream, p Preset) error {
	var err *C.rs_error
	C.rs_enable_stream_preset((*C.rs_device)(d), (C.rs_stream)(s), (C.rs_preset)(p), &err)
	if err != nil {
		return errorFrom(err)
	}
	return nil
}

// EnableStream enables a specific stream and requests specific properties
func (d *Device) EnableStream(s Stream, w int, h int, f Format, fr int) error {
	var err *C.rs_error
	C.rs_enable_stream((*C.rs_device)(d), (C.rs_stream)(s), C.int(w), C.int(h), (C.rs_format)(f), C.int(fr), &err)
	if err != nil {
		return errorFrom(err)
	}
	return nil
}

// GetDepthScale retrieves mapping between the units of the depth image and meters
func (d *Device) GetDepthScale() (float32, error) {
	var err *C.rs_error
	s := (float32)(C.rs_get_device_depth_scale((*C.rs_device)(d), &err))
	if err != nil {
		return 0.0, errorFrom(err)
	}
	return s, nil
}

// GetStreamWidth retrieves the width in pixels of a specific stream, equivalent to the width field from the stream's intrinsic
func (d *Device) GetStreamWidth(s Stream) (int, error) {
	var err *C.rs_error
	w := (int)(C.rs_get_stream_width((*C.rs_device)(d), (C.rs_stream)(s), &err))
	if err != nil {
		return 0, errorFrom(err)
	}
	return w, nil
}

// GetStreamHeight retrieves the height in pixels of a specific stream, equivalent to the height field from the stream's intrinsic
func (d *Device) GetStreamHeight(s Stream) (int, error) {
	var err *C.rs_error
	h := (int)(C.rs_get_stream_height((*C.rs_device)(d), (C.rs_stream)(s), &err))
	if err != nil {
		return 0, errorFrom(err)
	}
	return h, nil
}

//GetStreamIntrinsics retrieves intrinsic camera parameters for a specific stream
func (d *Device) GetStreamIntrinsics(s Stream) (*Intrinsics, error) {
	var err *C.rs_error
	var i Intrinsics
	C.rs_get_stream_intrinsics((*C.rs_device)(d), (C.rs_stream)(s), (*C.rs_intrinsics)(&i), &err)
	if err != nil {
		return nil, errorFrom(err)
	}
	return &i, nil
}

//GetExtrinsics retrieves extrinsic transformation between the viewpoints of two different streams
func (d *Device) GetExtrinsics(from, to Stream) (*Extrinsics, error) {
	var err *C.rs_error
	var e Extrinsics
	C.rs_get_device_extrinsics((*C.rs_device)(d), (C.rs_stream)(from), (C.rs_stream)(to), (*C.rs_extrinsics)(&e), &err)
	if err != nil {
		return nil, errorFrom(err)
	}
	return &e, nil
}

// PollForFrames provides the abiility to call a function when new frame
// data is avaliable
func (d *Device) PollForFrames(ctx context.Context, fn func()) {
	// NOTE(@benhinchley): We now run outside of the main event loop,
	// but are ignoring any potential errors rasied from the
	// C.rs_poll_for_frames function call
	go pollForFramesWorker(ctx, (*C.rs_device)(d), fn)
}

func pollForFramesWorker(ctx context.Context, dev *C.rs_device, fn func()) {
	var err *C.rs_error
	for {
		select {
		case <-ctx.Done():
			return
		default:
			switch C.rs_poll_for_frames(dev, &err) {
			case 0:
				// there is no frame
				continue
			case 1:
				// there is a frame
				fn()
			}
		}
	}
}

// GetFrameData retrieves the contents of the latest frame on a stream
func (d *Device) GetFrameData(s Stream) (unsafe.Pointer, error) {
	var err *C.rs_error
	df := C.rs_get_frame_data((*C.rs_device)(d), (C.rs_stream)(s), &err)
	if err != nil {
		return nil, errorFrom(err)
	}
	return df, nil
}
