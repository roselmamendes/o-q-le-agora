package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"artigos"
	"service"
	"bancoDeDados"
)

func TestObtemCategorias(t *testing.T) {
	bancoDeDados.LimpaBanco()
	expectedCategoria := artigos.Categoria{Nome: "Feminismo Negro"}
	service.AdicionarCategoria(expectedCategoria)

	categorias, erro := service.ObtemCategorias()
	mostraErro(t, erro, "Erro no metodo ObtemCategorias")
	
	assert.Equal(t, 1, len(categorias))
	assert.Equal(t, expectedCategoria, categorias[0])
}

// func TestObtemCategoriasPrioritarias(t *testing.T) {
// 	expectedCategoria := artigos.Categoria{Artigos: []artigos.Artigo{{Url: "url2", Lido: false}}, Nome: "Prioridade", Prioridade: true, NaoLidos: 1}

// 	categorias, erro := service.ObtemPrioridadeDeLeitura()
// 	mostraErro(t, erro, "Erro no metodo ObtemPrioridadeDeLeitura")
// 	assert.Equal(t, 2, len(categorias))
// 	assert.Equal(t, expectedCategoria.Nome, categorias[0].Nome)
// 	assert.Equal(t, 0, categorias[1].NaoLidos)
// }

// func TestPossivelContarArtigosNaoLidos(t *testing.T) {
// 	categoria := artigos.Categoria{Nome: "Feminismo Negro"}
// 	artigosNaoLidos := service.ContaArtigosNaoLidosDe(categoria)
// 	assert.Equal(t, 2, artigosNaoLidos)
// }


func mostraErro(t *testing.T, erro string, mensagem string) {
	if erro != "" {
		t.Fatalf(mensagem, erro)
	}
}

