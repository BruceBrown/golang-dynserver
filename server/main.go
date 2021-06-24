// This is a simple test of dynamically loading a library and calling into it's factory interface
package main

import (
	"fmt"
	"plugin"

	"github.com/brucebrown/golang-dynserver/common"
)

func main() {
	fmt.Println("opening plugin")
	p, err := plugin.Open("test.so")
	if err != nil {
		fmt.Println("plugin open failed")
		panic(err)
	}
	getRunable, err := p.Lookup("GetRunable")
	if err != nil {
		fmt.Println("plugin lookup failed")
		panic(err)
	}

	runner := getRunable.(func() common.Runable)()
	fmt.Println("created=", runner)

	started := runner.Start()
	fmt.Println("started=", started)
	running := runner.Run()
	fmt.Println("running=", running)
	stopped := runner.Stop()
	fmt.Println("stopped=", stopped)

}
