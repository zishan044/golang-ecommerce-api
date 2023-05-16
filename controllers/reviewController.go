package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-ecommerce-api/initializers"
	"github.com/golang-ecommerce-api/models"
)

func CreateReview(c *fiber.Ctx) error {
	review := models.Review{}

	if parseErr := c.BodyParser(&review); parseErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "could not parse body",
		})
		return nil
	}

	if insertErr := initializers.DB.Create(&review).Error; insertErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not create review",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "review created successfully",
	})
	return nil
}

func GetReviews(c *fiber.Ctx) error {
	reviews := []models.Review{}
	if queryErr := initializers.DB.Find(&reviews).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not fetch all reviews",
		})
		return nil
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "reviews fetched succesfully",
		"data": reviews,
	})
	return nil
}

func GetReviewByID(c *fiber.Ctx) error {
	review := models.Review{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if queryErr := initializers.DB.Find("id = ?", id).First(&review).Error; queryErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not find review",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg":  "review found successfully",
		"data": review,
	})
	return nil
}

func UpdateReview(c *fiber.Ctx) error {
	return nil
}

func DeleteReview(c *fiber.Ctx) error {
	review := models.Review{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "id cannot be empty",
		})
		return nil
	}

	if deleteErr := initializers.DB.Delete(review, id).Error; deleteErr != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "could not delete review",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"msg": "review deleted sucessfully",
	})
	return nil
}
