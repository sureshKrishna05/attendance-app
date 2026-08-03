package main

import (
	"context"
	"log"

	"attendance-api/internal/config"
	"attendance-api/internal/bootstrap"
	"attendance-api/internal/modules/users"
)

func main() {
	cfg := config.LoadConfig()
	db := bootstrap.ConnectDatabase(cfg)
	defer db.Close()

	// Initialize tables
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(50) PRIMARY KEY,
			password_hash VARCHAR(255) NOT NULL,
			role VARCHAR(20) NOT NULL CHECK (role IN ('student', 'faculty', 'admin')),
			name VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS attendance_sessions (
			id VARCHAR(50) PRIMARY KEY,
			subject_id VARCHAR(50) NOT NULL,
			class_id VARCHAR(50) NOT NULL,
			faculty_id VARCHAR(50) NOT NULL REFERENCES users(id),
			date DATE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS attendance_records (
			session_id VARCHAR(50) NOT NULL REFERENCES attendance_sessions(id) ON DELETE CASCADE,
			student_id VARCHAR(50) NOT NULL REFERENCES users(id),
			is_present BOOLEAN NOT NULL,
			PRIMARY KEY (session_id, student_id)
		);

		CREATE TABLE IF NOT EXISTS classes (
			id VARCHAR(50) PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			semester INT NOT NULL,
			section VARCHAR(10) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS class_students (
			class_id VARCHAR(50) NOT NULL REFERENCES classes(id),
			student_id VARCHAR(50) NOT NULL REFERENCES users(id),
			PRIMARY KEY (class_id, student_id)
		);

		CREATE TABLE IF NOT EXISTS faculty_assignments (
			class_id VARCHAR(50) NOT NULL REFERENCES classes(id),
			subject_id VARCHAR(50) NOT NULL,
			faculty_id VARCHAR(50) NOT NULL REFERENCES users(id),
			PRIMARY KEY (class_id, subject_id, faculty_id)
		);

		CREATE TABLE IF NOT EXISTS timetable_slots (
			id VARCHAR(50) PRIMARY KEY,
			class_id VARCHAR(50) NOT NULL REFERENCES classes(id),
			day_of_week VARCHAR(15) NOT NULL,
			start_time TIME NOT NULL,
			end_time TIME NOT NULL,
			subject_id VARCHAR(50) NOT NULL,
			faculty_id VARCHAR(50) NOT NULL REFERENCES users(id),
			room VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS internal_exams (
			id VARCHAR(50) PRIMARY KEY,
			class_id VARCHAR(50) NOT NULL REFERENCES classes(id),
			subject_id VARCHAR(50) NOT NULL,
			name VARCHAR(100) NOT NULL,
			max_marks INT NOT NULL,
			date DATE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS internal_marks (
			exam_id VARCHAR(50) NOT NULL REFERENCES internal_exams(id) ON DELETE CASCADE,
			student_id VARCHAR(50) NOT NULL REFERENCES users(id),
			marks_obtained INT NOT NULL,
			PRIMARY KEY (exam_id, student_id)
		);
	`
	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	repo := users.NewRepository(db)
	service := users.NewService(repo)

	// Seed Student
	err = service.RegisterUser(context.Background(), "21CS001", "password123", "student", "Arun Kumar")
	if err != nil {
		log.Printf("Failed to seed student (might already exist): %v", err)
	} else {
		log.Println("Seeded student: 21CS001")
	}

	// Seed Faculty
	err = service.RegisterUser(context.Background(), "FAC001", "password123", "faculty", "Dr. Ramesh")
	if err != nil {
		log.Printf("Failed to seed faculty (might already exist): %v", err)
	} else {
		log.Println("Seeded faculty: FAC001")
	}

	// Seed Classes and Assignments
	classQuery := `INSERT INTO classes (id, name, semester, section) VALUES ('CS_SEM5_A', 'B.Sc Computer Science', 5, 'A') ON CONFLICT DO NOTHING;`
	db.Exec(context.Background(), classQuery)

	enrollQuery := `INSERT INTO class_students (class_id, student_id) VALUES ('CS_SEM5_A', '21CS001') ON CONFLICT DO NOTHING;`
	db.Exec(context.Background(), enrollQuery)

	assignmentQuery := `INSERT INTO faculty_assignments (class_id, subject_id, faculty_id) VALUES ('CS_SEM5_A', 'CS501', 'FAC001') ON CONFLICT DO NOTHING;`
	db.Exec(context.Background(), assignmentQuery)

	// Seed Timetable
	timetableQuery := `INSERT INTO timetable_slots (id, class_id, day_of_week, start_time, end_time, subject_id, faculty_id, room) 
	VALUES 
	('TT1', 'CS_SEM5_A', 'Monday', '09:00:00', '10:00:00', 'CS501', 'FAC001', 'Room 101'),
	('TT2', 'CS_SEM5_A', 'Wednesday', '10:00:00', '11:00:00', 'CS501', 'FAC001', 'Room 101')
	ON CONFLICT DO NOTHING;`
	db.Exec(context.Background(), timetableQuery)

	// Seed Internal Exams & Marks
	examQuery := `INSERT INTO internal_exams (id, class_id, subject_id, name, max_marks, date)
	VALUES ('EXAM1', 'CS_SEM5_A', 'CS501', 'Internal Assessment 1', 50, '2026-08-01') ON CONFLICT DO NOTHING;`
	db.Exec(context.Background(), examQuery)

	markQuery := `INSERT INTO internal_marks (exam_id, student_id, marks_obtained)
	VALUES ('EXAM1', '21CS001', 45) ON CONFLICT DO NOTHING;`
	db.Exec(context.Background(), markQuery)

	log.Println("Seeding complete.")
}
