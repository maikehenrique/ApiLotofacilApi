<p align="center">
  <img src="https://user-images.githubusercontent.com/48946749/147809259-e7b15a3b-2e90-42c2-abaf-a6cacdc77e03.png">
  <h2 align="center">API Lotofácil CAIXA</h2>
  <p align="left">
    Free API to track Lotofácil lottery results <a href="http://loterias.caixa.gov.br/wps/portal/loterias">Lotteries CAIXA</a>.<br>
    Developed as a way to improve and establish knowledge, it seeks to be a free, robust and open tool for those who want to carry out implementations like me.<br>
    I use the API created to create a free statistics website and help those who play specifically in the Lotofácil lottery, I leave the link for those who want to access and test it <a href="https://api-lotofacil-api-builded.vercel.app/">Sorte Lotofácil</a>.<br>
    With time I will be improving both API and my statistics site.<br>
  </p>

</p>

## Return Examples
Currently, the database contains the lotofácil games but can be easily adapted to all other Caixa lotteries. It hasn't been adapted yet because my main focus is the lotofácil due to creating a specific statistics site for the same

* **Latest Result**

[GET]
```https://sortelotofacil.herokuapp.com/api/lotofacil/latest```

Return Examples: 

```json
{
    "registros": [
        {
            "loteria": "lotofacil",
            "nome": "Lotofácil",
            "concurso": 2642,
            "data": "20/10/2022",
            "local": "ESPAÇODASORTE em SÃOPAULO,SP",
            "dezenas": [
                "02",
                "04",
                "05",
                "08",
                "09",
                "13",
                "14",
                "15",
                "16",
                "17",
                "19",
                "20",
                "21",
                "22",
                "24"
            ],
            "premiacoes": [
                {
                    "acertos": "15 Pontos",
                    "vencedores": 3,
                    "premio": "505.605,73"
                },
                {
                    "acertos": "14 Pontos",
                    "vencedores": 275,
                    "premio": "1.652,16"
                },
                {
                    "acertos": "13 Pontos",
                    "vencedores": 9355,
                    "premio": "25,00"
                },
                {
                    "acertos": "12 Pontos",
                    "vencedores": 113258,
                    "premio": "10,00"
                },
                {
                    "acertos": "11 Pontos",
                    "vencedores": 620320,
                    "premio": "5,00"
                }
            ],
            "estadosPremiados": [
                {
                    "nome": "Goiás",
                    "uf": "GO",
                    "vencedores": "1",
                    "latitude": "36.7156122",
                    "longitude": "-95.9435511",
                    "cidades": [
                        {
                            "cidade": "Heitoraí",
                            "vencedores": "1",
                            "latitude": "-15.7227332",
                            "longitude": "-49.82713750000001"
                        }
                    ]
                },
                {
                    "nome": "Minas Gerais",
                    "uf": "MG",
                    "vencedores": "1",
                    "latitude": "-18.512178",
                    "longitude": "-44.5550308",
                    "cidades": [
                        {
                            "cidade": "Muzambinho",
                            "vencedores": "1",
                            "latitude": "-21.3712593",
                            "longitude": "-46.5232197"
                        }
                    ]
                },
                {
                    "nome": "Pará",
                    "uf": "PA",
                    "vencedores": "1",
                    "latitude": "-1.9981271",
                    "longitude": "-54.9306152",
                    "cidades": [
                        {
                            "cidade": "Ananindeua",
                            "vencedores": "1",
                            "latitude": "-1.3650671",
                            "longitude": "-48.3746372"
                        }
                    ]
                }
            ],
            "acumulou": false,
            "acumuladaProxConcurso": "R$ 1,5 Milhão ",
            "dataProxConcurso": "21/10/2022",
            "proxConcurso": 2643
        }
    ]
}
```

* **Specific Result**

[GET]
```https://sortelotofacil.herokuapp.com/api/lotofacil/<concurso>```

Lotofácil, concurso 2642: https://sortelotofacil.herokuapp.com/api/lotofacil/2642

