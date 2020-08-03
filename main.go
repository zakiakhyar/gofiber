package main

import {
	"api"
	"github.com/gofiber/fiber"
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