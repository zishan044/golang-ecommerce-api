package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-ecommerce-api/initializers"
	"github.com/golang-ecommerce-api/models"
)

func CreateSupplier(c *fiber.Ctx) error {
	supplier := models.Supplier{}

	if parseErr := c.BodyParser(&supplier); parseErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "could not parse body",
		})
		return nil
	}

	if insertErr := initializers.DB.Create(&supplier).Error; insertErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not create supplier",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "supplier created succesfully",
		"data": supplier,
	})
	return nil
}

func GetSuppliers(c *fiber.Ctx) error {
	suppliers := []models.Supplier{}

	if queryErr := initializers.DB.Find(&suppliers).Error; queryErr != nil {
		c.Status(http.StatusOK).JSON(&fiber.Map{
			"error": "could not fetch suppliers",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "suppliers fetched succesfully",
		"data": suppliers,
	})
	return nil
}

func GetSupplierByID(c *fiber.Ctx) error {
	supplier := models.Supplier{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Find("id = ?", id).First(&supplier).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not fetch supplier",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "supplier found succesfully",
		"data": supplier,
	})
	return nil
}

func UpdateSupplier(c *fiber.Ctx) error {
	return nil
}

func DeleteSupplier(c *fiber.Ctx) error {
	supplier := models.Supplier{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Delete(supplier, id).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not delete supplier",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "supplier deleted succesfully",
	})
	return nil
}
