package logics

import (
	"github.com/orignellc/global-p2p-api/database"
	"github.com/orignellc/global-p2p-api/pkg/util"
)

var client = database.Client{}

func NewClient(clientData *database.ClientData) (*database.Client, *util.ErrorWrapper) {

	var client database.Client
	errors := util.ValidateStruct(clientData)

	if errors != nil {
		return &client, &util.ErrorWrapper{
			Type: util.ValidationError,
			Details: errors,
		}
	}

	data := client.Insert(clientData)

	return data, nil
}

func Clients() *[]database.Client {

	var clients = new([]database.Client)

	client.List(clients)

	return clients
}

func GetClientById(id uint) *database.Client {
	var _client = new(database.Client)

	client.ById(_client, id)

	return _client
}

