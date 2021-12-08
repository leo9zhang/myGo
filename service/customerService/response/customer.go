package response

import "myGo/model/customer"

type CustomerResponse struct {
	Customer customer.Customer `json:"customerRouter"`
}
