package masterDataDomainEntity

type GetDistrictCollectionParameterV1 struct{ *GetCollectionParameterV1 }

type DefaultIdnDistrictCollectionWithLongNameV1 struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	LongName string `json:"longName"`
}

type DefaultIdnDistrictDetailV1 struct {
	Slug        string                     `json:"slug"`
	Code        string                     `json:"code"`
	DatiName    string                     `json:"datiName"`
	Name        string                     `json:"name"`
	Alias       string                     `json:"alias"`
	IdnProvince DefaultIdnProvinceDetailV1 `json:"idnProvince"`
}
