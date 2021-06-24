// This is a simple test of dynamically loading a library and calling into it's factory interface
package main

import (
	"encoding/json"
	"fmt"
	"github.com/brucebrown/golang-dynserver/common"
	"io/ioutil"
	"plugin"
)

type MyType map[string][]string

type TestJson struct {
	plugins []string
}

func main() {
	var data MyType

	file, err := ioutil.ReadFile("dynserver.json")
	if err != nil {
		fmt.Println("failed to read dynserver.json")
		panic(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("failed to unmarshal dynserver.json")
		panic(err)
	}
	names := data["plugins"]
	var factories []common.ConnectionFactory
	var connections []common.Connection
	for _, v := range names {
		name := v + ".so"
		p, err := plugin.Open(name)
		if err != nil {
			fmt.Println("plugin open failed", name)
			panic(err)
		}
		getFactory, err := p.Lookup("GetFactory")
		if err != nil {
			fmt.Println("plugin lookup failed")
			panic(err)
		}

		factory := getFactory.(func() common.ConnectionFactory)()
		factories = append(factories, factory)

		started := factory.Start()
		fmt.Println("started=", started)

		connection := factory.CreateConnection()
		if connection == nil {
			fmt.Println("Create Connect failed")
		}
		connections = append(connections, connection)
		connStarted := connection.Start()
		fmt.Println("conn started=", connStarted)
		connRunning := connection.Run()
		fmt.Println("conn running=", connRunning)
	}

	// shutdown
	for _, connection := range connections {
		stopped := connection.Stop()
		fmt.Println("conn stopped=", stopped)
	}
	for _, factory := range factories {
		stopped := factory.Stop()
		fmt.Println("factory stopped=", stopped)
	}
}
