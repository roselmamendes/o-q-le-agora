package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"artigos"
)

func TestObtemCategorias(t *testing.T) {
	artigoCriado := artigos.Artigo{Url: "http://algumtexto.com", Lido: false}
	artigoCriado2 := artigos.Artigo{Url: "http://algumtexto.com", Lido: false}
	var artigosCriados []artigos.Artigo
	artigosCriados = append(artigosCriados, artigoCriado)
	artigosCriados = append(artigosCriados, artigoCriado2)
	expectedCategoria := artigos.Categoria{Artigos: artigosCriados, Nome: "Feminismo Negro"}
	
	categorias := artigos.ObtemCategorias()
	assert.Equal(t, 3, len(categorias))
	assert.Equal(t, expectedCategoria, categorias[0])
}

func TestObtemCategoriasPrioritarias(t *testing.T) {
	expectedCategoria := artigos.Categoria{Artigos: []artigos.Artigo{{Url: "url2", Lido: false}}, Nome: "Prioridade", Prioridade: true, NaoLidos: 1}

	categorias := artigos.ObtemPrioridadeDeLeitura()
	assert.Equal(t, 2, len(categorias))
	assert.Equal(t, expectedCategoria, categorias[0])
	assert.Equal(t, 0, categorias[1].NaoLidos)
}

func TestPossivelContarArtigosNaoLidos(t *testing.T) {
	artigosNaoLidos := artigos.ContaArtigosNaoLidosDe("Feminismo Negro")
	assert.Equal(t, 2, artigosNaoLidos)
}
