package model

// Identifier is contains information whether it is deleted, under review, blocked, etc.
type Identifier struct {
	tableName struct{} `pg:"identifier"`

	IdentifierID int    `pg:"identifier_id" json:"identifier_id"`
	Info         string `pg:"info" json:"info"`
}
