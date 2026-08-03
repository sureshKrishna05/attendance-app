package classes

import (
	"attendance-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, service Service) {
	classesGroup := router.Group("/classes")

	// Faculty Route: Get Assigned Classes
	classesGroup.Get("/faculty", middleware.RoleMiddleware("faculty"), func(c *fiber.Ctx) error {
		facultyID := c.Locals("userID").(string)

		assignments, err := service.GetFacultyAssignments(c.Context(), facultyID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch assigned classes"})
		}

		return c.JSON(fiber.Map{
			"faculty_id":  facultyID,
			"assignments": assignments,
		})
	})
}
