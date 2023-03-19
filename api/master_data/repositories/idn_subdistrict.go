package masterDataRepository

import (
	"database/sql"
	"strings"

	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

func (r MasterDataRepository) IdnSubdistrict() domainInterface.IdnSubdistrictRepositories {
	return IdnSubdistrictRepository{&r}
}

func (r IdnSubdistrictRepository) CollectionV1(queries domainEntity.GetSubdistrictCollectionParameterV1, withPagination bool) ([]models.IdnSubdistrict, *paginationHelper.PaginationV1, error) {
	var results []models.IdnSubdistrict
	offset := (queries.Page - 1) * queries.Limit

	rawQuery := "SELECT s.* FROM idn_subdistrict AS s JOIN idn_district AS d ON d.pkid = s.idn_district_pkid JOIN idn_province as p ON p.pkid = d.idn_province_pkid ORDER BY p.name ASC, d.name ASC, s.name ASC"
	if withPagination {
		rawQuery = "SELECT s.* FROM idn_subdistrict AS s JOIN idn_district AS d ON d.pkid = s.idn_district_pkid JOIN idn_province as p ON p.pkid = d.idn_province_pkid ORDER BY p.name ASC, d.name ASC, s.name ASC LIMIT @limit OFFSET @offset"
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_subdistrict WHERE name LIKE @fkeyword ORDER BY LOCATE(@keyword, name) LIMIT @limit OFFSET @offset"
		}
	} else {
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_subdistrict WHERE name LIKE @fkeyword ORDER BY LOCATE(@keyword, name)"
		}
	}

	err := r.Databases.MySqlDB.
		Raw(
			rawQuery,
			sql.Named("keyword", queries.Search),
			sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1)),
			sql.Named("offset", offset),
			sql.Named("limit", queries.Limit),
		).
		Preload("IdnDistrict").
		Preload("IdnDistrict.IdnProvince").
		Find(&results).
		Error

	if err != nil {
		return nil, nil, err
	}

	if withPagination {
		rawQuery = "SELECT pkid FROM idn_subdistrict WHERE name LIKE @fkeyword"
		session := r.Databases.MySqlDB.
			Raw(rawQuery, sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1))).
			Find(&[]models.IdnSubdistrict{})

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

func (r IdnSubdistrictRepository) DetailV1(model models.IdnSubdistrict) (*models.IdnSubdistrict, error) {
	err := r.Databases.MySqlDB.
		Where(&model).
		Preload("IdnDistrict").
		Preload("IdnDistrict.IdnProvince").
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

func (r IdnSubdistrictRepository) UrbanVillageCollectionV1(model models.IdnSubdistrict, queries domainEntity.GetUrbanVillageCollectionBySubdistrictParameterV1, withPagination bool) ([]models.IdnUrbanVillage, *paginationHelper.PaginationV1, error) {
	var results []models.IdnUrbanVillage
	offset := (queries.Page - 1) * queries.Limit

	rawQuery := `
		SELECT uv.* FROM idn_urban_village AS uv
		JOIN idn_subdistrict AS s ON s.pkid = uv.idn_subdistrict_pkid
		JOIN idn_district AS d ON d.pkid = s.idn_district_pkid
		JOIN idn_province AS p ON p.pkid = d.idn_province_pkid
		WHERE uv.idn_subdistrict_pkid = @subdistrictPkid
		ORDER BY
			p.name ASC,
			d.name ASC,
			s.name ASC
	`
	if withPagination {
		rawQuery = `
			SELECT uv.* FROM idn_urban_village AS uv
			JOIN idn_subdistrict AS s ON s.pkid = uv.idn_subdistrict_pkid
			JOIN idn_district AS d ON d.pkid = s.idn_district_pkid
			JOIN idn_province AS p ON p.pkid = d.idn_province_pkid
			WHERE uv.idn_subdistrict_pkid = @subdistrictPkid
			ORDER BY
				p.name ASC,
				d.name ASC,
				s.name ASC
			LIMIT @limit OFFSET @offset
		`

		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_urban_village WHERE idn_subdistrict_pkid = @subdistrictPkid AND name LIKE @fkeyword ORDER BY LOCATE(@keyword, name) LIMIT @limit OFFSET @offset"
		}
	} else {
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_urban_village WHERE idn_subdistrict_pkid = @subdistrictPkid AND name LIKE @fkeyword ORDER BY LOCATE(@keyword, name)"
		}
	}

	err := r.Databases.MySqlDB.
		Raw(
			rawQuery,
			sql.Named("subdistrictPkid", model.Pkid),
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
		rawQuery = "SELECT pkid FROM idn_urban_village WHERE idn_subdistrict_pkid = @subdistrictPkid AND name LIKE @fkeyword"
		session := r.Databases.MySqlDB.
			Raw(rawQuery, sql.Named("subdistrictPkid", model.Pkid), sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1))).
			Find(&[]models.IdnUrbanVillage{})

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
