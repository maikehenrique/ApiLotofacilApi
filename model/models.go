package model

import "log"

var LogXmlJson log.Logger

type Dashboard struct {
	UltimosConcursos int64 `json:"ultimos_concursos"`
}

type NumerosSeremInseridos struct {
	Registros []struct {
		Concurso int64 `json:"concurso"`
	} `json:"registros"`
}

type ArrayLoteria struct {
	Registros []struct {
		Acumulado                    bool     `json:"acumulado"`
		DataApuracao                 string   `json:"dataApuracao"`
		DataProximoConcurso          string   `json:"dataProximoConcurso"`
		DezenasSorteadasOrdemSorteio []string `json:"dezenasSorteadasOrdemSorteio"`
		ExibirDetalhamentoPorCidade  bool     `json:"exibirDetalhamentoPorCidade"`
		ID                           any      `json:"id"`
		IndicadorConcursoEspecial    int      `json:"indicadorConcursoEspecial"`
		ListaDezenas                 []string `json:"listaDezenas"`
		ListaDezenasSegundoSorteio   any      `json:"listaDezenasSegundoSorteio"`
		ListaMunicipioUFGanhadores   []struct {
			Vencedores     int    `json:"ganhadores"`
			Nome           string `json:"municipio"`
			NomeFatansiaUL string `json:"nomeFatansiaUL"`
			Posicao        int    `json:"posicao"`
			Serie          string `json:"serie"`
			Uf             string `json:"uf"`
		} `json:"listaMunicipioUFGanhadores"`
		ListaRateioPremio []struct {
			Acertos    string  `json:"descricaoFaixa"`
			Faixa      int     `json:"faixa"`
			Vencedores int     `json:"numeroDeGanhadores"`
			Premio     float64 `json:"valorPremio"`
		} `json:"listaRateioPremio"`
		ListaResultadoEquipeEsportiva  any     `json:"listaResultadoEquipeEsportiva"`
		LocalSorteio                   string  `json:"localSorteio"`
		NomeMunicipioUFSorteio         string  `json:"nomeMunicipioUFSorteio"`
		NomeTimeCoracaoMesSorte        string  `json:"nomeTimeCoracaoMesSorte"`
		Numero                         int     `json:"numero"`
		NumeroConcursoAnterior         int     `json:"numeroConcursoAnterior"`
		NumeroConcursoFinal05          int     `json:"numeroConcursoFinal_0_5"`
		NumeroConcursoProximo          int     `json:"numeroConcursoProximo"`
		NumeroJogo                     int     `json:"numeroJogo"`
		Observacao                     string  `json:"observacao"`
		PremiacaoContingencia          any     `json:"premiacaoContingencia"`
		TipoJogo                       string  `json:"tipoJogo"`
		TipoPublicacao                 int     `json:"tipoPublicacao"`
		UltimoConcurso                 bool    `json:"ultimoConcurso"`
		ValorArrecadado                float64 `json:"valorArrecadado"`
		ValorAcumuladoConcurso05       float64 `json:"valorAcumuladoConcurso_0_5"`
		ValorAcumuladoConcursoEspecial float64 `json:"valorAcumuladoConcursoEspecial"`
		ValorAcumuladoProximoConcurso  float64 `json:"valorAcumuladoProximoConcurso"`
		ValorEstimadoProximoConcurso   float64 `json:"valorEstimadoProximoConcurso"`
		ValorSaldoReservaGarantidora   float64 `json:"valorSaldoReservaGarantidora"`
		ValorTotalPremioFaixaUm        float64 `json:"valorTotalPremioFaixaUm"`
	} `json:"registros"`
}

type BuildYourGame struct {
	Dezenas []int `json:"dezenas"`
}

type LoteriaCaixa struct {
	Acumulado                    bool     `json:"acumulado"`
	DataApuracao                 string   `json:"dataApuracao"`
	DataProximoConcurso          string   `json:"dataProximoConcurso"`
	DezenasSorteadasOrdemSorteio []string `json:"dezenasSorteadasOrdemSorteio"`
	ExibirDetalhamentoPorCidade  bool     `json:"exibirDetalhamentoPorCidade"`
	ID                           any      `json:"id"`
	IndicadorConcursoEspecial    int      `json:"indicadorConcursoEspecial"`
	ListaDezenas                 []string `json:"listaDezenas"`
	ListaDezenasSegundoSorteio   any      `json:"listaDezenasSegundoSorteio"`
	ListaMunicipioUFGanhadores   []struct {
		Vencedores     int    `json:"ganhadores"`
		Nome           string `json:"municipio"`
		NomeFatansiaUL string `json:"nomeFatansiaUL"`
		Posicao        int    `json:"posicao"`
		Serie          string `json:"serie"`
		Uf             string `json:"uf"`
	} `json:"listaMunicipioUFGanhadores"`
	ListaRateioPremio []struct {
		Acertos    string  `json:"descricaoFaixa"`
		Faixa      int     `json:"faixa"`
		Vencedores int     `json:"numeroDeGanhadores"`
		Premio     float64 `json:"valorPremio"`
	} `json:"listaRateioPremio"`
	ListaResultadoEquipeEsportiva  any     `json:"listaResultadoEquipeEsportiva"`
	LocalSorteio                   string  `json:"localSorteio"`
	NomeMunicipioUFSorteio         string  `json:"nomeMunicipioUFSorteio"`
	NomeTimeCoracaoMesSorte        string  `json:"nomeTimeCoracaoMesSorte"`
	Numero                         int     `json:"numero"`
	NumeroConcursoAnterior         int     `json:"numeroConcursoAnterior"`
	NumeroConcursoFinal05          int     `json:"numeroConcursoFinal_0_5"`
	NumeroConcursoProximo          int     `json:"numeroConcursoProximo"`
	NumeroJogo                     int     `json:"numeroJogo"`
	Observacao                     string  `json:"observacao"`
	PremiacaoContingencia          any     `json:"premiacaoContingencia"`
	TipoJogo                       string  `json:"tipoJogo"`
	TipoPublicacao                 int     `json:"tipoPublicacao"`
	UltimoConcurso                 bool    `json:"ultimoConcurso"`
	ValorArrecadado                float64 `json:"valorArrecadado"`
	ValorAcumuladoConcurso05       float64 `json:"valorAcumuladoConcurso_0_5"`
	ValorAcumuladoConcursoEspecial float64 `json:"valorAcumuladoConcursoEspecial"`
	ValorAcumuladoProximoConcurso  float64 `json:"valorAcumuladoProximoConcurso"`
	ValorEstimadoProximoConcurso   float64 `json:"valorEstimadoProximoConcurso"`
	ValorSaldoReservaGarantidora   float64 `json:"valorSaldoReservaGarantidora"`
	ValorTotalPremioFaixaUm        float64 `json:"valorTotalPremioFaixaUm"`
}
