package main

import {
	"github.com/zakiakhyar/gofiber/api"
	"github.com/gofiber/fiber"
	"github.com/valyala/fasthttp v1.34.0"
}

func beranda(c *fiber.Ctx){
	c.Send("Beranda")
}

func setupRoutes(app *fiber.App){
	app.Static("/", "./public")

	app.Get("/api/ambilsemua/data/", api.GetDatas)
	app.Get("/api/ambilsatu/data/", api.GetData)
	app.Get("/api/tambah/data/", api.NewData)
	app.Get("/api/hapus/data/", api.DeleteData)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}
