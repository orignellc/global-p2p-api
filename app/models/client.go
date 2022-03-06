package models

import (
	"gorm.io/gorm"
)

type (
	Client struct {
		gorm.Model
		ClientBasic `gorm:"embedded"`
		DomainWhitelists string `json:"domain_whitelists" gorm:"type:text" validate:"required,dive,required"`
	}

	ClientBasic struct {
		CompanyName      string `json:"company_name,unique" validate:"required,max=256"`
		Email            string `json:"email" gorm:"unique" validate:"required,email"`
		WebhookURL       string `json:"webhook_url" validate:"required,url"`
		CallbackURL      string `json:"callback_url" validate:"url"`
		Website          string `json:"website" validate:"required,url"`
		Bio              string	`json:"bio" gorm:"type:text" validate:"required,max=256"`
		Phone            string `json:"phone" gorm:"default:null" validate:"required,max=15"`
		Verified         bool `json:"verified" gorm:"default:false"`
		Active           bool `json:"active" gorm:"default:false"`
		SecretKey        string `json:"secret_key" gorm:"size:40,default:null"`
		PublicKey        string `json:"public_key" gorm:"size:40,default:null"`
	}
)