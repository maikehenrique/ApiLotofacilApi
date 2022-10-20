package model

import "log"

var LogXmlJson log.Logger

const (
	DiretorioLog = "log"
	ArquivoLog   = "apilotofacil.log"
)

type Dashboard struct {
	UltimosConcursos int64 `json:"ultimos_concursos"`
}

type NumerosSeremInseridos struct {
	Registros []struct {
		Concurso int64 `json:"concurso"`
	} `json:"registros"`
}

type EstadosPremiacoesSorte []struct {
	CodigoFaixa int    `json:"CodigoFaixa"`
	Uf          string `json:"SiglaEstado"`
	Latitude    string `json:"Latitude"`
	Longitude   string `json:"Longitude"`
	Nome        string `json:"NomeEstado"`
	Vencedores  int    `json:"Quantidade"`
	Cidades     []struct {
		CodigoCidade int    `json:"CodigoCidade"`
		Cidade       string `json:"NomeCidade"`
		Vencedores   int    `json:"Quantidade"`
		Latitude     string `json:"Latitude"`
		Longitude    string `json:"Longitude"`
	} `json:"PremiacaoPorCidade"`
}

type Loteria struct {
	Loteria    string   `json:"loteria"`
	Nome       string   `json:"nome"`
	Concurso   int      `json:"concurso"`
	Data       string   `json:"data"`
	Local      string   `json:"local"`
	Dezenas    []string `json:"dezenas"`
	Premiacoes []struct {
		Acertos    string `json:"acertos"`
		Vencedores int    `json:"vencedores"`
		Premio     string `json:"premio"`
	} `json:"premiacoes"`
	EstadosPremiados []struct {
		Nome       string `json:"nome"`
		Uf         string `json:"uf"`
		Vencedores string `json:"vencedores"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
		Cidades    []struct {
			Cidade     string `json:"cidade"`
			Vencedores string `json:"vencedores"`
			Latitude   string `json:"latitude"`
			Longitude  string `json:"longitude"`
		} `json:"cidades"`
	} `json:"estadosPremiados"`
	Acumulou              bool   `json:"acumulou"`
	AcumuladaProxConcurso string `json:"acumuladaProxConcurso"`
	DataProxConcurso      string `json:"dataProxConcurso"`
	ProxConcurso          int    `json:"proxConcurso"`
}

type ArrayLoteria struct {
	Registros []struct {
		Loteria    string   `json:"loteria"`
		Nome       string   `json:"nome"`
		Concurso   int      `json:"concurso"`
		Data       string   `json:"data"`
		Local      string   `json:"local"`
		Dezenas    []string `json:"dezenas"`
		Premiacoes []struct {
			Acertos    string `json:"acertos"`
			Vencedores int    `json:"vencedores"`
			Premio     string `json:"premio"`
		} `json:"premiacoes"`
		EstadosPremiados []struct {
			Nome       string `json:"nome"`
			Uf         string `json:"uf"`
			Vencedores string `json:"vencedores"`
			Latitude   string `json:"latitude"`
			Longitude  string `json:"longitude"`
			Cidades    []struct {
				Cidade     string `json:"cidade"`
				Vencedores string `json:"vencedores"`
				Latitude   string `json:"latitude"`
				Longitude  string `json:"longitude"`
			} `json:"cidades"`
		} `json:"estadosPremiados"`
		Acumulou              bool   `json:"acumulou"`
		AcumuladaProxConcurso string `json:"acumuladaProxConcurso"`
		DataProxConcurso      string `json:"dataProxConcurso"`
		ProxConcurso          int    `json:"proxConcurso"`
	} `json:"registros"`
}

type EstadosPremiados struct {
	Nome       string `json:"nome"`
	Uf         string `json:"uf"`
	Vencedores string `json:"vencedores"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Cidades    []struct {
		Cidade     string `json:"cidade"`
		Vencedores string `json:"vencedores"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
	} `json:"cidades"`
}

type Premiacoes struct {
	Acertos    string `json:"acertos"`
	Vencedores int    `json:"vencedores"`
	Premio     string `json:"premio"`
}

type Cidade struct {
	Cidade     string `json:"cidade"`
	Vencedores string `json:"vencedores"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}

type BuildYourGame struct {
	Dezenas []int `json:"dezenas"`
}
