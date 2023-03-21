package models

type IdnProvince struct {
	Pkid           int8 `gorm:"primaryKey" json:"-"`
	Slug           string
	Code           string
	IsoCode        string
	Name           string
	GeographicArea string
	Capital        string
	Image          string
}

func (IdnProvince) TableName() string {
	return "idn_province"
}
