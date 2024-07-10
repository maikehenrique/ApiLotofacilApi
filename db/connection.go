package db

import (
	"apilotofacil/configs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Open database connection
func OpenConnection() (*sql.DB, error) {
	var nameMethod = "|OpenConnection"
	sc := ""
	conf := configs.GetDB()

	sc = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Name, conf.SslMode)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
		panic(err)
	}

	err = conn.Ping()
	return conn, err
}
