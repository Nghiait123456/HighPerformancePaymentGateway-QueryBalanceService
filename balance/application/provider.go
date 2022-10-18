package application

import (
	"github.com/google/wire"
	"github.com/high-performance-payment-gateway/balance-service/balance/domain/query/request_balance_query"
)

var ProviderAllPartnerBalanceQuery = wire.NewSet(
	NewAllPartnerBalance,
	wire.Bind(new(request_balance_query.OneRequestInterface), new(*request_balance_query.OneRequest)),
)

var ProviderService = wire.NewSet(
	NewService, ProviderAllPartnerBalanceQuery, wire.Bind(new(ServiceInterface), new(*Service)),
)
