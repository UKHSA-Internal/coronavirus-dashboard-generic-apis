package main

import (
	"generic_apis/api"
)

func main() {

	service := &api.Api{}
	service.Run(":5100")

} // main
