package application

import (
	"github.com/high-performance-payment-gateway/balance-service/balance/application/query_request_balance"
	"github.com/high-performance-payment-gateway/balance-service/balance/domain/query/request_balance_query"
)

type (
	AllPartnerBalanceQuery struct {
		RequestBalanceQuery request_balance_query.OneRequestInterface
	}

	AllPartnerBalanceQueryInterface interface {
		Init() error
		GetOneRequestBalance(p query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse
	}
)

func (a *AllPartnerBalanceQuery) Init() error {
	return nil
}

func (a *AllPartnerBalanceQuery) GetOneRequestBalance(p query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse {
	return a.RequestBalanceQuery.HandleOneRequestQuery(p)
}

func NewAllPartnerBalance(RequestBalanceQuery request_balance_query.OneRequestInterface) AllPartnerBalanceQueryInterface {
	return &AllPartnerBalanceQuery{
		RequestBalanceQuery: RequestBalanceQuery,
	}
}
