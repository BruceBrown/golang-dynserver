package main

import (
	"fmt"

	"github.com/brucebrown/golang-dynserver/common"
)

// called onetime, when loaded
func init() {
	fmt.Println("other init called")
}

type Test struct {
	common.ConnectionFactory
	initialized bool
}

func (test Test) Start() bool {
	fmt.Println("other started")
	return true
}
func (test Test) Stop() bool {
	fmt.Println("other stopped")
	return true
}
func (test Test) CreateConnection() common.Connection {
	fmt.Println("other creating connection")
	return TestConnection{}
}

type TestConnection struct {
	common.Connection
}

func (conn TestConnection) Start() bool {
	fmt.Println("other connection started")
	return true
}
func (conn TestConnection) Stop() bool {
	fmt.Println("other connection stopped")
	return true
}
func (conn TestConnection) Run() bool {
	fmt.Println("other connection running")
	return true
}

// return the runable factory
func GetFactory(config common.ConnectorConfig) common.ConnectionFactory {
	fmt.Println("other GetFactory")
	if !test.initialized {
		test = Test{initialized: true}
	}
	return test
}

var test = Test{}
