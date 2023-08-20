package main

import (
	api "github.com/vitorpereira/api-template-go/api/rest"
	"github.com/vitorpereira/api-template-go/config"
)

func main() {
	config.Initialize()
	api.Initialize()
}
