package masterDataRepository

import (
	"database/sql"
	"strings"

	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

func (r MasterDataRepository) IdnDistrict() domainInterface.IdnDistrictRepositories {
	return IdnDistrictRepository{&r}
}

func (r IdnDistrictRepository) CollectionV1(model models.IdnDistrict, queries *domainEntity.GetDistrictCollectionParameterV1, withPagination bool) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error) {
	var results []models.IdnDistrict
	offset := (queries.Page - 1) * queries.Limit

	rawQuery := "SELECT * FROM idn_district WHERE idn_province_pkid = @provincePkid ORDER BY name ASC"
	if withPagination {
		rawQuery = "SELECT * FROM idn_district WHERE idn_province_pkid = @provincePkid ORDER BY name ASC LIMIT @limit OFFSET @offset"
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_district WHERE idn_province_pkid = @provincePkid and NAME lIKE @fkeyword ORDER BY LOCATE(@keyword, name) LIMIT @limit OFFSET @offset"
		}
	} else {

	}

	err := r.Databases.MySqlDB.
		Raw(
			rawQuery,
			sql.Named("provincePkid", model.IdnProvince.Pkid),
			sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1)),
			sql.Named("keyword", queries.Search),
			sql.Named("offset", offset),
			sql.Named("limit", queries.Limit),
		).
		Find(&results).Error

	if err != nil {
		return nil, nil, err
	}

	for i, _ := range results {
		results[i].IdnProvince = model.IdnProvince
	}

	if withPagination {
		rawQuery = "SELECT pkid from idn_district WHERE idn_province_pkid = @provincePkid AND name like @fkeyword"
		session := r.Databases.MySqlDB.
			Raw(rawQuery, sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1)), sql.Named("provincePkid", model.IdnProvince.Pkid)).
			Find(&[]models.IdnDistrict{})
		if session.Error != nil {
			return nil, nil, err
		}

		resPagination, err := paginationHelper.NewPagination(session.RowsAffected, queries.Page, queries.Limit)
		if session.Error != nil {
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
