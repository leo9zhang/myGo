package apiCustomer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myGo/common/request"
	"myGo/common/response"
	"myGo/global"
	"myGo/model/customer"
	customerRes "myGo/model/customer/response"
	"myGo/utils"
)

type CustomerApi struct {
}

// @Tags Customer
// @Summary 创建客户
// @accept application/json
// @Produce application/json
// @Param data body customer.Customer true "客户用户名, 客户手机号码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customer/customer [post]
func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer customer.Customer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//customerRouter.SysUserID = utils.GetUserID(c)
	//customerRouter.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	if err := customerService.CreateExaCustomer(customer); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Customer
// @Summary 删除客户
// @accept application/json
// @Produce application/json
// @Param data body customer.Customer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/customer [delete]
func (e *CustomerApi) DeleteExaCustomer(c *gin.Context) {
	var customer customer.Customer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := customerService.DeleteExaCustomer(customer); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Customer
// @Summary 更新客户信息
// @accept application/json
// @Produce application/json
// @Param data body customer.Customer true "客户ID, 客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/customer [put]
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
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Customer
// @Summary 获取单一客户信息
// @accept application/json
// @Produce application/json
// @Param data query customer.Customer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
func (e *CustomerApi) GetExaCustomer(c *gin.Context) {
	var customer customer.Customer
	_ = c.ShouldBindQuery(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, data := customerService.GetExaCustomer(customer.ID)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(customerRes.CustomerResponse{Customer: data}, "获取成功", c)
	}
}

// @Tags Customer
// @Summary 分页获取权限客户列表
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customerRouter/customerList [get]
func (e *CustomerApi) GetExaCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, customerList, total := customerService.GetCustomerInfoList(pageInfo)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
