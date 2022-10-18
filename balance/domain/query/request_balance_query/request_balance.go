package request_balance_query

import (
	"github.com/high-performance-payment-gateway/balance-service/balance/domain/command/request_balance_cmd"
	"github.com/high-performance-payment-gateway/balance-service/balance/infrastructure/cache/redis"
	"github.com/high-performance-payment-gateway/balance-service/balance/infrastructure/db/orm"
	log "github.com/sirupsen/logrus"
)

type (
	OneRequest struct {
		Data DataQuery
	}

	DataQuery = orm.BalanceRequestLog

	DataResponse struct {
		Data    DataQuery
		Status  bool
		Message string
	}

	ParamQuery struct {
		OrderId uint64
	}

	StatusQueryDB struct {
		IsResponse bool
	}

	OneRequestInterface interface {
		GetFromCache(qp ParamQuery) ResponseQueryCache
		HandleOneRequestQuery(qp ParamQuery) DataResponse
	}
)

func (or *OneRequest) GetFromCache(qp ParamQuery) ResponseQueryCache {
	qc := NewQueryCache()
	return qc.QueryRequestBalance(qp)
}

func (or *OneRequest) HandleOneRequestQuery(qp ParamQuery) DataResponse {
	rsQrCache := or.GetFromCache(qp)
	if rsQrCache.Err != nil {
		log.WithFields(log.Fields{
			"errMessage": rsQrCache.Err.Error(),
		}).Error("query one request fr DB error")

		return DataResponse{
			Data:    DataQuery{},
			Status:  false,
			Message: "Please try again late",
		}
	}

	if rsQrCache.IsUseDataForResponse == true {
		return DataResponse{
			Data:    rsQrCache.Data,
			Status:  true,
			Message: "Success",
		}
	}

	if rsQrCache.IsContinueUpdateCacheFrDB == true {
		//todo get fr global value and pass
		cacheUpdate := request_balance_cmd.NewCacheUpdate(ParamQuery{}, redis.RedLockCluster{}, &redis.RedisCluster{})
		rs := cacheUpdate.HandleRequestUpdateCacheFrDB()
		if rs.IsOrderIdExist == true && rs.Err == nil {
			return DataResponse{
				Data:    rs.Data,
				Status:  true,
				Message: "success",
			}
		}

		return DataResponse{
			Data:    DataQuery{},
			Status:  false,
			Message: "have error when update cache",
		}
	}

	if rsQrCache.IsUseDataForResponse == false && rsQrCache.IsContinueUpdateCacheFrDB == false {
		return DataResponse{
			Data:    DataQuery{},
			Status:  false,
			Message: "data is refreshing, please wait and try again late",
		}
	}

	//default error
	return DataResponse{
		Status: false,
	}
}

func NewOneRequest() OneRequestInterface {
	return &OneRequest{}
}
