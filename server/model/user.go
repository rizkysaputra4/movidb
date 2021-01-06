package model

import (
	"time"
)

// UserShortInfo ...
type UserShortInfo struct {
	tableName struct{} `pg:"user_short_info"`

	UserID      int    `pg:"user_id" json:"user_id"`
	UserName    string `pg:"user_name" json:"user_name"`
	CountryID   string `pg:"country_id" json:"country_id"`
	Password    string `pg:"password" json:"password"`
	Email       string `pg:"email" json:"email"`
	LastRequest string `pg:"last_request" json:"last_request"`
}

// UserInformation contains all full informations about user
type UserInformation struct {
	tableName struct{} `pg:"user_information"`

	UserID      int       `pg:"user_id" json:"user_id"`
	Birthdate   time.Time `pg:"birthdate" json:"birthdate"`
	SignUpDate  int       `pg:"signup_date" json:"signup_date"`
	Bio         string    `pg:"bio" json:"bio"`
	FBLink      string    `pg:"fb_link" json:"fb_link"`
	TwitterLink string    `pg:"twitter_link" json:"twitter_link"`
	IGLink      string    `pg:"ig_link" json:"ig_link"`
	Sex         bool      `pg:"sex" json:"sex"`
	LastRequest time.Time `pg:"last_request" json:"last_request"`
}
