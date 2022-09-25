package utils

import (
	"fmt"
	"log"
)

func CreateServer(port int) {
	app := fiber.New()

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
