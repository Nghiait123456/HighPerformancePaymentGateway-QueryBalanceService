package application

import (
	"github.com/google/wire"
	"github.com/high-performance-payment-gateway/balance-service/balance/domain/query/request_balance_query"
)

var ProviderAllPartnerBalanceQuery = wire.NewSet(
	NewAllPartnerBalanceQuery,
	request_balance_query.NewOneRequest,
	wire.Bind(new(AllPartnerBalanceQueryInterface), new(*AllPartnerBalanceQuery)),
)

var ProviderService = wire.NewSet(
	NewService, ProviderAllPartnerBalanceQuery, wire.Bind(new(ServiceInterface), new(*Service)),
)
