package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-ecommerce-api/initializers"
	"github.com/golang-ecommerce-api/models"
)

func SignUp(c *fiber.Ctx) error {
	customer := models.Customer{}

	parseErr := c.BodyParser(&customer)
	if parseErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "could not parse request body",
		})
		return nil
	}

	if insertErr := initializers.DB.Create(&customer).Error; insertErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not create cutomer",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "customer successfully created",
		"data": customer,
	})
	return nil
}

func Login(c *fiber.Ctx) error {
	return nil
}

func GetCustomers(c *fiber.Ctx) error {
	customers := []models.Customer{}

	if queryErr := initializers.DB.Find(&customers).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not fetch all customers",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "succesfully fetched all customers",
		"data": customers,
	})
	return nil
}

func GetCustomerByID(c *fiber.Ctx) error {
	customer := models.Customer{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Find("id = ?", id).First(&customer).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not find customer",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "succesfully found customer",
		"data": customer,
	})
	return nil
}

func UpdateCustomer(c *fiber.Ctx) error {
	return nil
}

func DeleteCustomer(c *fiber.Ctx) error {
	customer := models.Customer{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if deleteErr := initializers.DB.Delete(customer, id).Error; deleteErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "failed to delete customer",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "customer deleted succesfully",
	})
	return nil
}
