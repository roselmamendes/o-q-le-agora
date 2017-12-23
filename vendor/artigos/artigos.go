package artigos

type Categoria struct {
	Artigos []Artigo
	Nome string
}

type Artigo struct {
	Url string
	Lido bool
}

var Categorias []Categoria

func init() {
	var artigosFemiNegro []Artigo
	var artigo1 = Artigo{Url: "http://algumtexto.com", Lido: false}
	var artigo2 = Artigo{Url: "http://algumtexto.com", Lido: false}
	artigosFemiNegro = append(artigosFemiNegro, artigo1)
	artigosFemiNegro = append(artigosFemiNegro, artigo2)

	var feminismoNegro = Categoria{Artigos: artigosFemiNegro, Nome: "Feminismo Negro"}
	Categorias = append(Categorias, feminismoNegro)
}
