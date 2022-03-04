package models

import (
	"gorm.io/gorm"
)

type (
	Client struct {
		gorm.Model
		ClientBasic `gorm:"embedded"`
		DomainWhitelists string `json:"domain_whitelists" gorm:"type:text"`
	}

	ClientBasic struct {
		CompanyName      string `json:"company_name"`
		Email            string `json:"email"`
		WebhookURL       string `json:"webhook_url"`
		CallbackURL      string `json:"callback_url"`
		Website          string `json:"website"`
		Bio              string	`json:"bio" gorm:"type:text"`
		Phone            string `json:"phone" gorm:"default:null"`
		Verified         bool `json:"verified" gorm:"default:false"`
		Active           bool `json:"active" gorm:"default:false"`
		SecretKey        string `json:"secret_key" gorm:"size:40"`
		PublicKey        string `json:"public_key" gorm:"size:40"`
	}
)