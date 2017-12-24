package main

import (
	"gopkg.in/macaron.v1"
	"artigos"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Categorias"] = artigos.ObtemPrioridadeDeLeitura()
		ctx.HTML(200, "index") // 200 is the response code.
	})

	m.Get("/todas-categorias", func(ctx *macaron.Context) {
		ctx.Data["Categorias"] = artigos.ObtemCategorias()
		ctx.HTML(200, "todas-categorias") // 200 is the response code.
	})
	
	m.Run()
}
