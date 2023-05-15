package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-ecommerce-api/initializers"
	"github.com/golang-ecommerce-api/models"
)

func CreateCategory(c *fiber.Ctx) error {
	category := models.Category{}

	if parseErr := c.BodyParser(&category); parseErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "could not parse body",
		})
		return nil
	}

	if insertErr := initializers.DB.Create(&category).Error; insertErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not create category",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "category created succesfully",
		"data": category,
	})
	return nil
}

func GetCategories(c *fiber.Ctx) error {
	categories := []models.Category{}

	if queryErr := initializers.DB.Find(&categories).Error; queryErr != nil {
		c.Status(http.StatusOK).JSON(&fiber.Map{
			"error": "could not fetch categories",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "categories fetched succesfully",
		"data": categories,
	})
	return nil
}

func GetCategoryByID(c *fiber.Ctx) error {
	category := models.Category{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Find("id = ?", id).First(&category).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not fetch category",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "category found succesfully",
		"data": category,
	})
	return nil
}

func UpdateCategory(c *fiber.Ctx) error {
	return nil
}

func DeleteCategory(c *fiber.Ctx) error {
	category := models.Category{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Delete(category, id).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not delete category",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "category deleted succesfully",
	})
	return nil
}
