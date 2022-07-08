package routes

import (
	"go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here

	app.Post("/login", controllers.TestLogin)
	app.Get("/dog", controllers.GetDogs)
	app.Post("/dog", controllers.AddDog)
	app.Get("/dog/v2", controllers.GetDog)
	app.Put("/dog/:id", controllers.UpdateDog)
	app.Delete("/dog/:id", controllers.RemoveDog)

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")
	v1.Get("/employee", controllers.GetEmployee)
	v1.Post("/employee", controllers.AddEmployee)
	v1.Get("/employee/v2", controllers.GetEmployee)
	v1.Put("/employee/:id", controllers.UpdateEmployee)
	v1.Delete("/employee/:id", controllers.RemoveEmployee)

}
