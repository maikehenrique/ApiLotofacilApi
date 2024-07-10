package model

type Server struct {
	Production     bool     `json:"production,omitempty"`
	Port           int      `json:"port,omitempty"`
	HideBanner     bool     `json:"hideBanner,omitempty"`
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
}
