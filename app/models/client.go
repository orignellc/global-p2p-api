package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	CompanyName string
	Email string
	WebhookURL string
	CallbackURL string
	DomainWhitelists []string
	Website string
	Bio string
	Phone string
	Verified bool
	Active bool
	SecretKey string
	PublicKey string
}
