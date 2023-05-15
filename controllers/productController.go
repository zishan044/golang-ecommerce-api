package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-ecommerce-api/initializers"
	"github.com/golang-ecommerce-api/models"
)

func CreateProduct(c *fiber.Ctx) error {
	product := models.Product{}

	if parseErr := c.BodyParser(&product); parseErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "could not parse request body",
		})
		return nil
	}

	if insertErr := initializers.DB.Create(&product).Error; insertErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not create product",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "product successfully created",
		"data": product,
	})
	return nil
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	if queryErr := initializers.DB.Find(&products).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not find products",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "succesfully fetched all products",
		"data": products,
	})
	return nil
}

func GetProductByID(c *fiber.Ctx) error {
	product := models.Product{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Find("id = ?", id).First(&product).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not find product",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "found product succesfully",
		"data": product,
	})
	return nil
}

func UpdateProduct(c *fiber.Ctx) error {
	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	product := models.Product{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if deleteErr := initializers.DB.Delete(product, id).Error; deleteErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not delete product",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "product deleted succesfully",
	})
	return nil
}
