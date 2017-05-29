package librealsense

import "testing"

func TestGetDeviceCount(t *testing.T) {
	c, err := CreateContext()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	n, err := c.GetDeviceCount()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%d connected devices", n)

	c.Close()
}
