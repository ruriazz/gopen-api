package masterDataRepository

import (
	"database/sql"
	"strings"

	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

func (r MasterDataRepository) IdnProvince() domainInterface.IdnProvinceRepositories {
	return IdnProvinceRepository{&r}
}

func (r IdnProvinceRepository) CollectionV1(queries *domainEntity.GetProvinceCollectionParameterV1, pagination bool) ([]models.IdnProvince, *paginationHelper.PaginationV1, error) {
	var results []models.IdnProvince

	offset := (queries.Page - 1) * queries.Limit

	rawQuery := "SELECT * FROM idn_province ORDER BY geographic_area ASC, name ASC LIMIT"
	if pagination {
		rawQuery = "SELECT * FROM idn_province ORDER BY geographic_area ASC, name ASC LIMIT @limit OFFSET @offset"
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_province WHERE name LIKE @fkeyword ORDER BY LOCATE(@keyword, name) LIMIT @limit OFFSET @offset"
		}
	} else {
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_province WHERE name LIKE @fkeyword ORDER BY LOCATE(@keyword, name)"
		}
	}

	err := r.Databases.MySqlDB.
		Raw(rawQuery,
			sql.Named("keyword", queries.Search),
			sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1)),
			sql.Named("limit", queries.Limit),
			sql.Named("offset", offset),
		).
		Find(&results).Error

	if err != nil {
		return nil, nil, err
	}

	if pagination {
		rawQuery = "SELECT pkid from idn_province WHERE name like @fkeyword"
		session := r.Databases.MySqlDB.
			Raw(rawQuery, sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1))).
			Find(&[]models.IdnProvince{})
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

func (r IdnProvinceRepository) DistrictCollectionV1(model models.IdnProvince, queries *domainEntity.GetDistrictCollectionByProvinceParameterV1, withPagination bool) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error) {
	var results []models.IdnDistrict
	offset := (queries.Page - 1) * queries.Limit

	rawQuery := "SELECT * FROM idn_district WHERE idn_province_pkid = @provincePkid ORDER BY name ASC"
	if withPagination {
		rawQuery = "SELECT * FROM idn_district WHERE idn_province_pkid = @provincePkid ORDER BY name ASC LIMIT @limit OFFSET @offset"
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_district WHERE idn_province_pkid = @provincePkid and name lIKE @fkeyword ORDER BY LOCATE(@keyword, name) LIMIT @limit OFFSET @offset"
		}
	} else {
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_district WHERE idn_province_pkid = @provincePkid AND name lIKE @fkeyword ORDER BY LOCATE(@keyword, name)"
		}
	}

	err := r.Databases.MySqlDB.
		Raw(
			rawQuery,
			sql.Named("provincePkid", model.Pkid),
			sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1)),
			sql.Named("keyword", queries.Search),
			sql.Named("offset", offset),
			sql.Named("limit", queries.Limit),
		).
		Find(&results).Error

	if err != nil {
		return nil, nil, err
	}

	if withPagination {
		rawQuery = "SELECT pkid from idn_district WHERE idn_province_pkid = @provincePkid AND name like @fkeyword"
		session := r.Databases.MySqlDB.
			Raw(rawQuery, sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1)), sql.Named("provincePkid", model.Pkid)).
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

func (r IdnProvinceRepository) DetailV1(model models.IdnProvince) (*models.IdnProvince, error) {
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

	return &model, nil
}
