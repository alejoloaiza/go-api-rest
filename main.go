package main

import (
	"go-rest-api/extra"
	"go-rest-api/restserver"
)

func main() {
	extra.DBConnectPostgres()
	restserver.CreateRestServer()
}
