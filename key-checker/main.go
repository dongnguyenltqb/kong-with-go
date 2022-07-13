package main

import (
	"log"
	"net/http"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

func main() {
	server.StartServer(New, Version, Priority)
}

var Version = "0.2"
var Priority = 1

type Config struct {
	Whitelists []string `json:"whitelists"`
	Message    string   `json:"message"`
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	apiKey, err := kong.Request.GetQueryArg("api-key")
	if err != nil {
		log.Printf("Error reading 'api-key' query string: %s", err.Error())
	}
	whitelists := conf.Whitelists
	found := false
	for _, item := range whitelists {
		if item == apiKey {
			found = true
			break
		}
	}
	if !found {
		kong.Response.Exit(http.StatusForbidden, conf.Message, make(map[string][]string))
	}
}
