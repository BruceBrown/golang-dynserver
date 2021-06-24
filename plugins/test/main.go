package main

import (
	"fmt"

	"github.com/brucebrown/golang-dynserver/common"
)

// called onetime, when loaded
func init() {
	fmt.Println("init called")
}

type Test struct {
	common.ConnectionFactory
}

func (test Test) Start() bool {
	fmt.Println("test started")
	return true
}
func (test Test) Stop() bool {
	fmt.Println("test stopped")
	return true
}
func (test Test) CreateConnection() common.Connection {
	fmt.Println("test creating connection")
	return TestConnection{}
}

type TestConnection struct {
	common.Connection
}

func (conn TestConnection) Start() bool {
	fmt.Println("test connection started")
	return true
}
func (conn TestConnection) Stop() bool {
	fmt.Println("test connection stopped")
	return true
}
func (conn TestConnection) Run() bool {
	fmt.Println("test connection running")
	return true
}

// return the runable factory
func GetFactory() common.ConnectionFactory {
	fmt.Println("test GetFactory")

	return test
}

var test = Test{}
