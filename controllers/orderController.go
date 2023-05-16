package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-ecommerce-api/initializers"
	"github.com/golang-ecommerce-api/models"
)

func CreateOrder(c *fiber.Ctx) error {
	order := models.Order{}

	if parseErr := c.BodyParser(&order); parseErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "could not parse body",
		})
		return nil
	}

	if insertErr := initializers.DB.Create(&order).Error; insertErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not create order",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "order created successfully",
	})
	return nil
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	if queryErr := initializers.DB.Find(&orders).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not fetch all orders",
		})
		return nil
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "orders fetched succesfully",
		"data": orders,
	})
	return nil
}

func GetOrderByID(c *fiber.Ctx) error {
	order := models.Order{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Find("id = ?", id).First(&order).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not find order",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "order found successfully",
		"data": order,
	})
	return nil
}

func UpdateOrder(c *fiber.Ctx) error {
	return nil
}

func DeleteOrder(c *fiber.Ctx) error {
	order := models.Order{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if deleteErr := initializers.DB.Delete(order, id).Error; deleteErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not delete order",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "order deleted sucessfully",
	})
	return nil
}
