package main

import (
	"estudo/5gin-rest-go/database"
	"estudo/5gin-rest-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
