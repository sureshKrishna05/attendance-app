package marks

type Exam struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	SubjectID string `json:"subject_id"`
	Name      string `json:"name"`      // e.g., "Internal Assessment 1"
	MaxMarks  int    `json:"max_marks"`
	Date      string `json:"date"`      // format YYYY-MM-DD
}

type Mark struct {
	ExamID        string `json:"exam_id"`
	StudentID     string `json:"student_id"`
	MarksObtained int    `json:"marks_obtained"`
}

// Request Payload for Faculty uploading marks
type UploadMarksRequest struct {
	ClassID   string `json:"class_id"`
	SubjectID string `json:"subject_id"`
	ExamName  string `json:"exam_name"`
	MaxMarks  int    `json:"max_marks"`
	Date      string `json:"date"`
	Marks     []Mark `json:"marks"`
}

// Response structure for Student viewing marks
type StudentMarkResult struct {
	SubjectID     string `json:"subject_id"`
	ExamName      string `json:"exam_name"`
	MarksObtained int    `json:"marks_obtained"`
	MaxMarks      int    `json:"max_marks"`
	Date          string `json:"date"`
}
