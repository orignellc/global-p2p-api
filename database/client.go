package database

import (
	"gorm.io/gorm"
	"strings"
)

type (
	Client struct {
		gorm.Model
		ClientBasic      `gorm:"embedded"`
		DomainWhitelists string `json:"domain_whitelists" gorm:"type:text"`
	}

	ClientBasic struct {
		CompanyName     string `json:"company_name" gorm:"unique" validate:"required,max=256"`
		Email           string `json:"email" gorm:"unique" validate:"required,email"`
		Password        string `json:"password,omitempty" gorm:"" validate:"required,alphanum"`
		WebhookURL      string `json:"webhook_url" validate:"required,url"`
		CallbackURL     string `json:"callback_url,omitempty" validate:"url"`
		Website         string `json:"website" validate:"required,url"`
		Bio             string	`json:"bio" gorm:"type:text" validate:"required,max=256"`
		Phone           string `json:"phone" gorm:"default:null" validate:"required,max=15"`
		Verified        bool `json:"verified,omitempty" gorm:"default:false"`
		Active          bool `json:"active,omitempty" gorm:"default:false"`
		CountryCode		string `json:"country_code,omitempty"`
		Country      	Country `gorm:"foreignKey:CountryCode" validate:"-"`
		SecretKey       string `json:"secret_key,omitempty" gorm:"size:40,default:null"`
		PublicKey       string `json:"public_key,omitempty" gorm:"size:40,default:null"`
	}

	ClientData struct {
		ClientBasic
		DomainWhitelists []string `json:"domain_whitelists" validate:"required,dive,required"`
	}
)

func (*ClientData) ModelName() string {
	return "client"
}

func (client *Client) Public() *ClientData  {
	return &ClientData{
		ClientBasic: ClientBasic{
			CompanyName: client.CompanyName,
			Email: client.Email,
			WebhookURL:  client.WebhookURL,
			CallbackURL:  client.CallbackURL,
			Website:  client.Website,
			Bio:  client.Bio,
			Phone:  client.Phone,
			Verified:  client.Verified,
			Active:  client.Active,
			PublicKey:  client.PublicKey,
			Country:  client.Country,
		},
		DomainWhitelists: strings.Split(client.DomainWhitelists, ","),
	}
}

func (client Client) Insert(clientData *ClientData) *Client {
	model := Client{
		ClientBasic: clientData.ClientBasic,
		DomainWhitelists: strings.Join(clientData.DomainWhitelists, ","),
	}

	DB.Create(&model)

	DB.Preload("Country").Find(&model, model.ID)

	return &model
}

func (client Client) List(clients *[]Client) *gorm.DB {

	result := DB.Find(clients)

	return result
}

func (client Client) ById(_client *Client, id uint) *Client {

	DB.Find(_client,id)

	return _client
}