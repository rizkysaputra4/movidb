package model

// MovieType is identifier of a moview whether
// a movie, series, kdrama, jdrama, etc.
type MovieType struct {
	tableName struct{} `pg:"movie_type"`

	TypeID   int    `pg:"type_id" json:"type_id"`
	TypeName string `pg:"type_name" json:"type_name"`
}
