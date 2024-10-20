package types

type ReqConfig struct {
	Limit     int
	Offset    int
	SearchFor string
}

func NewReqConfig(pageSize, page int, searchFor string) *ReqConfig {

	return &ReqConfig{
		Limit:     pageSize,
		Offset:    (page - 1) * pageSize,
		SearchFor: searchFor,
	}

}
