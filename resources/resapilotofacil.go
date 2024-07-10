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

// Insert a lottery into the database
func InsertLoterry(c echo.Context) error {
	var nameMethod = "|InsertLoterry"
	loteria := new(model.LoteriaCaixa)

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

// List a lottery in the database
func ListLottery(c echo.Context) error {
	var nameMethod = "|ListLottery"
	var result string

	concurso := c.Param("concurso")
	if concurso == "" {
		concurso = "0"
	}

	concursoint, err := strconv.ParseInt(concurso, 10, 64)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
		return c.JSON(http.StatusBadRequest, "Concurso informado inválido.")
	}

	loteria, err := dao.ScrapingApiCaixa(concurso)
	if err != nil {
		concursos, err := dao.ListLottery(concursoint)
		if err != nil {
			log.Printf(nameMethod + "|Error|" + err.Error())
		}
		result = concursos
		valorJson := json.RawMessage(result)

		return c.JSON(http.StatusOK, valorJson)
	} else {
		var arrayResult model.ArrayLoteria
		arrayResult.Registros = append(arrayResult.Registros, loteria)
		return c.JSON(http.StatusOK, arrayResult)
	}
}

// List all lottery in the database
func ListAllLottery(c echo.Context) error {
	var nameMethod = "|ListAllLottery"
	var result string

	concursos, err := dao.ListAllLottery()
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
	}
	result = concursos
	valorJson := json.RawMessage(result)

	return c.JSON(http.StatusOK, valorJson)
}

// Search last contest drawn
func SearchLastContest(c echo.Context) error {
	var nameMethod = "|SearchLastContest"

	loteria, err := dao.ScrapingApiCaixa("latest")
	if err != nil {
		concursos, err := dao.ListLottery(0)
		if err != nil {
			log.Printf(nameMethod + "|Error|" + err.Error())
		}

		result := concursos
		valorJson := json.RawMessage(result)

		return c.JSON(http.StatusOK, valorJson)

	} else {
		_, err := dao.InsertLoterry(loteria)
		if err != nil {
			log.Println(nameMethod+"|Error|", err)
		}
		var arrayResult model.ArrayLoteria
		arrayResult.Registros = append(arrayResult.Registros, loteria)

		return c.JSON(http.StatusOK, arrayResult)
	}
}

// Dashboard search
func ListDashboard(c echo.Context) error {
	var nameMethod = "|ListDashboard"
	var result string

	ultimosConcurso := c.Param("ultimos_concurso")
	if ultimosConcurso == "" {
		ultimosConcurso = "0"
	}

	ultimosConcursoInt, err := strconv.ParseInt(ultimosConcurso, 10, 64)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
		return c.JSON(http.StatusBadRequest, "Parâmetro informado inválido para últimos concursos.")
	}

	dashboard_result, err := dao.GetDashboard(ultimosConcursoInt)
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
	}
	result = dashboard_result

	valorJson := json.RawMessage(result)
	return c.JSON(http.StatusOK, valorJson)
}

// Brings the results and hits of the assembled game
func BuildYourGame(c echo.Context) error {
	var nameMethod = "|BuildYourGame"
	var result string

	var requestDezenas = new(model.BuildYourGame)

	if err := c.Bind(&requestDezenas); err != nil { //Verifica JSON
		log.Println(nameMethod+"|Error|", err)
		return c.JSON(http.StatusBadRequest, "Erro no JSON do BODY.")
	}

	dao_result, err := dao.BuildYourGame(*requestDezenas)
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
	}
	result = dao_result

	valorJson := json.RawMessage(result)
	return c.JSON(http.StatusOK, valorJson)
}

func ValidateServerAtive(c echo.Context) error {
	return c.JSON(http.StatusOK, "Server active")
}
