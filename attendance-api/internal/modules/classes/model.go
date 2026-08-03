package classes

type Class struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Semester int    `json:"semester"`
	Section  string `json:"section"`
}

type FacultyAssignment struct {
	ClassID   string `json:"class_id"`
	SubjectID string `json:"subject_id"`
	ClassName string `json:"class_name"`
	Semester  int    `json:"semester"`
	Section   string `json:"section"`
}
