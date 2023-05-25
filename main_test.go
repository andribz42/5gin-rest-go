package main

import (
	"estudo/5gin-rest-go/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaSaudacao(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Andy", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	if resposta.Code != http.StatusOK {
		t.Fatalf("Status Error: Valor recebido foi %d e o esperado era %d",
			resposta.Code, http.StatusOK)
	}
}
