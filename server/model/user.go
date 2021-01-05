package model

// UserShortInfo ...
type UserShortInfo struct {
	tableName struct{} `pg:"user_short_info"`

	UserID       int    `pg:"user_id" json:"user_id"`
	UserFullName string `pg:"user_full_name" json:"user_full_name"`
	CountryID    int `pg:"country_id" json:"country_id"`
}

// UserInformation contains all full informations about user
type UserInformation struct {
	tableName struct{} `pg:"user_information"`

	UserID      int    `pg:"user_id" json:"user_id"`
	Birthdate   string `pg:"birthdate" json:"birthdate"`
	SignUpDate  int `pg:"signup_date" json:"signup_date"`
	Bio         string `pg:"bio" json:"bio"`
	FBLink      string `pg:"fb_link" json:"fb_link"`
	TwitterLink string `pg:"twitter_link" json:"twitter_link"`
	IGLink      string `pg:"ig_link" json:"ig_link"`
	Sex         bool   `pg:"sex" json:"sex"`
	LastRequest int `pg:"last_request" json:"last_request"`
}