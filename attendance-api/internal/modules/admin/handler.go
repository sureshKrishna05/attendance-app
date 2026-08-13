package admin

import (
	"attendance-api/internal/config"
	"attendance-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, cfg *config.Config, svc Service) {
	adminGroup := router.Group("/admin", middleware.AdminMiddleware(cfg))

	adminGroup.Post("/users/bulk", func(c *fiber.Ctx) error {
		var users []BulkUser
		if err := c.BodyParser(&users); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		if err := svc.BulkInsertUsers(c.Context(), users); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Users inserted successfully"})
	})

	adminGroup.Post("/classes", func(c *fiber.Ctx) error {
		var class Class
		if err := c.BodyParser(&class); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		if err := svc.CreateClass(c.Context(), &class); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Class created successfully"})
	})

	adminGroup.Post("/timetable", func(c *fiber.Ctx) error {
		var slot TimetableSlot
		if err := c.BodyParser(&slot); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		if err := svc.InsertTimetableSlot(c.Context(), &slot); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Timetable slot created successfully"})
	})

	adminGroup.Post("/assignments", func(c *fiber.Ctx) error {
		var assignment FacultyAssignment
		if err := c.BodyParser(&assignment); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		if err := svc.CreateFacultyAssignment(c.Context(), &assignment); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Faculty assignment created successfully"})
	})
}
