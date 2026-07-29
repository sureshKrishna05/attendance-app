-- =====================================================================================
-- UNIVERSITY ATTENDANCE MANAGEMENT SYSTEM - DATABASE SCHEMA
-- Version: 1.1 (Department Level Scope)
-- Architecture: Normalized Relational Database (Supabase-Inspired)
-- =====================================================================================

-- -------------------------------------------------------------------------------------
-- 1. SCHEMA SETUP & ROLES
-- -------------------------------------------------------------------------------------
CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS public;

-- -------------------------------------------------------------------------------------
-- 2. TYPE DEFINITIONS (ENUMS)
-- -------------------------------------------------------------------------------------
-- Profiles & Users
CREATE TYPE public.role_enum AS ENUM ('student', 'faculty', 'admin', 'hod');
CREATE TYPE public.profile_status AS ENUM ('active', 'inactive', 'suspended');

-- Academics
CREATE TYPE public.term_enum AS ENUM ('Odd', 'Even');

-- Specific Roles Status
CREATE TYPE public.student_status AS ENUM ('active', 'graduated', 'withdrawn');
CREATE TYPE public.faculty_status AS ENUM ('active', 'retired', 'on_leave');
CREATE TYPE public.class_student_status AS ENUM ('active', 'transferred');

-- Operations
CREATE TYPE public.attendance_session_status AS ENUM ('open', 'locked');
CREATE TYPE public.attendance_status AS ENUM ('present', 'absent', 'late', 'od', 'medical');
CREATE TYPE public.notification_target AS ENUM ('All', 'Department', 'Class', 'Student', 'Faculty');

-- -------------------------------------------------------------------------------------
-- 3. MOCK AUTH SCHEMA (Required for local testing of triggers)
-- -------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS auth.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    raw_user_meta_data JSONB,
    encrypted_password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- -------------------------------------------------------------------------------------
-- 4. BASE IDENTITY & ORGANIZATIONAL HIERARCHY
-- -------------------------------------------------------------------------------------

-- Profiles (Mirrored from auth)
CREATE TABLE public.profiles (
    id UUID PRIMARY KEY REFERENCES auth.users(id) ON DELETE CASCADE,
    email VARCHAR(255) UNIQUE NOT NULL,
    full_name VARCHAR(255),
    phone VARCHAR(20),
    avatar_url TEXT,
    role public.role_enum NOT NULL,
    status public.profile_status DEFAULT 'active',
    email_verified BOOLEAN DEFAULT FALSE,
    phone_verified BOOLEAN DEFAULT FALSE,
    is_first_login BOOLEAN DEFAULT TRUE,
    last_password_change TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);

-- Departments (hod_id foreign key added later to resolve circular dependency)
CREATE TABLE public.departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    short_name VARCHAR(20) NOT NULL,
    hod_id UUID 
);

-- Programmes
CREATE TABLE public.programmes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    department_id UUID REFERENCES public.departments(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    duration_years INTEGER NOT NULL
);

-- Semesters
CREATE TABLE public.semesters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    programme_id UUID REFERENCES public.programmes(id) ON DELETE CASCADE,
    semester_no INTEGER NOT NULL,
    academic_year VARCHAR(20) NOT NULL,
    term public.term_enum NOT NULL
);

-- Sections
CREATE TABLE public.sections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(20) NOT NULL
);

-- Subjects
CREATE TABLE public.subjects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    credits INTEGER NOT NULL,
    hours_per_week INTEGER NOT NULL,
    department_id UUID REFERENCES public.departments(id) ON DELETE CASCADE
);

-- -------------------------------------------------------------------------------------
-- 5. ROLE-SPECIFIC DETAILS & CIRCULAR DEPENDENCY RESOLUTION
-- -------------------------------------------------------------------------------------

-- Student Details
CREATE TABLE public.student_details (
    profile_id UUID PRIMARY KEY REFERENCES public.profiles(id) ON DELETE CASCADE,
    register_number VARCHAR(30) UNIQUE NOT NULL,
    roll_number VARCHAR(30),
    admission_number VARCHAR(30),
    admission_date DATE,
    department_id UUID REFERENCES public.departments(id),
    programme_id UUID REFERENCES public.programmes(id),
    semester_id UUID REFERENCES public.semesters(id),
    section_id UUID REFERENCES public.sections(id),
    dob DATE,
    gender VARCHAR(20),
    blood_group VARCHAR(5),
    address TEXT,
    parent_name VARCHAR(255),
    parent_phone VARCHAR(20),
    guardian_email VARCHAR(255),
    emergency_contact VARCHAR(20),
    photo_url TEXT,
    batch VARCHAR(20),
    status public.student_status DEFAULT 'active'
);

-- Faculty Details
CREATE TABLE public.faculty_details (
    profile_id UUID PRIMARY KEY REFERENCES public.profiles(id) ON DELETE CASCADE,
    employee_id VARCHAR(30) UNIQUE NOT NULL,
    designation VARCHAR(100),
    education TEXT,
    highest_qualification VARCHAR(100),
    specialization TEXT,
    research_area TEXT,
    department_id UUID REFERENCES public.departments(id),
    joining_date DATE,
    experience_years INTEGER,
    employment_type VARCHAR(50),
    dob DATE,
    address TEXT,
    office_phone VARCHAR(20),
    office_room VARCHAR(50),
    status public.faculty_status DEFAULT 'active'
);

-- Resolve Circular Dependency: Assign HOD to Department
ALTER TABLE public.departments 
ADD CONSTRAINT fk_department_hod 
FOREIGN KEY (hod_id) REFERENCES public.faculty_details(profile_id);

