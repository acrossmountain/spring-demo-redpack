package apis

import (
	"net/http"

	"github.com/acrossmounation/redpack/services"
	"github.com/acrossmounation/redpack/utils"

	"github.com/go-spring/spring-web"
)

type Account struct {
	AccountService services.AccountService `autowire:""`
}

// 账户创建
func (a *Account) Create(ctx SpringWeb.WebContext) {
	dto := services.AccountCreatedDTO{}
	err := ctx.Bind(&dto)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeRequestParamsError,
			Message: err.Error(),
		})
		return
	}
	account, err := a.AccountService.CreateAccount(dto)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeInnerServerError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Code:    utils.ResponseCodeOk,
		Message: "创建账户成功",
		Data:    account,
	})
	return
}

// 账户转账
func (a *Account) Transfer(ctx SpringWeb.WebContext) {
	transfer := services.AccountTransferDTO{}
	err := ctx.Bind(&transfer)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeRequestParamsError,
			Message: err.Error(),
		})
		return
	}
	status, err := a.AccountService.Transfer(transfer)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeInnerServerError,
			Message: err.Error(),
		})
		return
	}

	if status == services.TransferStatusSuccess {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeOk,
			Message: services.TransferStatusText[services.TransferStatusSuccess],
			Data:    status,
		})
	} else {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeBizTransferFailure,
			Message: services.TransferStatusText[status] + err.Error(),
		})
	}
	return
}

// 查询红包账户
func (a *Account) Envelope(ctx SpringWeb.WebContext) {
	userId := ctx.PathParam("userId")
	if len(userId) == 0 {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeRequestParamsError,
			Message: "请求参数错误",
		})
		return
	}

	account := a.AccountService.GetEnvelopeAccountByUserId(userId)
	if account == nil {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeInnerServerError,
			Message: "未查找到账户",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Code: utils.ResponseCodeOk,
		Data: account,
	})
	return
}

// 查询账户信息
func (a *Account) Account(ctx SpringWeb.WebContext) {
	accountNo := ctx.PathParam("accountNo")

	if len(accountNo) == 0 {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeRequestParamsError,
			Message: "请求参数错误",
		})
		return
	}

	account := a.AccountService.GetAccountByNo(accountNo)
	if account == nil {
		ctx.JSON(http.StatusOK, utils.Response{
			Code:    utils.ResponseCodeInnerServerError,
			Message: "未查找到账户",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Code: utils.ResponseCodeOk,
		Data: account,
	})
	return
}
