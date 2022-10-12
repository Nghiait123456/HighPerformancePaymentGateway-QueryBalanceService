package request_balance_cmd

import (
	"context"
	redlock "github.com/Nghiait123456/redlock"
	"github.com/high-performance-payment-gateway/balance-service/balance/domain/query/request_balance_query"
	"github.com/high-performance-payment-gateway/balance-service/balance/infrastructure/cache/redis"
	"strconv"
	"time"
)

type (
	UpdateCache struct {
		pr           ParamQuery
		clusterLock  redis.RedLockClusterInterface
		redisCluster redis.RedisClusterInterface
	}

	ParamQuery = request_balance_query.ParamQuery

	UpdateCacheInterface interface {
		createNewMutexLock()
		saveRecordLock()
		refreshLock(r *redlock.Mutex)
		getDataFrDB()
		updateCache()
		HandleRequestUpdateCacheFrDB()
	}
)

func (u *UpdateCache) createNewMutexLock() (*redlock.Mutex, error) {
	k := strconv.FormatUint(u.pr.OrderId, 10)
	customExpiry := redlock.OptionFunc(func(mutex *redlock.Mutex) {
		mutex.SetExpiry(8 * time.Second)
	})

	customTries := redlock.OptionFunc(func(mutex *redlock.Mutex) {
		mutex.SetTries(3)
	})

	customDelayFc := redlock.OptionFunc(func(mutex *redlock.Mutex) {
		mutex.SetDelayFunc(func(tries int) time.Duration {
			return 200 * time.Microsecond
		})
	})

	mutex := u.clusterLock.NewMutex(k, customExpiry, customTries, customDelayFc)
	if err := mutex.Lock(); err != nil {
		return &redlock.Mutex{}, err
	}

	return mutex, nil
}

func (u *UpdateCache) refreshLock(r *redlock.Mutex) (bool, error) {
	return r.Unlock()
}

func (u *UpdateCache) saveRecordLock() error {
	ctx := context.Background()
	k := strconv.FormatUint(u.pr.OrderId, 10)
	recordLock := CreateRecordLock()
	exp := time.Second * 9000
	return u.redisCluster.SetKeyWithStructValue(ctx, k, recordLock, exp)
}

func (u UpdateCache) getDataFrDB() {
	//todo get data from DB service
}

func (u *UpdateCache) updateCache() {
	//todo get data from DB service
}

func (u *UpdateCache) HandleRequestUpdateCacheFrDB() {
	//todo handle
}
