package router

import (
	"crud-restapi/controllers"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(App *fiber.App, db *gorm.DB) {
	userController := controllers.UserController{
		Database: db,
	}
	taskManagerController := controllers.TaskManagerController{
		Database: db,
	}

	App.Route("/api/user", func(router fiber.Router) {
		router.Post("/register", userController.Register)
		router.Post("/sign-in", userController.SignIn)
	})
	
	//Below this middleware, each router has a protection
	App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRECTKEY"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "No Authentication",
			})
		},
	}))

	App.Route("/api/task", func(router fiber.Router) {
		router.Get("/", taskManagerController.Get)
		router.Post("/", taskManagerController.Post)
		router.Delete("/:id", taskManagerController.Delete)
		router.Put("/:id", taskManagerController.Put)
	})
}
