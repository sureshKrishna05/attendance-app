package marks

import (
	"attendance-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, service Service) {
	marksGroup := router.Group("/marks")

	// Faculty Route: Upload Marks
	marksGroup.Post("/faculty", middleware.RoleMiddleware("faculty"), func(c *fiber.Ctx) error {
		facultyID := c.Locals("userID").(string)
		
		var req UploadMarksRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if err := service.UploadMarks(c.Context(), facultyID, req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload marks"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Marks uploaded successfully"})
	})

	// Student Route: View My Marks
	marksGroup.Get("/student", middleware.RoleMiddleware("student"), func(c *fiber.Ctx) error {
		studentID := c.Locals("userID").(string)

		marks, err := service.GetStudentMarks(c.Context(), studentID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch student marks"})
		}

		return c.JSON(fiber.Map{
			"student_id": studentID,
			"marks":      marks,
		})
	})
}
