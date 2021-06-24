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

type ConfigJson struct {
	Plugins []common.ConnectorConfig `json:"plugins"`
	B       []string                 `json:"b"`
	C       []string                 `json:"c"`
}

func main() {
	var data = ConfigJson{}

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

	var factories []common.ConnectionFactory
	var connections []common.Connection
	for _, cfg := range data.Plugins {
		name := cfg.Spec + ".so"
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
		factory := getFactory.(func(common.ConnectorConfig) common.ConnectionFactory)(cfg)
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
