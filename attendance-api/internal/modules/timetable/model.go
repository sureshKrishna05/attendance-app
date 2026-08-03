package timetable

type Slot struct {
	ID          string `json:"id"`
	ClassID     string `json:"class_id"`
	DayOfWeek   string `json:"day_of_week"`
	StartTime   string `json:"start_time"` // Stored as TIME in DB, can scan as string
	EndTime     string `json:"end_time"`
	SubjectID   string `json:"subject_id"`
	FacultyID   string `json:"faculty_id"`
	FacultyName string `json:"faculty_name,omitempty"` // Joined from users table
	Room        string `json:"room"`
}
