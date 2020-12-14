package utils

type ResponseCode int

const (
	ResponseCodeOk                 ResponseCode = 1000
	ResponseCodeValidationError    ResponseCode = 4000
	ResponseCodeRequestParamsError ResponseCode = 4100
	ResponseCodeInnerServerError   ResponseCode = 5000
	ResponseCodeBizError           ResponseCode = 6000
	ResponseCodeBizTransferFailure ResponseCode = 6010
)

type Response struct {
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}
