/*
Copyright © 2021 luckyhappycat Beijing China <luckyhappycat@126.com>

*/
package cpp

import (
	"github.com/pkg/errors"
	"log"
)

func NewCpp() *cpp {
	return &cpp{}
}

type cpp struct {
}

func (c *cpp) Create() error {
	log.Print("pkg/cpp/cpp.create()")
	if err := c.createCCppProperties(); err != nil {
		return errors.WithStack(err)
	}
	if err := c.createLaunch(); err != nil {
		return errors.WithStack(err)
	}
	if err := c.createSettings(); err != nil {
		return errors.WithStack(err)
	}
	if err := c.createTasks(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *cpp) createCCppProperties() error {
	log.Print("pkg/cpp/cpp.createCCppProperties()")
	return nil
}

func (c *cpp) createLaunch() error {
	log.Print("pkg/cpp/cpp.createLaunch()")
	return nil
}

func (c *cpp) createSettings() error {
	log.Print("pkg/cpp/cpp.createSettings()")
	return nil
}

func (c *cpp) createTasks() error {
	log.Print("pkg/cpp/cpp.createTasks()")
	return nil
}
