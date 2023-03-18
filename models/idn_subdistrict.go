package models

type IdnSubdistrict struct {
	Pkid            int16 `gorm:"primaryKey" json:"-"`
	IdnDistrictPkid int16 `json:"-"`
	Slug            string
	Code            string
	Name            string
	Alias           string
	IdnDistrict     IdnDistrict `gorm:"foreignKey:IdnDistrictPkid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:Pkid;"`
}

func (IdnSubdistrict) TableName() string {
	return "idn_subdistrict"
}
