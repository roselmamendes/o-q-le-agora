package artigos

import (
	"gopkg.in/mgo.v2/bson"
)

type Categoria struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Artigos []Artigo `bson:artigos`
	Nome string `bson:nome`
	Prioridade bool `bson:prioridade`
	NaoLidos int
}

type Artigo struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Url string `bson:url`
	Lido bool `bson:lido`
}

var categorias []Categoria

func init() {
	
}

func ObtemCategorias() []Categoria{
	return categorias
}
