package service

import "myGo/service/customerService"

type ServiceGroup struct {
	CustomerServiceGroup customerService.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
