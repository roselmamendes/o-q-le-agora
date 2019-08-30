package main

import (
	"gopkg.in/macaron.v1"
	"service"
	"fmt"
	"os"
	"bancoDeDados"
)
var CONNECTIONSTRING string

func init(){
	CONNECTIONSTRING = os.Getenv("CONNECTIONSTRING")
}

func main() {
	bancoDeDados.ConfiguraBanco(CONNECTIONSTRING)

	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		categorias, erro := service.ObtemPrioridadeDeLeitura()
		if erro != "" {
			fmt.Println("Erro no metodo ObtemPrioridadeDeLeitura", erro)
		} else {
			ctx.Data["Categorias"] = categorias
		}
		ctx.HTML(200, "index") // 200 is the response code.
	})

	m.Get("/todas-categorias", func(ctx *macaron.Context) {
		categorias , erro := service.ObtemCategorias()
		if erro != "" {
			fmt.Println("Erro no metodo ObtemCategorias", erro)
		}
		ctx.Data["Categorias"] = categorias
		ctx.HTML(200, "todas-categorias") // 200 is the response code.
	})

	m.Get("/categoria/:name", func(ctx *macaron.Context) {
		categoriaNome := ctx.Params(":name")
		categoriaParaMostrar, erro := service.ObtemCategoria(categoriaNome)
		if erro != "" {
			fmt.Println("Erro no metodo ObtemCategoria", erro)
		}
		ctx.Data["Categoria"] = categoriaParaMostrar
		ctx.HTML(200, "categoria")
	})
	
	m.Run()
}
