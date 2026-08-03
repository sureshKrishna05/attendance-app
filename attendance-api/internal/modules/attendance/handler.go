package attendance

import (
	"attendance-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, service Service) {
	attendanceGroup := router.Group("/attendance")

	// Faculty Route: Post Attendance
	attendanceGroup.Post("/", middleware.RoleMiddleware("faculty"), func(c *fiber.Ctx) error {
		var req TakeAttendanceRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		facultyID := c.Locals("userID").(string)

		if err := service.TakeAttendance(c.Context(), facultyID, req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Attendance recorded successfully"})
	})

	// Student Route: Get My Attendance
	attendanceGroup.Get("/", middleware.RoleMiddleware("student"), func(c *fiber.Ctx) error {
		studentID := c.Locals("userID").(string)

		records, err := service.GetStudentAttendance(c.Context(), studentID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch attendance"})
		}

		return c.JSON(fiber.Map{
			"student_id": studentID,
			"records":    records,
		})
	})
}
