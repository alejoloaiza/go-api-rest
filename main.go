package main

import (
	"go-rest-api/extra"
	"go-rest-api/restserver"
	"os"
)

func main() {
	arg := "./extra/config.json"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	_ = extra.GetConfig(arg)
	extra.DBConnectPostgres()
	restserver.CreateRestServer()
}
