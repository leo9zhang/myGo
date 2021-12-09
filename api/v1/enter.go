package v1

import (
	"myGo/api/v1/apiCustomer"
)

type ApiGroup struct {
	CustomerApiGroup apiCustomer.CustomerApi
}

var ApiGroupApp = new(ApiGroup)
