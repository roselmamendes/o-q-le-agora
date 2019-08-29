package service

import (
	"bancoDeDados"
	"gopkg.in/mgo.v2/bson"
	"artigos"
)

func ObtemCategorias() ([]artigos.Categoria, string) {
	return bancoDeDados.ObtemCategorias(bson.M{})
}

func ObtemPrioridadeDeLeitura() ([]artigos.Categoria, string){
	var prioridades []artigos.Categoria
	categorias, erro := ObtemCategorias()
	for _, categoria := range categorias {
		if categoria.Prioridade {
			categoria.NaoLidos = ContaArtigosNaoLidosDe(categoria)
			prioridades = append(prioridades, categoria)
		}
	}
	return prioridades, erro
}

func ContaArtigosNaoLidosDe(categoria artigos.Categoria) int {
	contador := 0
	
	for _, artigo := range categoria.Artigos {
		if !artigo.Lido {
			contador += 1
		}
	}
	return contador
}

func ObtemCategoria(nomeCategoria string) (artigos.Categoria, string) {
	categorias, erro := bancoDeDados.ObtemCategorias(bson.M{"name": nomeCategoria})
	return categorias[0], erro
}

func AdicionarCategoria(categoria artigos.Categoria) string {
	erro := bancoDeDados.AdicionarCategoria(categoria)
	return erro
}


