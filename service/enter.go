package service

import "myGo/service/customerService"

type ServiceGroup struct {
	CustomerServiceGroup customerService.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
