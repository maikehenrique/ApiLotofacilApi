package dao

import (
	"apilotofacil/db"
	"apilotofacil/model"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// Receives a lottery file and inserts it directly into the bank
func InsertLoterry(loteria model.LoteriaCaixa) (result string, err error) {
	var nameMethod = "|DaoInsertLoteria"
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO loteria (
				numero
				, tipo_jogo
				, data_apuracao
				, nome_municipio_uf_sorteio
				, dezenas_sorteadas
				, lista_rateio_premio
				, lista_municipio_ganhadores
				, acumulado
				, valor_acumulada_prox_sorteio
				, data_prox_concurso
			) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
			ON CONFLICT (numero) DO UPDATE
			SET	numero = EXCLUDED.numero
				,tipo_jogo = EXCLUDED.tipo_jogo
				,data_apuracao = EXCLUDED.data_apuracao
				,nome_municipio_uf_sorteio = EXCLUDED.nome_municipio_uf_sorteio
				,dezenas_sorteadas = EXCLUDED.dezenas_sorteadas
				,lista_rateio_premio = EXCLUDED.lista_rateio_premio
				,lista_municipio_ganhadores = EXCLUDED.lista_municipio_ganhadores
				,acumulado = EXCLUDED.acumulado
				,valor_acumulada_prox_sorteio = EXCLUDED.valor_acumulada_prox_sorteio
				,data_prox_concurso = EXCLUDED.data_prox_concurso
			RETURNING numero`
	tx, err := conn.Begin()
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
		return "", err
	}

	rows, err := tx.Query(sql,
		loteria.Numero,
		loteria.TipoJogo,
		loteria.DataApuracao,
		loteria.LocalSorteio+" - "+loteria.NomeMunicipioUFSorteio,
		convertObjectToJson(loteria.DezenasSorteadasOrdemSorteio),
		convertObjectToJson(loteria.ListaRateioPremio),
		convertObjectToJson(loteria.ListaMunicipioUFGanhadores),
		loteria.Acumulado,
		loteria.ValorAcumuladoProximoConcurso,
		loteria.DataProximoConcurso)
	if err != nil {
		log.Println(nameMethod+"|Error|", loteria.Numero, " - ", err)
		tx.Rollback()
		return "", err
	}

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			tx.Rollback()
			log.Println(nameMethod+"|Error|", err)
			return "", err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
		return "", err
	}

	return result, err
}

// Search for a lottery by ID, if it does not exist, all lotteries will be listed
func ListLottery(id int64) (result string, err error) {
	var nameMethod = "|DaoListarLoteria"

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	var row *sql.Row
	row = conn.QueryRow(`SELECT COALESCE(CAST('{"registros": ' || CAST(ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(r))) AS TEXT) || '}' AS JSONB),'{}')
							FROM (
								SELECT	numero
								, tipo_jogo "tipoJogo"
								, data_apuracao "dataApuracao"
								, nome_municipio_uf_sorteio "nomeMunicipioUFSorteio"
								, dezenas_sorteadas "dezenasSorteadasOrdemSorteio"
								, lista_rateio_premio "listaRateioPremio"
								, lista_municipio_ganhadores "listaMunicipioUFGanhadores"
								, acumulado
								, valor_acumulada_prox_sorteio "valorAcumuladoProximoConcurso"
								, data_prox_concurso "dataProximoConcurso"
							FROM loteria l 
							WHERE l.numero = (CASE WHEN $1 = 0 THEN (SELECT MAX(numero) FROM loteria) ELSE $1 END)
							)r `, id)

	err = row.Scan(&result)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
	}

	return result, err
}

func ListAllLottery() (result string, err error) {
	var nameMethod = "|DaoListAllLottery"

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	var row *sql.Row

	row = conn.QueryRow(`SELECT COALESCE(CAST('{"registros": ' || CAST(ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(r))) AS TEXT) || '}' AS JSONB),'{}')
		FROM (
			SELECT	numero
			, tipo_jogo "tipoJogo"
			, data_apuracao "dataApuracao"
			, nome_municipio_uf_sorteio "nomeMunicipioUFSorteio"
			, dezenas_sorteadas "dezenasSorteadasOrdemSorteio"
			, lista_rateio_premio "listaRateioPremio"
			, lista_municipio_ganhadores "listaMunicipioUFGanhadores"
			, acumulado
			, valor_acumulada_prox_sorteio "valorAcumuladoProximoConcurso"
			, data_prox_concurso "dataProximoConcurso"
		FROM loteria l 
		)r `)

	err = row.Scan(&result)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
	}

	return result, err
}

func ScrapingApiCaixa(concurso string) (result model.LoteriaCaixa, err error) {
	var nameMethod = "|ScrapingApiCaixa"
	var loteriaCaixa model.LoteriaCaixa

	var sql string
	if concurso != "latest" {
		_, err := strconv.Atoi(concurso)
		if err != nil {
			sql = ""
		} else {
			sql = concurso
		}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://servicebus2.caixa.gov.br/portaldeloterias/api/lotofacil/" + sql)
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
		return loteriaCaixa, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println(nameMethod+"|Error| %d %s", resp.StatusCode, resp.Status)
	}

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
	}

	err = json.Unmarshal(body, &loteriaCaixa)
	if err != nil {
		panic(err)
	}
	return loteriaCaixa, err
}

// Search the latest lottery contest
func SearchLastContestSorteOnline(contest string) (result string, err error) {
	var nameMethod = "|SearchLastContestSorteOnline"
	var sql string
	if contest != "latest" {
		_, err := strconv.Atoi(contest)
		if err != nil {
			sql = "latest"
		} else {
			sql = contest
		}
	}

	resp, err := http.Get("https://www.sorteonline.com.br/lotofacil/resultados/" + sql)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println(nameMethod+"|Error| %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div[class=DivDeVisibilidade]").Each(func(i int, s *goquery.Selection) {
		result = s.Find(".header-resultados__nro-concurso").Text()
	})

	return result, err
}

// Generates a sequence of contest numbers that are waiting to be entered in DB
func SearchContestNotEntered(id int64) (result string, err error) {
	var nameMethod = "|DaoSearchContestNotEntered"
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT COALESCE(CAST('{"registros": ' || CAST(ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(r))) AS TEXT) || '}' AS JSONB),'{}')
						FROM (
							SELECT generate_series(1,$1) concurso
						)r
						WHERE NOT EXISTS (SELECT 1 FROM loteria l WHERE l.numero = r.concurso)`, id)
	err = row.Scan(&result)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
	}
	return result, err
}

