package types

type ReqConfig struct {
	Limit  int
	Offset int
}

func NewReqConfig(pageSize, page int) *ReqConfig {

	return &ReqConfig{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	}

}
