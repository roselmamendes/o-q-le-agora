package artigos

import "gopkg.in/mgo.v2/bson"

type Categoria struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Artigos []Artigo `bson:artigos`
	Nome string `bson:nome`
	Prioridade bool `bson:prioridade`
	NaoLidos int
}

type Artigo struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Url string `bson:url`
	Lido bool `bson:lido`
}

var categorias []Categoria

func init() {
	var artigosFemiNegro []Artigo
	var artigo1 = Artigo{Url: "http://algumtexto.com", Lido: false}
	var artigo2 = Artigo{Url: "http://algumtexto.com", Lido: false}
	artigosFemiNegro = append(artigosFemiNegro, artigo1)
	artigosFemiNegro = append(artigosFemiNegro, artigo2)

	var feminismoNegro = Categoria{Artigos: artigosFemiNegro, Nome: "Feminismo Negro", Prioridade: false, NaoLidos: 0}
	var prioridade = Categoria{Artigos:[]Artigo{{Url: "url2", Lido: false}}, Nome:"Prioridade", Prioridade: true, NaoLidos: 0}
	var prioridade2 = Categoria{Artigos: []Artigo{{Url: "url1", Lido: true}}, Nome:"Prioridade2", Prioridade: true, NaoLidos: 0}
	categorias = append(categorias, feminismoNegro)
	categorias = append(categorias, prioridade)
	categorias = append(categorias, prioridade2)
}

func ObtemCategorias() []Categoria{
	return categorias
}
