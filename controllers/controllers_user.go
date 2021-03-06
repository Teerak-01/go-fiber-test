package controllers

import (
	"log"
	"strings"

	"go-fiber-test/database"
	m "go-fiber-test/models"

	"github.com/gofiber/fiber/v2"
)

func TestLogin(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	return c.SendString("Login Success")
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

//------------------------------------------------------

func GetEmployees(c *fiber.Ctx) error {
	db := database.DBConn
	var employee []m.Employee

	db.Find(&employee)
	return c.Status(200).JSON(employee)
}

func GetEmployee(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var employee []m.Employee

	result := db.Find(&employee, "employee_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&employee)
}

func AddEmployee(c *fiber.Ctx) error {
	db := database.DBConn
	var employee m.Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&employee)
	return c.Status(201).JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	db := database.DBConn
	var employee m.Employee
	id := c.Params("id")

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&employee)
	return c.Status(200).JSON(employee)
}

func RemoveEmployee(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var employee m.Employee

	result := db.Delete(&employee, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
