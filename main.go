package main

import (
	"example/m/fiber"
	"fmt"
)

func main() {
	fmt.Println("Main is running!")

	serv := fiber.New()
	fiber.Start(serv, fiber.ServerConfig{ListenPost: "8000"})
}
