package masterDataDomainEntity

type GetProvinceCollectionParameterV1 struct{ *GetCollectionParameterV1 }

type DefaultIdnProvinceCollectionV1 struct {
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type DefaultIdnProvinceDetailV1 struct {
	Slug           string `json:"slug"`
	Code           string `json:"code"`
	IsoCode        string `json:"isoCode"`
	Name           string `json:"name"`
	GeographicArea string `json:"geographicArea"`
	Capital        string `json:"capital"`
	Image          string `json:"image"`
}
