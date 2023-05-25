package routes

import (
	"estudo/5gin-rest-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.PATCH("/alunos/:id", controllers.AlteraAluno)
	r.DELETE("/alunos/:id", controllers.ExcluiAlunoPorID)
	r.Run()
}
