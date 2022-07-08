package main

import (
	"fmt"
	"go-fiber-test/controllers"
	"go-fiber-test/database"
	"go-fiber-test/routes"
	"log"

	m "go-fiber-test/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)

	var err error

	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	// log.Println(db)

	database.DBConn.AutoMigrate(&m.Dogs{})
	database.DBConn.AutoMigrate(&m.Employee{})
	fmt.Println("Migrated DB")
}

func main() {
	app := fiber.New()
	routes.UserRoute(app)
	initDatabase()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// type Person struct {
	// 	Name string `json:"name"`
	// 	Pass string `json:"pass"`
	// }

	// app.Post("/:num", func(c *fiber.Ctx) error {
	// 	p := new(Person)

	// 	if err := c.BodyParser(p); err != nil {
	// 		return err
	// 	}

	// 	log.Println(p.Name) // john
	// 	log.Println(p.Pass) // doe

	// 	return c.SendString("Login Success")

	// })

	// app.Get("/user/:name", func(c *fiber.Ctx) error {
	// 	name := c.Params("name") // "fenny"

	// 	log.Println(name)
	// 	return c.Status(200).JSON(name)
	// 	// ...
	// })

	// app.Get("/number/:num", func(c *fiber.Ctx) error {
	// 	num, err := c.ParamsInt("num") // "fenny"
	// 	//num := strconv.Atoi(number)
	// 	sum := 1
	// 	for i := num; i >= 1; i-- {
	// 		sum = sum * i
	// 	}
	// 	if err != nil {
	// 		return c.SendFile("file-does-not-exist")
	// 	}
	// 	log.Println(sum)
	// 	return c.Status(200).JSON(sum)
	// 	// ...
	// })

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1
	v1.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "772565",
		},
	}))
	v1.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("api/v1/list")
	}) // /api/v1/list
	v1.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("api/v1/user")
	})

	//------------------------------------------
	v1.Get("/employee", controllers.GetEmployee)
	v1.Post("/employee", controllers.AddEmployee)
	v1.Get("/employee/v2", controllers.GetEmployee)
	v1.Put("/employee/:id", controllers.UpdateEmployee)
	v1.Delete("/employee/:id", controllers.RemoveEmployee)
	v2 := api.Group("/v2")
	v2.Get("/employee", controllers.GetEmployee)
	// /api/v1/user
	// v1.Get("/employee", func(c *fiber.Ctx) error {

	// 	type Employee struct {
	// 		gorm.Model
	// 		EmployeeID int       `json:"employee_id"`
	// 		Name       string    `json:"name"`
	// 		LName      string    `json:"lname"`
	// 		Birthday   time.Time `json:"birthday"`
	// 		Age        int       `json:"age"`
	// 		Email      string    `json:"email"`
	// 		Tel        string    `json:"tel"`
	// 	}

	// 	employee := new(Employee)

	// 	if err := c.BodyParser(&employee); err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"message": err.Error(),
	// 		})
	// 	}

	// 	// 	//Do something else here

	// 	// 	//Return user
	// 	return c.JSON(employee)

	// }) // /api/v1/user

	log.Fatal(app.Listen(":3000"))
	app.Listen(":3000")
}