-- -------------------------------------------------------------------------------------
-- 6. CLASS & SUBJECT ASSIGNMENTS
-- -------------------------------------------------------------------------------------

-- Classes
CREATE TABLE public.classes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    programme_id UUID REFERENCES public.programmes(id),
    semester_id UUID REFERENCES public.semesters(id),
    section_id UUID REFERENCES public.sections(id),
    mentor_id UUID REFERENCES public.faculty_details(profile_id),
    room VARCHAR(30),
    strength INTEGER
);

-- Class Students
CREATE TABLE public.class_students (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    class_id UUID REFERENCES public.classes(id) ON DELETE CASCADE,
    student_id UUID REFERENCES public.student_details(profile_id) ON DELETE CASCADE,
    joined_on DATE DEFAULT CURRENT_DATE,
    status public.class_student_status DEFAULT 'active'
);

-- Faculty Assignments
CREATE TABLE public.faculty_assignments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    faculty_id UUID REFERENCES public.faculty_details(profile_id),
    class_id UUID REFERENCES public.classes(id),
    subject_id UUID REFERENCES public.subjects(id),
    academic_year VARCHAR(20) NOT NULL
);

-- Class Subjects
CREATE TABLE public.class_subjects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    class_id UUID REFERENCES public.classes(id),
    subject_id UUID REFERENCES public.subjects(id),
    faculty_assignment_id UUID REFERENCES public.faculty_assignments(id)
);

-- Timetable Slots
CREATE TABLE public.timetable_slots (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    class_subject_id UUID REFERENCES public.class_subjects(id) ON DELETE CASCADE,
    day_of_week INTEGER CHECK (day_of_week BETWEEN 1 AND 7),
    hour_number INTEGER NOT NULL,
    room VARCHAR(30)
);

-- -------------------------------------------------------------------------------------
-- 7. CORE OPERATIONS (ATTENDANCE & INTERNAL ASSESSMENTS)
-- -------------------------------------------------------------------------------------

-- Attendance Sessions (Linked directly to timetable slots per v1.1 specifications)
CREATE TABLE public.attendance_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    class_subject_id UUID REFERENCES public.class_subjects(id),
    timetable_slot_id UUID REFERENCES public.timetable_slots(id),
    faculty_id UUID REFERENCES public.faculty_details(profile_id),
    date DATE NOT NULL,
    remarks TEXT,
    status public.attendance_session_status DEFAULT 'open'
);

-- Attendance Records (Includes Correction Tracking)
CREATE TABLE public.attendance_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    attendance_session_id UUID REFERENCES public.attendance_sessions(id) ON DELETE CASCADE,
    student_id UUID REFERENCES public.student_details(profile_id),
    status public.attendance_status NOT NULL,
    marked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    updated_by UUID REFERENCES public.faculty_details(profile_id),
    correction_reason TEXT
);

-- Internal Exams
CREATE TABLE public.internal_exams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    class_subject_id UUID REFERENCES public.class_subjects(id) ON DELETE CASCADE,
    maximum_marks INTEGER NOT NULL,
    exam_date DATE NOT NULL,
    created_by UUID REFERENCES public.faculty_details(profile_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Internal Marks
CREATE TABLE public.internal_marks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    internal_exam_id UUID REFERENCES public.internal_exams(id) ON DELETE CASCADE,
    student_id UUID REFERENCES public.student_details(profile_id) ON DELETE CASCADE,
    marks_obtained DECIMAL(5,2) NOT NULL,
    remarks TEXT,
    entered_by UUID REFERENCES public.faculty_details(profile_id),
    entered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- -------------------------------------------------------------------------------------
-- 8. UTILITY & SYSTEM TABLES
-- -------------------------------------------------------------------------------------

-- Notifications
CREATE TABLE public.notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    target_type public.notification_target NOT NULL,
    created_by UUID REFERENCES public.profiles(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Files
CREATE TABLE public.files (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID REFERENCES public.profiles(id) ON DELETE CASCADE,
    file_name VARCHAR(255) NOT NULL,
    bucket_name VARCHAR(100) NOT NULL,
    visibility VARCHAR(50) DEFAULT 'private',
    mime_type VARCHAR(100) NOT NULL,
    file_size BIGINT NOT NULL,
    file_extension VARCHAR(10),
    checksum VARCHAR(255),
    path TEXT NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Audit Logs
CREATE TABLE public.audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES public.profiles(id),
    session_id VARCHAR(255),
    request_id VARCHAR(255),
    user_agent TEXT,
    action VARCHAR(100) NOT NULL,
    table_name VARCHAR(100) NOT NULL,
    record_id UUID NOT NULL,
    old_data JSONB,
    new_data JSONB,
    ip_address INET,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- -------------------------------------------------------------------------------------
-- 9. SUPABASE-INSPIRED AUTHENTICATION TRIGGERS
-- -------------------------------------------------------------------------------------

-- Function to handle new user signup and mirror to public.profiles
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS trigger AS $$
BEGIN
  INSERT INTO public.profiles (id, email, full_name, role)
  VALUES (
    new.id, 
    new.email, 
    new.raw_user_meta_data->>'full_name', 
    (new.raw_user_meta_data->>'role')::public.role_enum
  );
  RETURN new;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- Trigger attached to auth.users
CREATE TRIGGER on_auth_user_created
  AFTER INSERT ON auth.users
  FOR EACH ROW EXECUTE PROCEDURE public.handle_new_user();

-- =====================================================================================
-- END OF SCRIPT
-- =====================================================================================
