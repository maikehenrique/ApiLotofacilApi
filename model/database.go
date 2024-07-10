package model

type DBConfig struct {
	Host           string `json:"host,omitempty"`
	Port           uint16 `json:"port,omitempty"`
	Name           string `json:"name,omitempty"`
	Schema         string `json:"schema,omitempty"`
	User           string `json:"user,omitempty"`
	Password       string `json:"password,omitempty"`
	ConnectTimeout int    `json:"connectTimeout,omitempty"`

	SslMode string `json:"sslMode,omitempty"`
}