package masterDataDomainEntity

type GetUrbanVillageCollectionParameterV1 struct{ *GetCollectionParameterV1 }

type DefaultIdnUrbanVillageCollectionWithLongNameV1 struct {
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	LongName   string `json:"longName"`
	PostalCode string `json:"postalCode"`
}

type DefaultIdnUrbanVillageDetailV1 struct {
	Slug           string                        `json:"slug"`
	Code           string                        `json:"code"`
	Name           string                        `json:"name"`
	Alias          string                        `json:"alias"`
	PostalCode     string                        `json:"postalCode"`
	IdnSubdistrict DefaultIdnSubdistrictDetailV1 `json:"idnSubdistrict"`
}
