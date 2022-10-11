package request_balance_cmd

import "github.com/high-performance-payment-gateway/balance-service/balance/domain/query/request_balance_query"

type (
	RecordLock = request_balance_query.DataSavedCache

	RecordLockInterface interface {
		CreateRecordLock() error
	}
)

func CreateRecordLock() RecordLock {
	return RecordLock{
		Data:       request_balance_query.DataQuery{},
		IsHaveLock: true,
	}
}
