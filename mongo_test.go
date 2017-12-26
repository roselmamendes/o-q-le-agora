package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bancoDeDados"
	"gopkg.in/mgo.v2/bson"
	"artigos"
)

func TestObterCategorias(t *testing.T){
	err := bancoDeDados.AdicionaCategoria(artigos.Categoria{Nome:"Arquitetura", Prioridade:false})
	if err != "" {
		t.Fatalf(err)
	}
	categorias, error := bancoDeDados.ObtemCategorias(bson.M{})
	if error != "" {
		t.Fatalf("Error: %s", error)
	}
	assert.Equal(t, 1, len(categorias))

	err = bancoDeDados.LimpaBanco()
	if err != "" {
		t.Fatalf(err)
	}
}
