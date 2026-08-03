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
