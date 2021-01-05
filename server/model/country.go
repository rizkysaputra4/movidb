package model

// CountryList contain list of all country
type CountryList struct {
	tableName struct{} `pg:"country_list"`

	CountryID   int    `pg:"country_id" json:"country_id"`
	CountryName string `pg:"country_name" json:"country_name"`
	FlagLink    string `pg:"flag_link" json:"flag_link"`
}