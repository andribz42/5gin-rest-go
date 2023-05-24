package controllers

import (
	"estudo/5gin-rest-go/models"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(ctx *gin.Context) {
	ctx.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: ": "Tá de marola " + nome + "?",
	})
}
