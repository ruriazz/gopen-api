package models

type IdnDistrict struct {
	Pkid            int16 `gorm:"primaryKey" json:"-"`
	IdnProvincePkid int8  `gorm:"index" json:"-"`
	Slug            string
	Code            string
	DatiName        string
	Name            string
	Alias           string
	IdnProvince     IdnProvince `gorm:"foreignKey:IdnProvincePkid;references:Pkid"`
}

func (IdnDistrict) TableName() string {
	return "idn_district"
}
