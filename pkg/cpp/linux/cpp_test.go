package linux

import "testing"

func TestCreateCCppProperties(t *testing.T) {
	c := NewLinux()
	err := c.createCCppProperties()
	if err != nil {
		t.Error("createCCppProperties", err)
	}
}

func TestCreateLaunch(t *testing.T) {
	c := NewLinux()
	err := c.createLaunch()
	if err != nil {
		t.Error("createLaunch", err)
	}
}

func TestCreateSettings(t *testing.T) {
	c := NewLinux()
	err := c.createSettings()
	if err != nil {
		t.Error("createSettings", err)
	}
}

func TestCreateTasks(t *testing.T) {
	c := NewLinux()
	err := c.createTasks()
	if err != nil {
		t.Error("createTasks", err)
	}
}