```json
{
    "registros": [
        {
            "data": "20/10/2022",
            "nome": "Lotofácil",
            "local": "ESPAÇODASORTE em SÃOPAULO,SP",
            "dezenas": [
                "02",
                "04",
                "05",
                "08",
                "09",
                "13",
                "14",
                "15",
                "16",
                "17",
                "19",
                "20",
                "21",
                "22",
                "24"
            ],
            "loteria": "lotofacil",
            "acumulou": false,
            "concurso": 2642,
            "premiacoes": [
                {
                    "premio": "505.605,73",
                    "acertos": "15 Pontos",
                    "vencedores": 3
                },
                {
                    "premio": "1.652,16",
                    "acertos": "14 Pontos",
                    "vencedores": 275
                },
                {
                    "premio": "25,00",
                    "acertos": "13 Pontos",
                    "vencedores": 9355
                },
                {
                    "premio": "10,00",
                    "acertos": "12 Pontos",
                    "vencedores": 113258
                },
                {
                    "premio": "5,00",
                    "acertos": "11 Pontos",
                    "vencedores": 620320
                }
            ],
            "estados_premiados": [
                {
                    "uf": "GO",
                    "nome": "Goiás",
                    "cidades": [
                        {
                            "cidade": "Heitoraí",
                            "latitude": "-15.7227332",
                            "longitude": "-49.82713750000001",
                            "vencedores": "1"
                        }
                    ],
                    "latitude": "36.7156122",
                    "longitude": "-95.9435511",
                    "vencedores": "1"
                },
                {
                    "uf": "MG",
                    "nome": "Minas Gerais",
                    "cidades": [
                        {
                            "cidade": "Muzambinho",
                            "latitude": "-21.3712593",
                            "longitude": "-46.5232197",
                            "vencedores": "1"
                        }
                    ],
                    "latitude": "-18.512178",
                    "longitude": "-44.5550308",
                    "vencedores": "1"
                },
                {
                    "uf": "PA",
                    "nome": "Pará",
                    "cidades": [
                        {
                            "cidade": "Ananindeua",
                            "latitude": "-1.3650671",
                            "longitude": "-48.3746372",
                            "vencedores": "1"
                        }
                    ],
                    "latitude": "-1.9981271",
                    "longitude": "-54.9306152",
                    "vencedores": "1"
                }
            ],
            "data_prox_concurso": "21/10/2022",
            "acumulada_prox_concurso": "R$ 1,5 Milhão "
        }
    ]
}
```

## Documentação da API
 
**URL:* https://sortelotofacil.herokuapp.com/api/lotofacil

[GET]
  - /api/lotofacil                      Returns all contests available for search
  - /api/lotofacil/{Concurso}           Returns the result of a specific contest
  - /api/lotofacil/latest               Returns the last contest available for search
  - /api/utilidades/server_active       Returns whether the service is active

Note: The API was developed intelligently where every 24 hours it tries to update the database automatically, so it is not necessary to monitor or manually insert the new draws.

## Documentação para executar o projeto
    * Configure and install Golang packages
        - go mod init apilotofacil 
        - go get
    * Create a Postgresql Database
    * Create table in pattern
    ```sql
      CREATE TABLE LOTERIA (
            NUMERO 							BIGINT  NOT NULL,	
            TIPO_JOGO						TEXT NULL,
            DATA_APURACAO					TEXT NULL,
            NOME_MUNICIPIO_UF_SORTEIO		TEXT NULL,
            DEZENAS_SORTEADAS				JSONB NULL,
            LISTA_RATEIO_PREMIO				JSONB NULL,
            LISTA_MUNICIPIO_GANHADORES		JSONB NULL,
            ACUMULADO						BOOL NULL,
            VALOR_ACUMULADA_PROX_SORTEIO	NUMERIC(12,2),
            DATA_PROX_CONCURSO				TEXT NULL
            ,CONSTRAINT PK_NUMERO PRIMARY KEY (NUMERO)
    );
    ```
## Contribuição

Any contributions to this repository are welcome.

## Invitation to visit my website : <a href="https://api-lotofacil-api-builded.vercel.app/">Sorte Lotofácil</a>
I invite you who are interested in the Api to see the site that was developed with the data generated by the API
![image](https://user-images.githubusercontent.com/54610589/196836579-cb61c750-eb0f-4384-9dea-9741d5c3f3fa.png)
![image](https://user-images.githubusercontent.com/54610589/196836990-9c613264-834b-4ace-ace1-7b15502d6925.png)
