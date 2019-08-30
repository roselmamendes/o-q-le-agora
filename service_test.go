package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"artigos"
	"service"
	"bancoDeDados"
)

var connectionString = os.Getenv("CONNECTIONSTRING")

func TestObtemCategorias(t *testing.T) {
	bancoDeDados.ConfiguraBanco(connectionString)
	bancoDeDados.LimpaBanco()
	expectedCategoria := artigos.Categoria{Nome: "Feminismo Negro"}
	service.AdicionarCategoria(expectedCategoria)

	categorias, erro := service.ObtemCategorias()
	mostraErro(t, erro, "Erro no metodo ObtemCategorias")
	
	assert.Equal(t, 1, len(categorias))
	assert.Equal(t, expectedCategoria, categorias[0])
}

func TestObtemCategoriasPrioritarias(t *testing.T) {
	bancoDeDados.ConfiguraBanco(connectionString)
	bancoDeDados.LimpaBanco()
	expectedCategoria := artigos.Categoria{
		Artigos: []artigos.Artigo{{Url: "url2", Lido: false}}, 
		Nome: "Prioridade", 
		Prioridade: true, 
		NaoLidos: 3}
	service.AdicionarCategoria(expectedCategoria)
	
	categorias, erro := service.ObtemPrioridadeDeLeitura()

	mostraErro(t, erro, "Erro no metodo ObtemPrioridadeDeLeitura")
	assert.Equal(t, 1, len(categorias))
	assert.Equal(t, expectedCategoria.Nome, categorias[0].Nome)
	assert.Equal(t, 1, categorias[0].NaoLidos)
}

func TestPossivelContarArtigosNaoLidos(t *testing.T) {
	categoria := artigos.Categoria{
		Artigos: []artigos.Artigo{
			{Url: "url2", Lido: false},
			{Url: "url3", Lido: false},
		},
		Nome: "Feminismo Negro"}

	artigosNaoLidos := service.ContaArtigosNaoLidosDe(categoria)
	assert.Equal(t, 2, artigosNaoLidos)
}


func mostraErro(t *testing.T, erro string, mensagem string) {
	if erro != "" {
		t.Fatalf(mensagem, erro)
	}
}

