package configs

import (
	"os"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	Api APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port            string
	ReadTimeoutMin  int
	WriteTimeoutMin int
	TokenApiLoteria string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	SslMode  string
}

//Put default value if there is no configuration file
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("api.read-timeout-min", 120)
	viper.SetDefault("api.write-timeout-min", 120)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")

	viper.SetDefault("database.user", "5432")
	viper.SetDefault("database.ssl-mode", "disable")
}

//Load configuration file, if there is an environment variable called production, the ENV variables will be used
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.Api = APIConfig{
		Port:            viper.GetString("api.port"),
		ReadTimeoutMin:  viper.GetInt("api.read-timeout-min"),
		WriteTimeoutMin: viper.GetInt("api.write-timeout-min"),
		TokenApiLoteria: viper.GetString("api.token-api-loteria"),
	}

	sslMode := ""
	host := ""
	port := ""
	user := ""
	pass := ""
	dbName := ""

	if os.Getenv("PRODUCTION") != "" {
		sslMode = "require"
		host = os.Getenv("DATABASE.HOST")
		port = os.Getenv("DATABASE.PORT")
		user = os.Getenv("DATABASE.USER")
		pass = os.Getenv("DATABASE.PASS")
		dbName = os.Getenv("DATABASE.NAME")

	} else {
		host = viper.GetString("database.host")
		port = viper.GetString("database.port")
		user = viper.GetString("database.user")
		pass = viper.GetString("database.pass")
		dbName = viper.GetString("database.name")
		sslMode = viper.GetString("database.ssl-mode")
	}

	cfg.DB = DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Pass:     pass,
		Database: dbName,
		SslMode:  sslMode,
	}
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.Api.Port
}

func GetServerReadTimeoutMin() int {
	return cfg.Api.ReadTimeoutMin
}

func GetServerWriteTimeoutMin() int {
	return cfg.Api.WriteTimeoutMin
}

func GetTokenApiLoteria() string {
	return cfg.Api.TokenApiLoteria
}
