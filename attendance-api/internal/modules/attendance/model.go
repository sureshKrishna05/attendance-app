package attendance

import "time"

type Session struct {
	ID        string    `json:"id"`
	SubjectID string    `json:"subject_id"`
	ClassID   string    `json:"class_id"`
	FacultyID string    `json:"faculty_id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

type Record struct {
	SessionID string `json:"session_id"`
	StudentID string `json:"student_id"`
	IsPresent bool   `json:"is_present"`
}

// Request Payload for Faculty taking attendance
type TakeAttendanceRequest struct {
	SubjectID string   `json:"subject_id"`
	ClassID   string   `json:"class_id"`
	Date      string   `json:"date"` // e.g. "2025-05-20"
	Records   []Record `json:"records"`
}

type StudentAttendanceResult struct {
	SubjectID string `json:"subject_id"`
	Date      string `json:"date"`
	IsPresent bool   `json:"is_present"`
}
