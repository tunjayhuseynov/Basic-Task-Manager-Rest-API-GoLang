package controllers

import (
	"crud-restapi/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TaskManagerController struct {
	Database *gorm.DB
}

func (t *TaskManagerController) Get(c *fiber.Ctx) error {
	var task []models.Task
	result := t.Database.Find(&task)
	if result.Error != nil {
		log.Fatalf("Error querying users: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal error")
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func (t *TaskManagerController) Post(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	result := t.Database.Create(&task)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create task",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (t *TaskManagerController) Put(c *fiber.Ctx) error {
	id := c.Params("id")

	var task models.Task
	if err := t.Database.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	var updatedTask map[string]interface{}
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if name, exists := updatedTask["name"]; exists {
		task.Name = name.(string)
	}

	if priority, exists := updatedTask["priority"]; exists {
		task.Priority = uint(priority.(float64))
	}

	if err := t.Database.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update task",
		})
	}

	// Return the updated task with a success status
	return c.Status(fiber.StatusOK).JSON(task)
}

func (t *TaskManagerController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var task models.Task
	if err := t.Database.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	if err := t.Database.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete task",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
