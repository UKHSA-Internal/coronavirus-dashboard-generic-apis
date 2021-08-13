package main

import (
	"generic_apis/api"
	"generic_apis/base"
)

func main() {

	service := &base.Api{}
	api.Run(service)

} // main
