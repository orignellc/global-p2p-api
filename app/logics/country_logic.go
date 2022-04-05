package logics

import (
	"github.com/orignellc/global-p2p-api/database"
)

var country = database.Country{}

func Countries() *[]database.Country {

	var countries = new([]database.Country)

	country.List(countries)

	return countries
}

func GetCountryById(id uint) *database.Country {
	var _country = new(database.Country)

	country.ById(_country, id)

	return _country
}

