package request_balance_query

import "github.com/high-performance-payment-gateway/balance-service/balance/infrastructure/db/orm"

type (
	RecordLock struct {
		Data       orm.BalanceRequestLog
		IsHaveLock bool
	}

	RecordLockInterface interface {
		createRecordLock() error
	}
)
