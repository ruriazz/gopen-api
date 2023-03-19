package masterDataDomainEntity

type GetSubdistrictCollectionParameterV1 struct{ *GetCollectionParameterV1 }
type GetUrbanVillageCollectionBySubdistrictParameterV1 struct{ *GetCollectionParameterV1 }

type DefaultIdnSubdistrictCollectionWithLongNameV1 struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	LongName string `json:"longName"`
}

type DefaultIdnSubdistrictDetailV1 struct {
	Slug        string                     `json:"slug"`
	Code        string                     `json:"code"`
	Name        string                     `json:"name"`
	Alias       string                     `json:"alias"`
	IdnDistrict DefaultIdnDistrictDetailV1 `json:"idnDistrict"`
}

type DefaultIdnUrbanVillageCollectionV1 struct {
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	PostalCode string `json:"postalCode"`
}
