package timetable

import (
	"attendance-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, service Service) {
	timetableGroup := router.Group("/timetable")

	// Student Route: Get My Timetable
	timetableGroup.Get("/student", middleware.RoleMiddleware("student"), func(c *fiber.Ctx) error {
		studentID := c.Locals("userID").(string)

		slots, err := service.GetStudentTimetable(c.Context(), studentID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch student timetable"})
		}

		return c.JSON(fiber.Map{
			"student_id": studentID,
			"timetable":  slots,
		})
	})

	// Faculty Route: Get My Timetable
	timetableGroup.Get("/faculty", middleware.RoleMiddleware("faculty"), func(c *fiber.Ctx) error {
		facultyID := c.Locals("userID").(string)

		slots, err := service.GetFacultyTimetable(c.Context(), facultyID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch faculty timetable"})
		}

		return c.JSON(fiber.Map{
			"faculty_id": facultyID,
			"timetable":  slots,
		})
	})
}
