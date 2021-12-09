package customerService

import (
	"myGo/common/request"
	"myGo/global"
	"myGo/model/customer"
	"time"
)

type CustomerService struct {
}

func (exa *CustomerService) CreateExaCustomer(e customer.Customer) (err error) {
	e.CreatedTime = time.Now()
	e.UpdateTime = time.Now()
	err = global.Db.Create(&e).Error
	return err
}

func (exa *CustomerService) DeleteExaCustomer(e customer.Customer) (err error) {
	err = global.Db.Delete(&e).Error
	return err
}

func (exa *CustomerService) UpdateExaCustomer(e *customer.Customer) (err error) {
	err = global.Db.Save(e).Error
	return err
}

func (exa *CustomerService) GetExaCustomer(id uint) (err error, customer customer.Customer) {
	err = global.Db.Where("id = ?", id).First(&customer).Error
	return
}

func (exa *CustomerService) GetCustomerInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.Db.Model(&customer.Customer{})

	var CustomerList []customer.Customer
	err = db.Count(&total).Error
	if err != nil {
		return err, CustomerList, total
	} else {
		err = db.Limit(limit).Offset(offset).Find(&CustomerList).Error
	}
	return err, CustomerList, total
}
