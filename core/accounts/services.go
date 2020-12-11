package accounts

import (
	"errors"

	"github.com/acrossmounation/redpack/services"

	"github.com/go-playground/validator/v10"
	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-logger"
	"github.com/go-spring/spring-web"
	"github.com/shopspring/decimal"
)

func init() {
	SpringBoot.RegisterBean(new(accountService)).
		Export((*services.AccountService)(nil))
}

var _ services.AccountService = &accountService{}

type accountService struct {
	Domain *AccountDomain `autowire:""`
}

func (service *accountService) CreateAccount(dto services.AccountCreatedDTO) (*services.AccountDTO, error) {

	// 验证参数
	if err := SpringWeb.Validate(&dto); err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if !ok {
			SpringLogger.Error("验证错误：", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				// e.Translate()
				SpringLogger.Error(e)
			}
		}
		return nil, err
	}

	// 执行账户创建业务逻辑
	balance, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return nil, err
	}
	account := services.AccountDTO{
		AccountName:  dto.AccountName,
		AccountType:  dto.AccountType,
		CurrencyCode: dto.CurrencyCode,
		UserId:       dto.UserId,
		Username:     dto.Username,
		Balance:      balance,
		Status:       1,
	}
	return service.Domain.Create(account)
}

func (service *accountService) Transfer(dto services.AccountTransferDTO) (services.TransferStatus, error) {

	// 验证参数
	if err := SpringWeb.Validate(&dto); err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if !ok {
			SpringLogger.Error("验证错误：", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				// e.Translate()
				SpringLogger.Error(e)
			}
		}
		return services.TransferStatusFailure, err
	}

	// 执行业务逻辑
	amount, err := decimal.NewFromString(dto.AmountStr)
	if err != nil {
		return services.TransferStatusFailure, err
	}
	dto.Amount = amount
	if dto.ChangeFlag == services.ChangeFlagTransferOut {
		if dto.ChangeType > 0 {
			return services.TransferStatusFailure,
				errors.New("如果 changeFlag 为支出，那么 changeType 必须小于 0")
		}
	} else {
		if dto.ChangeType < 0 {
			return services.TransferStatusFailure,
				errors.New("如果 changeFlag 为收入，那么 changeType 必须大于 0")
		}
	}

	status, err := service.Domain.Transfer(dto)

	if status == services.TransferStatusSuccess {
		backwardDto := dto
		backwardDto.TradeBody = dto.TradeTarget
		backwardDto.TradeTarget = dto.TradeBody
		backwardDto.ChangeType = -dto.ChangeType
		backwardDto.ChangeFlag = -dto.ChangeFlag
		status, err := service.Domain.Transfer(backwardDto)
		return status, err
	}

	return status, err
}

func (service *accountService) StoreValue(dto services.AccountTransferDTO) (services.TransferStatus, error) {

	// 储值情况下，交易主体和交易对象都是本身
	dto.TradeTarget = dto.TradeBody
	dto.ChangeFlag = services.ChangeFlagTransferIn
	dto.ChangeType = services.ChangeTypeAccountStoreValue

	// 验证参数
	if err := SpringWeb.Validate(&dto); err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if !ok {
			SpringLogger.Error("验证错误：", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				// e.Translate()
				SpringLogger.Error(e)
			}
		}
		return services.TransferStatusFailure, err
	}

	// 执行业务逻辑
	amount, err := decimal.NewFromString(dto.AmountStr)
	if err != nil {
		return services.TransferStatusFailure, err
	}
	dto.Amount = amount
	return service.Domain.Transfer(dto)
}

func (service *accountService) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {
	return service.Domain.GetEnvelopeAccountByUserId(userId)
}

func (service *accountService) GetAccountByNo(accountNo string) *services.AccountDTO {
	return service.Domain.GetAccountByNo(accountNo)
}
