package users

import "time"

type User struct {
	ID           string    `json:"id"`            // e.g., '21CS001' or 'FAC001'
	PasswordHash string    `json:"-"`             // Hide from JSON serialization
	Role         string    `json:"role"`          // 'student' or 'faculty'
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
