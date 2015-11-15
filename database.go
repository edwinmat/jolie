package main

import "github.com/edwinmat/server-shared"

type Database interface {
	//FindApplication(id string) (*Application, error)
	GetApplications() ([]*Application, error)
	//UpdateApplication(app *Application, params map[string]interface{}) error
	SaveApplication(app *Application) error
	Handle(queues *shared.ConsumerQueues)

	//FindDevice(id string) (*Device, error)
	//GetDevice() ([]*Device, error)
	//UpdateDevice(device *Device, params map[string]interface{}) error
	//SaveDevice(device *Device) error
}
