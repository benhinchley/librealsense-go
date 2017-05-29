package librealsense

/*
#cgo linux darwin LDFLAGS: -L/usr/local/lib/ -lrealsense -I/usr/local/include/librealsense
#include <librealsense/rs.h>
*/
import "C"

// Intrinsics describes the intrinsic camera parameters for a specific stream
type Intrinsics C.rs_intrinsics

func (i *Intrinsics) Width() int {
	return (int)(i.width)
}
func (i *Intrinsics) Height() int {
	return (int)(i.height)
}
func (i *Intrinsics) PPX() float32 {
	return (float32)(i.ppx)
}
func (i *Intrinsics) PPY() float32 {
	return (float32)(i.ppy)
}
func (i *Intrinsics) FX() float32 {
	return (float32)(i.fx)
}
func (i *Intrinsics) FY() float32 {
	return (float32)(i.fy)
}
func (i *Intrinsics) Model() Distortion {
	return (Distortion)(i.model)
}
func (i *Intrinsics) Coeffs() [5]float32 {
	var c [5]float32
	for i, f := range i.coeffs {
		c[i] = (float32)(f)
	}
	return c
}

// Extrinsics describes the topology describing how the different devices are connected.
type Extrinsics C.rs_extrinsics
