package admin

type BulkUser struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Name     string `json:"name"`
}

type Class struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Semester int    `json:"semester"`
	Section  string `json:"section"`
}

type TimetableSlot struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	DayOfWeek string `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	SubjectID string `json:"subject_id"`
	FacultyID string `json:"faculty_id"`
	Room      string `json:"room"`
}

type FacultyAssignment struct {
	ClassID   string `json:"class_id"`
	SubjectID string `json:"subject_id"`
	FacultyID string `json:"faculty_id"`
}
