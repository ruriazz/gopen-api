package masterDataDomainInterface

type MasterDataHandlers interface {
	IdnProvince() IdnProvinceHandlers
	IdnDistrict() IdnDistrictHandlers
	IdnSubdistrict() IdnSubdistrictHandlers
	IdnUrbanVillage() IdnUrbanVillageHandlers
}

type MasterDataUsecases interface {
	IdnProvince() IdnProvinceUsecases
	IdnDistrict() IdnDistrictUsecases
	IdnSubdistrict() IdnSubdistrictUsecases
	IdnUrbanVillage() IdnUrbanVillageUsecases
}

type MasterDataRepositories interface {
	IdnProvince() IdnProvinceRepositories
	IdnDistrict() IdnDistrictRepositories
	IdnSubdistrict() IdnSubdistrictRepositories
	IdnUrbanVillage() IdnUrbanVillageRepositories
}

type MasterDataSerializers interface {
	IdnProvince() IdnProvinceSerializers
	IdnDistrict() IdnDistrictSerializers
	IdnSubdistrict() IdnSubdistrictSerializers
	IdnUrbanVillage() IdnUrbanVillageSerializers
}

type MasterDataValidators interface {
	IdnProvince() IdnProvinceValidators
	IdnDistrict() IdnDistrictValidators
	IdnSubdistrict() IdnSubdistrictValidators
	IdnUrbanVillage() IdnUrbanVillageValidators
}
