package main

import "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	
	// m.Get("/", func(ctx *macaron.Context) {
	// 	ctx.Data["Name"] = "roselma"
	// 	ctx.HTML(200, "index") // 200 is the response code.
	// })

	// m.Get("/data-visualization", func(ctx *macaron.Context) {
	// 	ctx.HTML(200, "data-visualization") // 200 is the response code.
	// })

	// m.Get("/login", func(ctx *macaron.Context) {
	// 	ctx.HTML(200, "login") // 200 is the response code.
	// })
	m.Run()
}
