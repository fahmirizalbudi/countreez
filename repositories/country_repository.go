package repositories

import (
	"countreez/structs"
	"database/sql"
)

func GetAllCountries(dbParam *sql.DB) (result []structs.Country, err error) {
	sql := "SELECT id, name, iso2, iso3, capital, region, language FROM countries"

	rows, err := dbParam.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var country structs.Country

		err = rows.Scan(&country.ID, &country.Name, &country.Iso2, &country.Iso3, &country.Capital, &country.Region, &country.Language)
		if err != nil {
			panic(err)
		}

		result = append(result, country)
	}
	return
}