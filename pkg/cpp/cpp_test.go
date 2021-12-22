package cpp

import "testing"

func TestCreateLaunch(t *testing.T) {
	c := NewCpp()
	err := c.createLaunch()
	if err != nil {
		t.Error("createLaunch", err)
	}
}
