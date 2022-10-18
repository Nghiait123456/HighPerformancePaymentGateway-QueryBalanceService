package application

import "github.com/high-performance-payment-gateway/balance-service/balance/application/query_request_balance"

type (
	ServiceInterface interface {
		Init()
		GetOneRequestBalance(pq query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse
	}
)
