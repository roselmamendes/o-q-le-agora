package main

// import (
// 	"testing"
// 	"github.com/stretchr/testify/assert"
// 	"bancoDeDados"
// 	"gopkg.in/mgo.v2/bson"
// 	"artigos"
// )

// func TestObterCategorias(t *testing.T){
	// err := bancoDeDados.AdicionaCategoria(artigos.Categoria{Nome:"Arquitetura", Prioridade:false})
	// mostraErro(t, err, "Erro no metodo AdicionaCategoria")

// 	categorias, error := bancoDeDados.ObtemCategorias(bson.M{})
// 	mostraErro(t, error, "Erro no metodo ObtemCategorias")

// 	assert.Equal(t, 3, len(categorias))
// 	assert.Equal(t, "Feminismo Negro", categorias[0].Nome)
// 	assert.Equal(t, []artigos.Artigo{}, categorias[0].Artigos)

// 	limpaBanco(t)
// }

// func TestObterCategoriasQuandoNaoHaCategoria(t *testing.T) {
// 	categorias, error := bancoDeDados.ObtemCategorias(bson.M{})
// 	assert.Equal(t, "", error)
// 	assert.Equal(t, 0, len(categorias))
// }

// func TestObterArtigos(t *testing.T) {
// 	err0 := bancoDeDados.AdicionaCategoria(artigos.Categoria{Nome:"Arquitetura", Prioridade:false})
// 	mostraErro(t, err0, "Erro no método AdicionaCategoria:")

// 	categoria := "Arquitetura"
// 	artigo := artigos.Artigo{Url: "http://teste.com", Lido: false}
	
// 	err1 := bancoDeDados.AdicionaArtigo(categoria, artigo)
// 	mostraErro(t, err1, "Erro no metodo AdicionaArtigo:")
	
// 	categorias, err2 := bancoDeDados.ObtemCategorias(bson.M{"nome": categoria})
// 	mostraErro(t, err2, "Erro no metodo ObtemCategorias:")

// 	assert.Equal(t, 1, len(categorias[0].Artigos))
// 	assert.Equal(t, "http://teste.com", categorias[0].Artigos[0].Url)

// 	limpaBanco(t)
// }

// func TestObterArtigosQuandoNaoHaCategoria(t *testing.T) {
// 	err := bancoDeDados.AdicionaArtigo("Não existe", artigos.Artigo{})
// 	assert.Equal(t, "AdicionaArtigo - Categoria Não existe não encontrada", err)
// }

// func limpaBanco(t *testing.T) {
// 	erro := bancoDeDados.LimpaBanco()
// 	mostraErro(t, erro, "Erro no método LimpaBanco:")
// }

// func mostraErro(t *testing.T, erro string, mensagem string) {
// 	if erro != "" {
// 		t.Fatalf(mensagem, erro)
// 	}
// }
