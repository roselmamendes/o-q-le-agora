package artigos

type Categoria struct {
	Artigos []Artigo `bson:artigos`
	Nome string `bson:nome`
	Prioridade bool `bson:prioridade`
	NaoLidos int
}

type Artigo struct {
	Url string `bson:url`
	Lido bool `bson:lido`
}

var categorias []Categoria

func init() {
	
}

func ObtemCategorias() []Categoria{
	return categorias
}
