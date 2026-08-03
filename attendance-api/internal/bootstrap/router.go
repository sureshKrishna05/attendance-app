package bootstrap

import (
	"attendance-api/internal/config"
	"attendance-api/internal/middleware"
	"attendance-api/internal/modules/attendance"
	"attendance-api/internal/modules/auth"
	"attendance-api/internal/modules/classes"
	"attendance-api/internal/modules/marks"
	"attendance-api/internal/modules/timetable"
	"attendance-api/internal/modules/users"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, db *pgxpool.Pool) {
	// Initialize Repositories
	userRepo := users.NewRepository(db)
	attendanceRepo := attendance.NewRepository(db)
	classesRepo := classes.NewRepository(db)
	timetableRepo := timetable.NewRepository(db)
	marksRepo := marks.NewRepository(db)
	
	// Initialize Services
	userService := users.NewService(userRepo)
	attendanceService := attendance.NewService(attendanceRepo)
	classesService := classes.NewService(classesRepo)
	timetableService := timetable.NewService(timetableRepo)
	marksService := marks.NewService(marksRepo)

	// API v1 Group
	api := app.Group("/api/v1")

	// Setup Auth Routes (Public, but requires API Key)
	auth.SetupRoutes(api, cfg, userService)

	// Protected Routes Group
	protected := api.Group("/", middleware.ProtectedRoute(cfg))

	// Setup Attendance Routes
	attendance.SetupRoutes(protected, attendanceService)

	// Setup Classes Routes
	classes.SetupRoutes(protected, classesService)

	// Setup Timetable Routes
	timetable.SetupRoutes(protected, timetableService)

	// Setup Marks Routes
	marks.SetupRoutes(protected, marksService)

	// Example Protected Route
	protected.Get("/protected", func(c *fiber.Ctx) error {
		userID := c.Locals("userID")
		role := c.Locals("role")
		return c.JSON(fiber.Map{
			"message": "You have accessed a protected route",
			"user":    userID,
			"role":    role,
		})
	})
}
