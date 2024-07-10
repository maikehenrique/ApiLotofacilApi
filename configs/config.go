package configs

import (
	"apilotofacil/model"
	"encoding/json"
	"io/ioutil"
)

var Conf *config

type config struct {
	Server   model.Server
	Database model.DBConfig
	content  []byte
}

type APIConfig struct {
	Port            string
	ReadTimeoutMin  int
	WriteTimeoutMin int
	TokenApiLoteria string
}

// Load configuration file, if there is an environment variable called production, the ENV variables will be used
func LoadFromJSON(fileName string) (err error) {
	Conf = new(config)

	if Conf.content, err = ioutil.ReadFile(fileName); err != nil {
		return err
	}

	if err = json.Unmarshal(Conf.content, &Conf); err != nil {
		return err
	}

	return nil
}

func Bytes() []byte {
	return Conf.content
}

func GetDB() model.DBConfig {
	return Conf.Database
}
