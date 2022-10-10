package request_balance_query

type (
	DataQueryDB     = DataQuery
	ResponseQueryDB struct {
		Data                 DataQueryDB
		IsUseDataForResponse bool
	}
)
