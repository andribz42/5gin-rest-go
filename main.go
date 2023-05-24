package main

import (
	"estudo/5gin-rest-go/models"
	"estudo/5gin-rest-go/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "A", CPF: "9518462951", RG: "123456798"},
		{Nome: "B", CPF: "7418259632", RG: "321654987"},
		{Nome: "C", CPF: "7894651322", RG: "564987321"},
	}
	routes.HandleRequests()
}
