package handler

import (
	"fmt"
	"net/http"

	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/route/middleware"
)

// CountryList contain list of all country
type CountryList struct {
	tableName struct{} `pg:"country_list"`

	CountryID   string `pg:"country_id" json:"country_id"`
	CountryName string `pg:"country_name" json:"country_name"`
	Alpha3      string `pg:"alpha_3" json:"alpha_3"`
}

// InsertCountry ...
func InsertCountry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting User")

	//var userShortInfo model.UserShortInfo
	// countryList := &CountryList{}

	// err := e.Bind(&countryList)
	// if err != nil {
	// 	fmt.Println("err1 ---------------------")
	// 	return err
	// }

	// isAlreadyExist := comp.CheckIfExist("country_id", countryList.CountryID, countryList)
	// if isAlreadyExist {
	// 	return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, "country already exist"))
	// }

	// _, err = db.DB.Model(countryList).Insert()

	// if err != nil {
	// 	fmt.Println("---------------------------")
	// 	panic(err)
	// }

	// return e.JSON(http.StatusOK, countryList)
}

// RoleOrderPermission ...
func RoleOrderPermission(w http.ResponseWriter, r *http.Request, obj interface{}, requestedRole int) (bool, error) {

	claims, _ := middleware.GetJWTClaims(w, r)

	claimRole := claims["role"]
	if claimRole == nil && claims != nil {
		err := fmt.Errorf("invalid token")
		return false, err
	}

	var subjectRole int
	if claimRole == nil {
		subjectRole = 41
	} else {
		subjectRole = int(claimRole.(float64))
	}

	var objRoleInDB int
	err := db.DB.Model(obj).
		Where("user_id = ?user_id").
		Column("role").Select(&objRoleInDB)
	if err != nil {
		return false, err
	}

	if subjectRole > objRoleInDB || subjectRole > requestedRole {
		return false, nil
	}

	return true, nil
}
