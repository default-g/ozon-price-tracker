package http

import "github.com/gofiber/fiber/v3"

var application struct {
	fiber *fiber.App
}

func GetApp() *fiber.App {
	if application.fiber == nil {
		application.fiber = fiber.New()
	}

	return application.fiber
}
