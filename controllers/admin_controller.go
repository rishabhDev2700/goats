package controllers

import (
	"rishabhdev2700/goats/models"

	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	category := models.Category{}
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON data"})
	}
	err := models.DB.Create(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to store data in the database",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Category created successfully"})

}


func CreateArticle(c *fiber.Ctx) error {
    article := models.Article{}
    if err := c.BodyParser(&article); err !=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message":"Error while parsing json data",
            "error":err,
        })
    }
	err:=models.DB.Create(&article)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Unable to save article in the database",
			"error":err,
		})
	}
    return c.JSON(fiber.Map{"message":"Article saved to database successfully"})
}

func CreateTag(c *fiber.Ctx)error{
	tag := models.Tag{}
	if err := c.BodyParser(&tag);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message":"Error while parsing json data",
            "error":err,
        })
	}
	if err:=models.DB.Create(&tag);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message":"Unable to store tag in the database",
            "error":err,
        })
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Tag created successfully",
	})
}