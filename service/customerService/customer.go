package customerService

import (
	"myGo/common/request"
	"myGo/global"
	"myGo/model/customer"
)

type CustomerService struct {
}

func (exa *CustomerService) CreateExaCustomer(e customer.Customer) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *CustomerService) DeleteExaCustomer(e customer.Customer) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (exa *CustomerService) UpdateExaCustomer(e *customer.Customer) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *CustomerService) GetExaCustomer(id uint) (err error, customer customer.Customer) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

func (exa *CustomerService) GetCustomerInfoList(dataId string, info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&customer.Customer{})

	var CustomerList []customer.Customer
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return err, CustomerList, total
	} else {
		err = db.Limit(limit).Offset(offset).Find(&CustomerList).Error
	}
	return err, CustomerList, total
}
