package apiCustomer

import "myGo/service"

type ApiGroup struct {
	CustomerApi
}

var customerService = service.ServiceGroupApp.CustomerServiceGroup.CustomerService
