package model

import (
	"time"

	"github.com/rizkysaputra4/moviwiki/api/db"
)

// UserShortInfo ...
type UserShortInfo struct {
	tableName struct{} `pg:"user_short_info"`

	UserID      int    `pg:"user_id, pk" json:"user_id,omitempty"`
	UserName    string `pg:"user_name" json:"user_name,omitempty"`
	CountryID   string `pg:"country_id" json:"country_id,omitempty"`
	Password    string `pg:"password" json:"password,omitempty"`
	Email       string `pg:"email" json:"email,omitempty"`
	Role        int    `pg:"role" json:"role,omitempty"`
	LastRequest string `pg:"last_request" json:"last_request,omitempty"`
}

// AdminInfo ...
type AdminInfo struct {
	tableName struct{} `pg:"user_short_info"`

	UserID    int    `pg:"user_id" json:"user_id ,omitempty"`
	UserName  string `pg:"user_name" json:"user_name,omitempty"`
	CountryID string `pg:"country_id" json:"country_id,omitempty"`
	Password  string `pg:"password" json:"password,omitempty"`
	Email     string `pg:"email" json:"email,omitempty"`
	Role      int    `pg:"role" json:"role,omitempty"`

	LastRequest string `pg:"last_request" json:"last_request,omitempty"`
}

// UserInformation contains all full informations about user
type UserInformation struct {
	tableName struct{} `pg:"user_information"`

	UserID                  int    `pg:",pk"`
	UserFullName            string `pg:"user_full_name" json:"user_full_name,omitempty"`
	Birthdate               string `pg:"birthdate" json:"birthdate,omitempty"`
	RegisterDate            string `pg:"signup_date" json:"signup_date,omitempty"`
	Bio                     string `pg:"bio" json:"bio,omitempty"`
	FBLink                  string `pg:"fb_link" json:"fb_link,omitempty"`
	TwitterLink             string `pg:"twitter_link" json:"twitter_link,omitempty"`
	IGLink                  string `pg:"ig_link" json:"ig_link,omitempty"`
	Sex                     bool   `pg:"sex" json:"sex"`
	ContributorPoint        int    `pg:"contributor_points" json:"contributor_points,omitempty"`
	MonthlyContributorPoint int    `pg:"monthly_contributor_points" json:"monthly_contributor_points,omitempty"`
}

// UpdateLastRequest ...
func (u *UserShortInfo) UpdateLastRequest() error {

	u.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")
	_, err := db.DB.Model(u).
		Where("user_id = ?user_id").
		Column("last_request").
		Update()

	if err != nil {
		return err
	}

	return nil
}
