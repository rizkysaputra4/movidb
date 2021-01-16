package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/comp"
	"github.com/rizkysaputra4/moviwiki/server/db"
)

// CountryList contain list of all country
type CountryList struct {
	tableName struct{} `pg:"country_list"`

	CountryID   string `pg:"country_id" json:"country_id"`
	CountryName string `pg:"country_name" json:"country_name"`
	Alpha3      string `pg:"alpha_3" json:"alpha_3"`
}

// InsertCountry ...
func InsertCountry(e echo.Context) error {
	fmt.Println("Inserting User")

	//var userShortInfo model.UserShortInfo
	countryList := &CountryList{}

	err := e.Bind(&countryList)
	if err != nil {
		fmt.Println("err1 ---------------------")
		return err
	}

	isAlreadyExist := comp.CheckIfExist("country_id", countryList.CountryID, countryList)
	if isAlreadyExist {
		return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, "country already exist"))
	}

	_, err = db.DB.Model(countryList).Insert()

	if err != nil {
		fmt.Println("---------------------------")
		panic(err)
	}

	return e.JSON(http.StatusOK, countryList)
}
