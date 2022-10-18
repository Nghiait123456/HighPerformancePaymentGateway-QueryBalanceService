package request_balance_cmd

type (
	RecordLock = DataSavedCache
)

func CreateRecordLock() RecordLock {
	return RecordLock{
		Data:       DataQuery{},
		IsHaveLock: true,
	}
}
