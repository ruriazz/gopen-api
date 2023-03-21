package masterDataRepository

import (
	"database/sql"
	"fmt"
	"strings"

	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/src/helpers/pagination"
	"github.com/ruriazz/gopen-api/src/models"
)

func (r MasterDataRepository) IdnDistrict() domainInterface.IdnDistrictRepositories {
	return IdnDistrictRepository{&r}
}

func (r IdnDistrictRepository) CollectionV1(queries *domainEntity.GetDistrictCollectionParameterV1, withPagination bool) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error) {
	var results []models.IdnDistrict
	offset := (queries.Page - 1) * queries.Limit

	rawQuery := "SELECT d.* FROM idn_district AS d JOIN idn_province AS p ON p.pkid = d.idn_province_pkid ORDER BY p.name ASC, d.name ASC"
	if withPagination {
		rawQuery = "SELECT d.* FROM idn_district AS d JOIN idn_province AS p ON p.pkid = d.idn_province_pkid ORDER BY p.name ASC, d.name ASC LIMIT @limit OFFSET @offset"
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_district WHERE name lIKE @fkeyword ORDER BY LOCATE(@keyword, name) LIMIT @limit OFFSET @offset"
		}
	} else {
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_district WHERE name lIKE @fkeyword ORDER BY LOCATE(@keyword, name)"
		}
	}

	fmt.Println(rawQuery)

	err := r.Databases.MySqlDB.
		Raw(
			rawQuery,
			sql.Named("keyword", queries.Search),
			sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1)),
			sql.Named("offset", offset),
			sql.Named("limit", queries.Limit),
		).
		Preload("IdnProvince").
		Find(&results).
		Error

	if err != nil {
		return nil, nil, err
	}

	if withPagination {
		rawQuery = "SELECT pkid from idn_district WHERE name like @fkeyword"
		session := r.Databases.MySqlDB.
			Raw(rawQuery, sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1))).
			Find(&[]models.IdnDistrict{})
		if session.Error != nil {
			return nil, nil, session.Error
		}

		resPagination, err := paginationHelper.NewPagination(session.RowsAffected, queries.Page, queries.Limit)
		if err != nil {
			return nil, nil, err
		}

		return results, resPagination, nil
	}

	return results, nil, nil
}

func (r IdnDistrictRepository) DetailV1(model models.IdnDistrict) (*models.IdnDistrict, error) {
	err := r.Databases.MySqlDB.
		Where(&model).
		First(&model).
		Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}

		return nil, err
	}

	province := models.IdnProvince{Pkid: model.IdnProvincePkid}
	if err := r.Databases.MySqlDB.First(&province).Error; err != nil {
		return nil, err
	}

	model.IdnProvince = province
	return &model, nil
}

func (r IdnDistrictRepository) SubdistrictCollectionV1(model models.IdnDistrict, queries domainEntity.GetSubdistrictCollectionByDistrictParameterV1, withPagination bool) ([]models.IdnSubdistrict, *paginationHelper.PaginationV1, error) {
	var results []models.IdnSubdistrict
	offset := (queries.Page - 1) * queries.Limit

	rawQuery := "SELECT * FROM idn_subdistrict WHERE idn_district_pkid = @districtPkid ORDER BY name ASC"
	if withPagination {
		rawQuery = "SELECT * FROM idn_subdistrict WHERE idn_district_pkid = @districtPkid ORDER BY name ASC LIMIT @limit OFFSET @offset"
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_subdistrict WHERE idn_district_pkid = @districtPkid AND name LIKE @fkeyword ORDER BY LOCATE(@keyword, name) LIMIT @limit OFFSET @offset"
		}
	} else {
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_subdistrict WHERE idn_district_pkid = @districtPkid AND name LIKE @fkeyword ORDER BY LOCATE(@keyword, name)"
		}
	}

	err := r.Databases.MySqlDB.
		Raw(
			rawQuery,
			sql.Named("districtPkid", model.Pkid),
			sql.Named("keyword", queries.Search),
			sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1)),
			sql.Named("offset", offset),
			sql.Named("limit", queries.Limit),
		).
		Find(&results).
		Error

	if err != nil {
		return nil, nil, err
	}

	if withPagination {
		rawQuery = "SELECT pkid from idn_subdistrict WHERE idn_district_pkid = @districtPkid AND name like @fkeyword"
		session := r.Databases.MySqlDB.
			Raw(rawQuery, sql.Named("districtPkid", model.Pkid), sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1))).
			Find(&[]models.IdnDistrict{})
		if session.Error != nil {
			return nil, nil, session.Error
		}

		resPagination, err := paginationHelper.NewPagination(session.RowsAffected, queries.Page, queries.Limit)
		if err != nil {
			return nil, nil, err
		}

		return results, resPagination, nil
	}

	return results, nil, nil
}
