package model

// MovieType is identifier of a moview whether
// a movie, series, kdrama, jdrama, etc.
type MovieType struct {
	tableName struct{} `pg:"movie_type"`

	TypeID   int    `pg:"type_id" json:"type_id"`
	TypeName string `pg:"type_name" json:"type_name"`
}

// MovieInformation contains all value about movie
type MovieInformation struct {
	tableName struct{} `pg:"movie_information"`

	MovieID           int     `pg:"movie_id" json:"movie_id"`
	MovieTitle        string  `pg:"movie_title" json:"movie_title"`
	MovieSynopsis     string  `pg:"movie_synopsis" json:"movie_synopsis"`
	ReleaseDate       string  `pg:"release_date" json:"release_date"`
	IMDBRating        float32 `pg:"imdb_rating" json:"imdb_rating"`
	IMDBNumbVote      int     `pg:"imdb_numb_vote" json:"imdb_numb_vote"`
	MetaCritics       float32 `pg:"metacritics" json:"metacritics"`
	SiteRating        float32 `pg:"site_rating" json:"site_rating"`
	SiteNumbVote      int     `pg:"site_numb_vote" json:"site_numb_vote"`
	PosterLink        string  `pg:"poster_link" json:"poster_link"`
	TrailerLink       string  `pg:"trailer_link" json:"trailer_link"`
	Duration          string  `pg:"duration" json:"duration"`
	Awards            string  `pg:"awards" json:"awards"`
	TypeID            int     `pg:"type_id" json:"type_id"`
	CountryID         string  `pg:"country_id" json:"country_id"`
	Language          string  `pg:"language" json:"language"`
	UniqueLink        string  `pg:"unique_link" json:"unique_link"`
	IdentifierID      int     `pg:"identifier_id" json:"identifier_id"`
	OverallRating     float32 `pg:"overall_rating" json:"overall_rating"`
	Popularity        int     `pg:"popularity" json:"popularity"`
	MonthlyPopularity int     `pg:"monthly_popularity" json:"monthly_popularity"`
	WeeklyPopularity  int     `pg:"weekly_popularity" json:"weekly_popularity"`
	DailtyPopularity  int     `pg:"daily_popularity" json:"daily_popularity"`
}

// MoviePeopleRole contains model for people who work in movie industry
type MoviePeopleRole struct {
	tableName struct{} `pg:"role_list"`

	RoleID   int8   `pg:"role_id" json:"role_id"`
	RoleName string `pg:"role_name" json:"role_name"`
}

// MovieCharacter table that contain all character in a movie
type MovieCharacter struct {
	tableName struct{} `pg:"movie_character"`

	MovieID     int    `pg:"movie_id" json:"movie_id"`
	PersonID    int    `pg:"person_id" json:"person_id"`
	PictureLink string `pg:"picture_link" json:"picture_link"`
	EpsID       int    `pg:"eps_id" json:"eps_id"`
	Lead        int    `pg:"lead" json:"lead"`
	Character   string `pg:"character" json:"character"`
	Info        string `pg:"info" json:"info"`
}
