package main

import (
	"fmt"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

func main() {
	server.StartServer(New, Version, Priority)
}

var Version = "0.2"
var Priority = 800

type Config struct {
	Prefix string `json:"prefix"`
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	kong.Log.Info(fmt.Sprintf("<path-prefix> <%s>", conf.Prefix))
}
