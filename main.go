package main

import (
	api "github.com/vitorpereira/api-template-go/api/rest"
	"github.com/vitorpereira/api-template-go/config"
	_ "github.com/vitorpereira/api-template-go/docs"
)

// @title Todo List API
func main() {
	config.Initialize()
	api.Initialize()
}
