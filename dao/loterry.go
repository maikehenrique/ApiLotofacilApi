package dao

import (
	"apilotofacil/db"
	"apilotofacil/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//Receives a lottery file and inserts it directly into the bank
func InsertLoterry(loteria model.Loteria) (result string, err error) {
	var nameMethod = "|DaoInsertLoteria"
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO loteria(concurso
						,loteria
						,nome
						,data
						,local
						,dezenas
						,premiacoes
						,estados_premiados
						,acumulou
						,acumulada_prox_concurso
						,data_prox_concurso
			) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING concurso`
	tx, err := conn.Begin()
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
		return "", err
	}

	rows, err := tx.Query(sql,
		loteria.Concurso,
		loteria.Loteria,
		loteria.Nome,
		loteria.Data,
		loteria.Local,
		convertObjectToJson(loteria.Dezenas),
		convertObjectToJson(loteria.Premiacoes),
		convertObjectToJson(loteria.EstadosPremiados),
		loteria.Acumulou,
		loteria.AcumuladaProxConcurso,
		loteria.DataProxConcurso)
	if err != nil {
		log.Println(nameMethod+"|Error| ", loteria.Concurso, "- ", err)
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

//Search for a lottery by ID, if it does not exist, all lotteries will be listed
func ListLottery(id int) (result string, err error) {
	var nameMethod = "|DaoListarLoteria"
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	var row *sql.Row
	if id != 0 {
		row = conn.QueryRow(`SELECT COALESCE(CAST('{"registros": ' || CAST(ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(r))) AS TEXT) || '}' AS JSONB),'{}')
							FROM (
								SELECT  *
								FROM loteria WHERE concurso = $1
							)r `, id)
	} else {
		row = conn.QueryRow(`SELECT COALESCE(CAST('{"registros": ' || CAST(ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(r))) AS TEXT) || '}' AS JSONB),'{}')
		FROM (
			SELECT  *
			FROM loteria
		)r `)
	}

	err = row.Scan(&result)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
	}
	return result, err
}

//Performs the extraction of lottery data directly from the website www.sorteonline.com.br
func ScrapingSorteOnline(concurso string) (result model.Loteria, err error) {
	var nameMethod = "|DaoScrapingSorteOnline"
	var sql string
	if concurso != "latest" {
		_, err := strconv.Atoi(concurso)
		if err != nil {
			sql = "latest"
		} else {
			sql = concurso
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

	dataAtual := time.Now()
	var loteria model.Loteria
	doc.Find("div[class=DivDeVisibilidade]").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title .name").Text()
		concurso := s.Find(".header-resultados__nro-concurso").Text()
		local := s.Find(".header-resultados__local-sorteio").Text()
		data := s.Find(".header-resultados__datasorteio").Text()
		data = strings.ToLower(data)
		if data == "hoje" {
			data = fmt.Sprintf("%d/%02d/%d", dataAtual.Day(), dataAtual.Month(), dataAtual.Year())
		}

		loteria.Nome = title
		loteria.Loteria = "lotofacil"
		loteria.Concurso, _ = strconv.Atoi(concurso)
		loteria.Local = local
		loteria.Data = data

		//Dezenas
		var arraydezenas []string
		s.Find("li[class=bg]").Each(func(i int, s *goquery.Selection) {
			arraydezenas = append(arraydezenas, s.Text())
		})
		loteria.Dezenas = arraydezenas

		//Premiações
		s.Find(".block-table .result .tr").Each(func(i int, s *goquery.Selection) {
			var premiacao model.Premiacoes
			s.Find(".td").Each(func(i int, s *goquery.Selection) {

				sampleRegexp := regexp.MustCompile(`[^\d]`)
				if i == 0 /*Acertos*/ {
					premiacao.Acertos = s.Text()

				}

				if i == 1 /*Ganhadores*/ {
					input := s.Text()
					result := sampleRegexp.ReplaceAllString(input, "")
					premiacao.Vencedores, _ = strconv.Atoi(result)
				}

				if i == 2 /*Prêmio*/ {
					premiacao.Premio = s.Text()
					loteria.Premiacoes = append(loteria.Premiacoes, premiacao)
				}
			})
		})

		// Estados premiados
		s.Find(".button-win").Each(func(i int, s *goquery.Selection) {
			estados, _ := s.Attr("data-estados-premiados")
			var estadospremiacoessorte model.EstadosPremiacoesSorte

			json.Unmarshal([]byte(estados), &estadospremiacoessorte)
			for _, k := range estadospremiacoessorte {
				var estadopremiacao model.EstadosPremiados

				estadopremiacao.Nome = k.Nome
				estadopremiacao.Uf = k.Uf
				estadopremiacao.Vencedores = fmt.Sprintf("%d", k.Vencedores)
				estadopremiacao.Latitude = k.Latitude
				estadopremiacao.Longitude = k.Longitude
				for _, j := range k.Cidades {
					var cidade model.Cidade
					cidade.Cidade = j.Cidade
					cidade.Latitude = j.Latitude
					cidade.Longitude = j.Longitude
					cidade.Vencedores = fmt.Sprintf("%d", j.Vencedores)
					estadopremiacao.Cidades = append(estadopremiacao.Cidades, cidade)
				}
				loteria.EstadosPremiados = append(loteria.EstadosPremiados, estadopremiacao)
			}

		})

		acumulou := s.Find(".acumulado").Size() > 0
		acumuladoProxConcurso := strings.Replace(s.Find(".estimative .value").Text(), "\n", "", -1)
		acumuladoProxConcurso = strings.Replace(acumuladoProxConcurso, "R$", "R$ ", -1)
		dataProxConcurso := s.Find(".foother-resultados__data-sorteio").Text()
		dataProxConcurso = strings.ToLower(dataProxConcurso)

		if dataProxConcurso != "" {
			if dataProxConcurso == "hoje" {
				dataProxConcurso = fmt.Sprintf("%d/%02d/%d", dataAtual.Day(), dataAtual.Month(), dataAtual.Year())
			} else if dataProxConcurso == "amanhã" {
				dataProxConcurso = fmt.Sprintf("%d/%02d/%d", dataAtual.Day()+1, dataAtual.Month(), dataAtual.Year())
			}
		}

		loteria.Acumulou = acumulou
		loteria.AcumuladaProxConcurso = acumuladoProxConcurso
		loteria.DataProxConcurso = dataProxConcurso
		loteria.ProxConcurso = loteria.Concurso + 1

		result = loteria
		//valorConvertido, _ := json.Marshal(loteria)
		//fmt.Println("valorConvertido:", string(valorConvertido))
	})

	return result, err
}

//Search the latest lottery contest
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

//Generates a sequence of contest numbers that are waiting to be entered in DB
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
						WHERE NOT EXISTS (SELECT 1 FROM loteria l WHERE l.concurso = r.concurso)`, id)
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

//It does the whole process of updating the database, searches for the last draw drawn and performs a search in the database to see which contest does not yet exist in the bank, and only the contests that do not exist will be inserted.
func UpdateDB() error {
	var nameMethod = "|UpdateDB"
	contest, err := SearchLastContestSorteOnline("latest")
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
	}

	log.Println(nameMethod+"|Info|Updating DB :", contest)

	contestInt, err := strconv.ParseInt(contest, 10, 64)
	if err != nil {
		log.Println(nameMethod+"|Error|", err)
	}

	contestInsert, err := SearchContestNotEntered(contestInt)
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
			loteria, err := ScrapingSorteOnline(contestString)
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
