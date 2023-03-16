package masterDataDomainEntity

type GetCollectionParameterV1 struct {
	Search string `form:"search"`
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
}
