package masterDataRepository

import (
	"database/sql"
	"strings"

	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

func (r MasterDataRepository) IdnUrbanVillage() domainInterface.IdnUrbanVillageRepositories {
	return IdnUrbanVillageRepository{&r}
}

func (r IdnUrbanVillageRepository) CollectionV1(queries domainEntity.GetUrbanVillageCollectionParameterV1, withPagination bool) ([]models.IdnUrbanVillage, *paginationHelper.PaginationV1, error) {
	var results []models.IdnUrbanVillage
	offset := (queries.Page - 1) * queries.Limit

	rawQuery := `
		SELECT uv.* FROM idn_urban_village AS uv
		JOIN idn_subdistrict AS s ON s.pkid = uv.idn_subdistrict_pkid
		JOIN idn_district AS d ON d.pkid = s.idn_district_pkid
		JOIN idn_province AS p ON p.pkid = d.idn_province_pkid
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
			ORDER BY
				p.name ASC,
				d.name ASC,
				s.name ASC
			LIMIT @limit OFFSET @offset
		`

		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_urban_village WHERE name LIKE @fkeyword ORDER BY LOCATE(@keyword, name) LIMIT @limit OFFSET @offset"
		}
	} else {
		if queries.Search != "" {
			rawQuery = "SELECT * FROM idn_urban_village WHERE name LIKE @fkeyword ORDER BY LOCATE(@keyword, name)"
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
		Preload("IdnSubdistrict").
		Preload("IdnSubdistrict.IdnDistrict").
		Preload("IdnSubdistrict.IdnDistrict.IdnProvince").
		Find(&results).
		Error

	if err != nil {
		return nil, nil, err
	}

	if withPagination {
		rawQuery = "SELECT pkid FROM idn_urban_village WHERE  name LIKE @fkeyword"
		session := r.Databases.MySqlDB.
			Raw(rawQuery, sql.Named("fkeyword", strings.Replace("%?%", "?", queries.Search, 1))).
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

func (r IdnUrbanVillageRepository) DetailV1(model models.IdnUrbanVillage) (*models.IdnUrbanVillage, error) {
	err := r.Databases.MySqlDB.
		Where(&model).
		Preload("IdnSubdistrict").
		Preload("IdnSubdistrict.IdnDistrict").
		Preload("IdnSubdistrict.IdnDistrict.IdnProvince").
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
