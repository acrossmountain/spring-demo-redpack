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

}

// 账户存储
func (a *Account) Store(ctx SpringWeb.WebContext) {

}

// 查询红包账户
func (a *Account) Envelope(ctx SpringWeb.WebContext) {
}

// 查询账户信息
