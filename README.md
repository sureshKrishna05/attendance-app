# University Attendance Management System

A production-ready modular student information system with attendance management as its first core module. Designed for department-level operations at Bharathidasan University.

## 🚀 Overview

The University Attendance Management System provides a secure, reliable, and scalable platform for faculty to mark attendance and internal marks, and for students to track their academic progress in real-time. The system is split into a **Flutter cross-platform frontend** and a **Go (Fiber) robust backend API**.

### 📱 Frontend (Flutter)
- **Framework**: Flutter
- **Supported Platforms**: Android, iOS
- **Features**:
  - Secure Login for Students and Faculty
  - Student Dashboard (Overall attendance, internal marks, timetable)
  - Faculty Dashboard (Manage assigned classes, mark attendance, upload marks)
  - Responsive, Modern "Minimal Government/University" Theme
- **Location**: `/frontend`

### ⚙️ Backend API (Go)
- **Language/Framework**: Go with GoFiber
- **Database**: PostgreSQL (using `pgx`)
- **Architecture**: Modular Monolith with Clean Architecture principles
- **Security**: 
  - JWT (JSON Web Tokens) for user authentication
  - Publishable API Key (`X-Api-Key`) for client verification
- **Location**: `/attendance-api`

## 🛠 Project Structure

```text
.
├── attendance-api/         # Go Backend Service
│   ├── cmd/server/         # Application Entrypoint
│   ├── configs/            # Configuration management (.env)
│   ├── internal/           # Business logic, routes, and middlewares
│   └── pkg/                # Reusable utilities (JWT, etc.)
├── frontend/               # Flutter Frontend Application
│   ├── lib/
│   │   ├── screens/        # UI Views (Login, Dashboards)
│   │   └── theme/          # UI Styling (Colors, Typography)
└── README.md               # Project Documentation
```

## 🔐 Security Features

1. **Client Verification**: The Go backend requires an `X-Api-Key` header with a valid publishable key on all protected API routes to prevent unauthorized access by third-party clients.
2. **User Authentication**: Upon a successful login, the API issues a signed JWT which the client must present as a `Bearer` token in the `Authorization` header.

## 🏃 Getting Started

### Running the Backend

1. Navigate to the API directory:
   ```bash
   cd attendance-api
   ```
2. Ensure you have your `.env` configured properly (refer to `configs/config.go` for required keys).
3. Install dependencies and run:
   ```bash
   go mod tidy
   go run cmd/server/main.go
   ```
The backend server will start on `http://localhost:8080`.

### Running the Frontend

1. Navigate to the Flutter directory:
   ```bash
   cd frontend
   ```
2. Run the application:
   ```bash
   flutter run
   ```

## 📄 License

This project is licensed under the [MIT License](LICENSE).
