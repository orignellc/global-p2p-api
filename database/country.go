package database

import (
	"gorm.io/gorm"
)

type (
	Country struct {
		gorm.Model
		Name string `json:"name" gorm:"unique,type:text" validate:"required,max=25"`
		Iso2Code string `json:"iso2_code" gorm:"size:2,unique" validate:"required"`
		Iso3Code string `json:"iso3_code" gorm:"size:3,unique" validate:"required"`
		DialCode string `json:"dial_code" gorm:"size:10,unique" validate:"required"`
	}
)

func (*Country) ModelName() string {
	return "country"
}

func (country Country) List(countries *[]Country) *gorm.DB {
	result := DB.Find(countries)

	return result
}

func (country Country) ById(_country *Country, id uint) *Country {

	DB.Find(_country,id)

	return _country
}