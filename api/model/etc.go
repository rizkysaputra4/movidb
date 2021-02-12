package model

// Identifier is contains information whether it is deleted, under review, blocked, etc.
type Identifier struct {
	tableName struct{} `pg:"identifier"`

	IdentifierID int    `pg:"identifier_id" json:"identifier_id"`
	Info         string `pg:"info" json:"info"`
}

// Country struct contains country table
type Country struct {
	tableName struct{} `pg:"country_list"`

	CountryName string `pg:"country_name" json:"N"`
	CountryID   string `pg:"country_id,pk" json:"I"`
}
