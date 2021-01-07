package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
)

// InsertCountry ...
func InsertCountry(e echo.Context) error {
	fmt.Println("Inserting User")

	//var userShortInfo model.UserShortInfo
	countryList := &model.CountryList{}

	err := e.Bind(&countryList)

	if err != nil {
		fmt.Println("err1 ---------------------")
		return err
	}

	_, err = db.DB.Model(countryList).Insert()

	if err != nil {
		fmt.Println("---------------------------")
		panic(err)
	}

	return e.JSON(http.StatusOK, countryList)
}
