package main

import (
	"gopkg.in/macaron.v1"
	"service"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Categorias"] = service.ObtemPrioridadeDeLeitura()
		ctx.HTML(200, "index") // 200 is the response code.
	})

	m.Get("/todas-categorias", func(ctx *macaron.Context) {
		ctx.Data["Categorias"] = service.ObtemCategorias()
		ctx.HTML(200, "todas-categorias") // 200 is the response code.
	})

	m.Get("/categoria/:name", func(ctx *macaron.Context) {
		categoriaNome := ctx.Params(":name")
		categoriaParaMostrar := service.ObtemCategoria(categoriaNome)
		ctx.Data["Categoria"] = categoriaParaMostrar
		ctx.HTML(200, "categoria")
	})
	
	m.Run()
}
