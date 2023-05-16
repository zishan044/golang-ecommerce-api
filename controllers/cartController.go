package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-ecommerce-api/initializers"
	"github.com/golang-ecommerce-api/models"
)

func CreateCart(c *fiber.Ctx) error {
	cart := models.Cart{}

	if parseErr := c.BodyParser(&cart); parseErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "could not parse body",
		})
		return nil
	}

	if insertErr := initializers.DB.Create(&cart).Error; insertErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not create cart",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "cart created successfully",
	})
	return nil
}

func GetCarts(c *fiber.Ctx) error {
	carts := []models.Cart{}
	if queryErr := initializers.DB.Find(&carts).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not fetch all carts",
		})
		return nil
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "carts fetched succesfully",
		"data": carts,
	})
	return nil
}

func GetCartByID(c *fiber.Ctx) error {
	cart := models.Cart{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Find("id = ?", id).First(&cart).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not find cart",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "cart found successfully",
		"data": cart,
	})
	return nil
}

func UpdateCart(c *fiber.Ctx) error {
	return nil
}

func DeleteCart(c *fiber.Ctx) error {
	cart := models.Cart{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if deleteErr := initializers.DB.Delete(cart, id).Error; deleteErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not delete cart",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "cart deleted sucessfully",
	})
	return nil
}
