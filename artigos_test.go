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
	assert.Equal(t, expectedCategoria, categorias[0])
}

func TestObtemCategoriasPrioritarias(t *testing.T) {
}
