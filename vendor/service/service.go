package service

import (
	"artigos"
)

func ObtemCategorias() []artigos.Categoria{
	return artigos.ObtemCategorias()
}

func ObtemPrioridadeDeLeitura() []artigos.Categoria{
	var prioridades []artigos.Categoria
	categorias := ObtemCategorias()
	for _, categoria := range categorias {
		if categoria.Prioridade {
			categoria.NaoLidos = ContaArtigosNaoLidosDe(categoria.Nome)
			prioridades = append(prioridades, categoria)
		}
	}
	return prioridades
}

func ContaArtigosNaoLidosDe(nomeCategoria string) int {
	categoria := ObtemCategoria(nomeCategoria)
	var contador int
	
	for _, artigo := range categoria.Artigos {
		if !artigo.Lido {
			contador += 1
		}
	}
	return contador
}

func ObtemCategoria(nomeCategoria string) (categoriaProcurada artigos.Categoria) {
	categorias := ObtemCategorias()
	for _, categoria := range categorias {
		if categoria.Nome == nomeCategoria {
			categoriaProcurada = categoria
		}
	}
	return categoriaProcurada
}


