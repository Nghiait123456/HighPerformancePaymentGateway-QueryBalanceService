package handle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/high-performance-payment-gateway/balance-service/balance/application"
	"github.com/high-performance-payment-gateway/balance-service/balance/application/query_request_balance"
	"github.com/high-performance-payment-gateway/balance-service/balance/infrastructure/server/web_server"
	validate_api "github.com/high-performance-payment-gateway/balance-service/balance/interfaces/controller/api/validate"
	"github.com/high-performance-payment-gateway/balance-service/balance/interfaces/controller/dto/api/dto_api_request"
	"github.com/high-performance-payment-gateway/balance-service/balance/interfaces/controller/dto/api/dto_api_response"
	validate_base "github.com/high-performance-payment-gateway/balance-service/balance/pkg/external/validate"
)

type (
	RequestBalanceQuery struct {
		sv application.ServiceInterface
	}

	RequestBalanceQueryResponse struct {
		HttpStatus int
		Status     string
		Code       int
		Message    string
	}
)

func (r *RequestBalanceQuery) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(web_server.MapBase{
		"status": "ok",
	})
}

func (r *RequestBalanceQuery) GetOneRequestBalance(c *fiber.Ctx) error {
	rqDto := dto_api_request.NewRequestBalanceQuery()
	res, errB := rqDto.BindDataDto(c)
	if errB != nil {
		return res.Response(c)
	}

	validate := validate_api.ValidateApiRequestBalance{
		VB:  validate_base.NewBaseValidate(),
		Dto: *rqDto,
	}

	validate.Init()
	resV, errV := validate.Validate()
	if errV != nil {
		return resV.Response(c)
	}

	paramSV := query_request_balance.ParamQueryOneBalance{
		OrderId: rqDto.Request.OrderId,
	}

	resRQB := r.sv.GetOneRequestBalance(paramSV)
	resProcess := dto_api_response.NewResponseRequestBalanceDto()
	resProcess.MappingFrServiceRequestBalanceResponse(resRQB)

	return resProcess.Response(c)
}

func NewRequestBalanceQuery(sv application.ServiceInterface) *RequestBalanceQuery {
	return &RequestBalanceQuery{
		sv: sv,
	}
}
