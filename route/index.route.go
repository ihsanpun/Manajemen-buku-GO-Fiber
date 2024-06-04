package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihsanpun/go-fiber-part2/config"
	"github.com/ihsanpun/go-fiber-part2/handler"
	"github.com/ihsanpun/go-fiber-part2/middleware"
	"github.com/ihsanpun/go-fiber-part2/utils"
)

func RouteInit(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Aku adalah index")
	})
	r.Post("/login", handler.LoginHandler)

	r.Static("/public", config.ProjectRootPath+"/public/asset")
	r.Get("/user", middleware.Auth, handler.UserHandlerRead)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Delete("/user/:id", handler.UserHandlerDelete)

	r.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)
	r.Get("/book", handler.BookHandlerRead)

	r.Post("/gallery", utils.HandleMultipleFile, handler.PhotoHandlerCreate)
	r.Delete("/gallery/:id", handler.PhotoHandlerDelete)
}
