package api

import "github.com/gofiber/fiber"

func GetDatas(c *fiber.Ctx) {
	c.Send("Semua data")
}

func GetData(c *fiber.Ctx) {
	c.Send("Satu data")
}

func NewData(c *fiber.Ctx) {
	c.Send("Buat data")
}

func DeleteData(c *fiber.Ctx) {
	c.Send("Hapus data")
}

func main() {

}
