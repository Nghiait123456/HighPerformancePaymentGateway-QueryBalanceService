package request_balance_cmd

type (
	UpdateCache struct {
	}

	UpdateCacheInterface interface {
		createNewMutexLock()
		saveRecordLock()
		getDataFrDB()
		updateCache()
		HandleRequestUpdateCacheFrDB()
	}
)

func (u *UpdateCache) createNewMutexLock() {

}
