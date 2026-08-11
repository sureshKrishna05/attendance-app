package main

import (
	"context"
	"fmt"
	"log"

	"attendance-api/internal/config"
	"attendance-api/internal/bootstrap"
	"attendance-api/internal/modules/users"
)

func main() {
	cfg := config.LoadConfig()
	db := bootstrap.ConnectDatabase(cfg)
	defer db.Close()

	repo := users.NewRepository(db)
	service := users.NewService(repo)

	// Seed 2 Faculties
	faculties := []struct{ id, name string }{
		{"FAC001", "Dr. Ramesh"},
		{"FAC002", "Prof. Suresh"},
	}
	for _, f := range faculties {
		if err := service.RegisterUser(context.Background(), f.id, "password123", "faculty", f.name); err != nil {
			log.Printf("Faculty %s already exists or error: %v", f.id, err)
		} else {
			log.Printf("Seeded faculty: %s", f.id)
		}
	}

	// Seed 10 Students
	for i := 1; i <= 10; i++ {
		studentID := fmt.Sprintf("21CS%03d", i)
		studentName := fmt.Sprintf("Student %d", i)
		if err := service.RegisterUser(context.Background(), studentID, "password123", "student", studentName); err != nil {
			log.Printf("Student %s already exists or error: %v", studentID, err)
		} else {
			log.Printf("Seeded student: %s", studentID)
		}
	}

	// Seed 3 Classes
	classQueries := []string{
		`INSERT INTO classes (id, name, semester, section) VALUES ('CS_SEM5_A', 'B.Tech CSE', 5, 'A') ON CONFLICT DO NOTHING;`,
		`INSERT INTO classes (id, name, semester, section) VALUES ('CS_SEM5_B', 'B.Tech CSE', 5, 'B') ON CONFLICT DO NOTHING;`,
		`INSERT INTO classes (id, name, semester, section) VALUES ('CS_SEM3_A', 'B.Tech CSE', 3, 'A') ON CONFLICT DO NOTHING;`,
	}
	for _, q := range classQueries {
		db.Exec(context.Background(), q)
	}

	// Assign Students to Classes
	// 21CS001 to 21CS004 -> CS_SEM5_A
	// 21CS005 to 21CS007 -> CS_SEM5_B
	// 21CS008 to 21CS010 -> CS_SEM3_A
	for i := 1; i <= 10; i++ {
		classID := "CS_SEM3_A"
		if i <= 4 {
			classID = "CS_SEM5_A"
		} else if i <= 7 {
			classID = "CS_SEM5_B"
		}
		studentID := fmt.Sprintf("21CS%03d", i)
		db.Exec(context.Background(), `INSERT INTO class_students (class_id, student_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;`, classID, studentID)
	}

	// Faculty Assignments
	// Dr. Ramesh teaches CS501 in SEM5_A and SEM5_B
	// Prof. Suresh teaches CS301 in SEM3_A
	assignments := []struct{ class, subject, faculty string }{
		{"CS_SEM5_A", "CS501", "FAC001"},
		{"CS_SEM5_B", "CS501", "FAC001"},
		{"CS_SEM3_A", "CS301", "FAC002"},
	}
	for _, a := range assignments {
		db.Exec(context.Background(), `INSERT INTO faculty_assignments (class_id, subject_id, faculty_id) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;`, a.class, a.subject, a.faculty)
	}

	// Seed Timetable
	timetables := []struct{ id, class, day, start, end, subject, faculty, room string }{
		{"TT1", "CS_SEM5_A", "Monday", "09:00:00", "10:00:00", "CS501", "FAC001", "Room 101"},
		{"TT2", "CS_SEM5_A", "Wednesday", "10:00:00", "11:00:00", "CS501", "FAC001", "Room 101"},
		{"TT3", "CS_SEM5_B", "Tuesday", "09:00:00", "10:00:00", "CS501", "FAC001", "Room 102"},
		{"TT4", "CS_SEM5_B", "Thursday", "10:00:00", "11:00:00", "CS501", "FAC001", "Room 102"},
		{"TT5", "CS_SEM3_A", "Monday", "11:00:00", "12:00:00", "CS301", "FAC002", "Room 201"},
		{"TT6", "CS_SEM3_A", "Friday", "10:00:00", "11:00:00", "CS301", "FAC002", "Room 201"},
	}
	for _, t := range timetables {
		db.Exec(context.Background(), `INSERT INTO timetable_slots (id, class_id, day_of_week, start_time, end_time, subject_id, faculty_id, room) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING;`, t.id, t.class, t.day, t.start, t.end, t.subject, t.faculty, t.room)
	}

	log.Println("Massive Data Seeding complete.")
}
