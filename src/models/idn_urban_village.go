package models

type IdnUrbanVillage struct {
	Pkid               int32 `gorm:"primaryKey" json:"-"`
	IdnSubdistrictPkid int16 `gorm:"index" json:"-"`
	Slug               string
	Code               string
	Name               string
	Alias              string
	PostalCode         string
	IdnSubdistrict     IdnSubdistrict `gorm:"foreignKey:IdnSubdistrictPkid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:Pkid;"`
}

func (IdnUrbanVillage) TableName() string {
	return "idn_urban_village"
}
