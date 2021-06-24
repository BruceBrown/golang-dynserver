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
	fmt.Println("started")
	return true
}
func (test Test) Stop() bool {
	fmt.Println("stopped")
	return true
}
func (test Test) CreateConnection() Connection {
	fmt.Println("creating connection")
	return true
}

// return the runable factory
func GetFactory() common.ConnectionFactory {
	fmt.Println("GetFactory")

	return test
}

var test = Test{}
