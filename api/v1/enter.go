package v1

import (
	"myGo/api/v1/apiCustomer"
)

type ApiGroup struct {
	apiCustomer.CustomerApi
}

var ApiGroupApp = new(ApiGroup)
