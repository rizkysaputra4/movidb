package model

// CountryList contain list of all country
type CountryList struct {
	tableName struct{} `pg:"country_list"`

	CountryID   string    `pg:"country_id" json:"country_id"`
	CountryName string `pg:"country_name" json:"country_name"`
	Alpha3    string `pg:"alpha_3" json:"alpha_3"`
}