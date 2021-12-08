package apiCustomer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myGo/common/request"
	"myGo/common/response"
	"myGo/global"
	"myGo/model/customer"
	"myGo/utils"
)

type CustomerApi struct {
}

func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer customer.Customer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//customerRouter.SysUserID = utils.GetUserID(c)
	//customerRouter.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	//if err := customerService.CreateExaCustomer(customerRouter); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

func (e *CustomerApi) DeleteExaCustomer(c *gin.Context) {
	var customer customer.Customer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := customerService.DeleteExaCustomer(customer); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (e *CustomerApi) UpdateExaCustomer(c *gin.Context) {
	var customer customer.Customer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := customerService.UpdateExaCustomer(&customer); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (e *CustomerApi) GetExaCustomer(c *gin.Context) {
	var customer customer.Customer
	_ = c.ShouldBindQuery(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//err, data := customerService.GetExaCustomer(customerRouter.ID)
	//if err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(consumer.CustomerResponse{Customer: data}, "获取成功", c)
	//}
}

func (e *CustomerApi) GetExaCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//err, customerList, total := customerService.GetCustomerInfoList(utils.GetUserAuthorityId(c), pageInfo)
	//if err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败"+err.Error(), c)
	//} else {
	//	response.OkWithDetailed(response.PageResult{
	//		List:     customerList,
	//		Total:    total,
	//		Page:     pageInfo.Page,
	//		PageSize: pageInfo.PageSize,
	//	}, "获取成功", c)
	//}
}
