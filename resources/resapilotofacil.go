package resources

import (
	"apilotofacil/dao"
	"apilotofacil/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Insert a lottery into the database
func InsertLoterry(c echo.Context) error {
	var nameMethod = "|InsertLoterry"
	loteria := new(model.Loteria)

	if err := c.Bind(loteria); err != nil { //Verifica JSON
		log.Println(nameMethod+"|Error|", err)
		return c.JSON(http.StatusBadRequest, "Erro no JSON do BODY.")
	}
	_, err := dao.InsertLoterry(*loteria)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
	}
	return c.JSON(http.StatusOK, loteria)
}

//List a lottery in the database
func ListLottery(c echo.Context) error {
	var nameMethod = "|ListLottery"
	var result string

	concurso := c.Param("concurso")
	if concurso == "" {
		concurso = "0"
	}

	concursoint, err := strconv.Atoi(concurso)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
		return c.JSON(http.StatusBadRequest, "Concurso informado inválido.")
	}

	concursos, err := dao.ListLottery(concursoint)
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
	}
	result = concursos

	valorJson := json.RawMessage(result)
	return c.JSON(http.StatusOK, valorJson)
}

//Search last contest drawn
func SearchLastContest(c echo.Context) error {
	var nameMethod = "|SearchLastContest"
	log.Println("nameMethod:", nameMethod)
	loteria, err := dao.ScrapingSorteOnline("latest")
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
	} else {
		_, err := dao.InsertLoterry(loteria)
		if err != nil {
			log.Println(nameMethod+"|Error|", err)
		}
	}
	var arrayResult model.ArrayLoteria
	arrayResult.Registros = append(arrayResult.Registros, loteria)

	return c.JSON(http.StatusOK, arrayResult)
}
