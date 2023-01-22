package fiber

import "github.com/gofiber/fiber/v2"

type FiberServer struct {
	server *fiber.App
}

type ServerConfig struct {
	ListenPost string
}

func New() *FiberServer {
	app := fiber.New()

	f := &FiberServer{
		server: app,
	}

	return f
}

func Start(f *FiberServer, cfg ServerConfig) {
	f.addTodoRoute(f.server)
	f.server.Listen(":" + cfg.ListenPost)
}
