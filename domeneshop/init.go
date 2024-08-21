package domeneshop

import (
	"encoding/base64"
	cfg "main/config"
)

var Token string
var BaseURL = "https://api.domeneshop.no/v0"

func Init(cfg *cfg.Config) {

	username := cfg.Token
	password := cfg.Secret

	// Create the Basic Auth string
	auth := username + ":" + password
	Token = base64.StdEncoding.EncodeToString([]byte(auth))

}
