package main

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"id":   "1",
		"Nome": "Manuel",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos)
	r.Run()
}
