package librealsense

import "testing"

func TestCreateContext(t *testing.T) {
	c, err := CreateContext()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	c.Close()
}

func TestGetDevice(t *testing.T) {
	c, err := CreateContext()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if n, _ := c.GetDeviceCount(); n == 0 {
		t.Error("no device connected")
		t.FailNow()
	}

	d, err := c.GetDevice(0)
	if err != nil {
		t.Error(err)
	}
	n, _ := d.Name()
	t.Logf("device name: %s", n)

	c.Close()
}
