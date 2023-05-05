package main

import (
	"log"

	"github.com/TrHung-297/fountain"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

func init() {
}

type NullInstance struct {
	state int
	m     func()
}

func (e *NullInstance) GetIdentification() (addr, dcName, serverName, serverID string) {
	return "", "", "", ""
}

func (e *NullInstance) Initialize() error {
	g_log.V(1).Info("null instance initialize...")
	e.state = 1
	return nil
}

func (e *NullInstance) RunLoop() {
	g_log.V(1).Infof("null run_loop...")
	e.state = 2
	e.m()
}

func (e *NullInstance) Destroy() {
	g_log.V(1).Infof("null destroy...")
	e.state = 3
}

func main() {
	instance := &NullInstance{}
	instance.m = func() {
		fountain.QuitAppInstance()
	}

	fountain.DoMainAppInstance(instance)

	result := instance.state

	log.Printf("result: %d", result)
}
