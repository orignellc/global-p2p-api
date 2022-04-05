package database

import (
	"encoding/json"
	"fmt"
	"github.com/orignellc/global-p2p-api/pkg/util"
	zeroLog "github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"strings"
)

func RunSeeder() {
	//seedCountries()
	seedClients()
}

func seedCountries() {
	var countries = new([]Country)

	country  := Country{}

	result := country.List(countries)

	if result.RowsAffected == 0 {
		seederFile, err := os.Open(fmt.Sprintf("%s/countries.json", util.SeederPath()))

		if err != nil{
			panic(err.Error())
		}

		defer func(seederFile *os.File) {
			err := seederFile.Close()
			if err != nil {
				panic("Cannot close seeder file")
			}
		}(seederFile)

		byteValue, err := ioutil.ReadAll(seederFile)

		if err != nil{
			panic(err.Error())
		}

		err = json.Unmarshal(byteValue, &countries)
		if err != nil {
			zeroLog.Err(err).
				Str("File Path", "country.json").
				Msg("Unable to Unmarshal countries")
		}

		for _, country := range *countries {
			DB.Create(&country)
		}
	}
}

func seedClients() {
	var clients = new([]Client)
	var clientsBasic = new([]ClientBasic)

	client  := Client{}

	result := client.List(clients)

	if result.RowsAffected == 0 {
		filePath := fmt.Sprintf("%s/clients.json", util.SeederPath())
		seederFile, err := os.Open(filePath)

		if err != nil{
			panic(err.Error())
		}

		defer func(seederFile *os.File) {
			err := seederFile.Close()
			if err != nil {
				panic("Cannot close seeder file")
			}
		}(seederFile)

		byteValue, err := ioutil.ReadAll(seederFile)

		if err != nil{
			panic(err.Error())
		}

		err = json.Unmarshal(byteValue, &clientsBasic)
		if err != nil {
			zeroLog.Err(err).
				Str("File Path", filePath).
				Msg("Unable to Unmarshal countries")
		}

		for _, clientBasic := range *clientsBasic {
			model := Client{
				ClientBasic: clientBasic,
				DomainWhitelists: strings.Join([]string{
					"http://127.0.0.1",
					"http://localhost",
					"https://piggyfi.africa",
				}, ","),
			}

			DB.Create(&model)
		}
	}
}
