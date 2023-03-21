package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/settings"
	"github.com/ruriazz/gopen-api/src/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	setting, err := settings.NewSettings()
	if err != nil {
		panic(err)
	}

	database, err := databases.NewDatabases(setting).ConnectMySql()
	if err != nil {
		panic(err)
	}

	if err := loadIdnProvince(database); err != nil {
		panic(err)
	}

	if err := loadIdnDistrict(database); err != nil {
		panic(err)
	}

	if err := loadIdnSubdistrict(database); err != nil {
		panic(err)
	}

	if err := loadIdnUrbanVillage(database); err != nil {
		panic(err)
	}
}

func loadIdnProvince(database *gorm.DB) error {
	type idnProvinceMaster struct {
		ID             int    `json:"id,omitempty"`
		Slug           string `json:"slug,omitempty"`
		Code           string `json:"code,omitempty"`
		ISOCode        string `json:"isoCode,omitempty"`
		Name           string `json:"name,omitempty"`
		GeographicArea string `json:"geographicArea,omitempty"`
		Capital        string `json:"capital,omitempty"`
		Image          string `json:"image,omitempty"`
	}

	fmt.Println("Start IDN Province Data importing...")
	var results []idnProvinceMaster
	client := resty.New()

	res, err := client.R().Get("https://raw.githubusercontent.com/ruriazz/master-data/master/idn-administrative-area/json/provinsi.json")
	if err != nil {
		return err
	}

	json.Unmarshal(res.Body(), &results)
	fmt.Printf("%d IDN Province Data collected\n", len(results))
	for i, result := range results {
		idnProvince := models.IdnProvince{
			Pkid:           int8(result.ID),
			Slug:           result.Slug,
			Code:           result.Code,
			IsoCode:        result.ISOCode,
			Name:           result.Name,
			GeographicArea: result.GeographicArea,
			Capital:        result.Capital,
			Image:          result.Image,
		}

		database.Clauses(clause.OnConflict{DoNothing: true}).Create(&idnProvince)
		fmt.Printf("[%d] %s Created\n", len(results)-(i+1), result.Name)
	}

	fmt.Println("IDN Province Data import finish.")
	return nil
}

func loadIdnDistrict(database *gorm.DB) error {
	type idnDistrictMaster struct {
		ID            int
		IdnProvinsiID int `json:"idnProvinsiId"`
		Slug          string
		Code          string
		DatiName      string `json:"datiName"`
		Name          string
		Alias         string
	}

	fmt.Println("Start IDN District Data importing...")
	var results []idnDistrictMaster
	client := resty.New()

	res, err := client.R().Get("https://raw.githubusercontent.com/ruriazz/master-data/master/idn-administrative-area/json/kabupaten.json")
	if err != nil {
		return err
	}

	json.Unmarshal(res.Body(), &results)
	fmt.Printf("%d IDN District Data collected\n", len(results))
	for i, result := range results {
		idnDistrict := models.IdnDistrict{
			Pkid:            int16(result.ID),
			IdnProvincePkid: int8(result.IdnProvinsiID),
			Slug:            result.Slug,
			Code:            result.Code,
			DatiName:        result.DatiName,
			Name:            result.Name,
			Alias:           result.Alias,
		}

		if err := database.Clauses(clause.OnConflict{DoNothing: true}).Create(&idnDistrict).Error; err != nil {
			fmt.Printf("ERROR create %s ", result.Name)
			fmt.Println(err)
		}
		fmt.Printf("[%d] %s Created\n", len(results)-(i+1), result.Name)
	}

	fmt.Println("IDN District Data import finish.")
	return nil
}

func loadIdnSubdistrict(database *gorm.DB) error {
	type idnSubdistrictMaster struct {
		ID            int
		IdnDistrictId int `json:"idnKabupatenId"`
		Slug          string
		Code          string
		Name          string
		Alias         string
	}

	fmt.Println("Start IDN Subdistrict Data importing...")
	var results []idnSubdistrictMaster
	client := resty.New()

	res, err := client.R().Get("https://raw.githubusercontent.com/ruriazz/master-data/master/idn-administrative-area/json/kecamatan.json")
	if err != nil {
		return err
	}

	json.Unmarshal(res.Body(), &results)
	fmt.Printf("%d IDN Subdistrict Data collected\n", len(results))
	for i, result := range results {
		idnSubdistrict := models.IdnSubdistrict{
			Pkid:            int16(result.ID),
			IdnDistrictPkid: int16(result.IdnDistrictId),
			Slug:            result.Slug,
			Code:            result.Code,
			Name:            result.Name,
			Alias:           result.Alias,
		}

		if err := database.Clauses(clause.OnConflict{DoNothing: true}).Create(&idnSubdistrict).Error; err != nil {
			fmt.Printf("ERROR create %s ", result.Name)
			fmt.Println(err)
		}
		fmt.Printf("[%d] %s Created\n", len(results)-(i+1), result.Name)
	}

	fmt.Println("IDN Subdistrict Data import finish.")
	return nil
}

func loadIdnUrbanVillage(database *gorm.DB) error {
	type idnUrbanVillageMaster struct {
		ID                 int
		IdnSubdistrictPkid int `json:"idnKecamatanId"`
		Slug               string
		Code               string
		Name               string
		Alias              string
		PostalCode         string
	}

	fmt.Println("Start IDN Urban Village Data importing...")
	var results []idnUrbanVillageMaster
	client := resty.New()

	res, err := client.R().Get("https://raw.githubusercontent.com/ruriazz/master-data/master/idn-administrative-area/json/kelurahan.json")
	if err != nil {
		return err
	}

	json.Unmarshal(res.Body(), &results)
	fmt.Printf("%d IDN Urban Village Data collected\n", len(results))
	for i, result := range results {
		idnUrbanVilage := models.IdnUrbanVillage{
			Pkid:               int32(result.ID),
			IdnSubdistrictPkid: int16(result.IdnSubdistrictPkid),
			Slug:               result.Slug,
			Code:               result.Code,
			Name:               result.Name,
			Alias:              result.Alias,
			PostalCode:         result.PostalCode,
		}

		if err := database.Clauses(clause.OnConflict{DoNothing: true}).Create(&idnUrbanVilage).Error; err != nil {
			fmt.Printf("ERROR create %s ", result.Name)
			fmt.Println(err)
		}
		fmt.Printf("[%d] %s Created\n", len(results)-(i+1), result.Name)
	}

	fmt.Println("IDN Urban Village Data import finish.")
	return nil
}