func convertObjectToJson(i interface{}) []byte {
	valueConverted, err := json.Marshal(i)
	if err != nil {
		return []byte("{}")
	}
	return valueConverted
}

// It does the whole process of updating the database, searches for the last draw drawn and performs a search in the database to see which contest does not yet exist in the bank, and only the contests that do not exist will be inserted.
func UpdateDB() error {
	var nameMethod = "|UpdateDB"
	ultimoConcurso, err := ScrapingApiCaixa("latest")
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
	}

	contestInsert, err := SearchContestNotEntered(int64(ultimoConcurso.Numero))
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
	} else {

		var numbersToInsert model.NumerosSeremInseridos
		err = json.Unmarshal([]byte(contestInsert), &numbersToInsert)
		if err != nil {
			log.Println(nameMethod+"|Error|", err)
		}
		for _, numero := range numbersToInsert.Registros {
			contestString := strconv.FormatInt(numero.Concurso, 10)
			loteria, err := ScrapingApiCaixa(contestString)
			if err != nil {
				log.Printf(nameMethod + "|Error|" + err.Error())
			} else {
				_, err := InsertLoterry(loteria)
				if err != nil {
					log.Println(nameMethod+"|Error|", err)
				} else {
					log.Println(nameMethod+"|Info|Concurso inserido com sucesso :", contestString)
				}
			}
		}
	}
	return err
}

// Fetch the dashboard data to show on the frontend
func GetDashboard(ultimos_concursos int64) (dashboard string, err error) {
	var nameMethod = "|GetDashboard"
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	row := conn.QueryRow(`SELECT f.rretorno  
					FROM f_entidade_dashboard('obter_dashboard','{"ultimos_concursos":` + strconv.FormatInt(ultimos_concursos, 10) + `}')	f`)
	err = row.Scan(&dashboard)
	if err != nil {
		log.Println(nameMethod+"|Error|", ultimos_concursos)
	}
	return
}

// Set up the lottery game and bring the results and successes
func BuildYourGame(requestDezenas model.BuildYourGame) (result string, err error) {
	var nameMethod = "|BuildYourGame"
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	row := conn.QueryRow(`SELECT rretorno 
					FROM f_entidade_generica('conferir_jogo_montado',$1)`, convertObjectToJson(requestDezenas))
	err = row.Scan(&result)
	if err != nil {
		log.Println(nameMethod+"|Error|", requestDezenas)
	}
	return
}
