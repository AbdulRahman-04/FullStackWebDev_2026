package models

import "time"

type User struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Role string `gorm:"size:30;default:'user'"`

	FullName string `gorm:"size:50;not null"`
	Email string `gorm:"size:50;not null;uniqueIndex"`
	Password string `gorm:"not null"`

	Phone string `gorm:"size:20;not null"`

	EmailVerified bool `gorm:"default:false"`
	PhoneVerified bool `gorm:"default:false"`
 
	EmailVerifyToken string 
	PhoneVerifyToken string 

	RefreshToken string 
	RefreshExpiry time.Time

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// CREATE TABLE users (
// 	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
// 	role VARCHAR(30) DEFAULT 'user',
// 	full_name VARCHAR(50) NOT NULL,
// 	email VARCHAR(50) NOT NULL UNIQUE,
// 	password TEXT not NULL,
// 	phone VARCHAR(20) NOT NULL,
// 	email_verified BOOLEAN DEFAULT false,
// 	phone_verified BOOLEAN DEFAULT false,

// 	email_verify_token TEXT,
// 	phone_verify_token TEXT,

// 	refresh_token TEXT,
// 	refresh_expiry TIMESTAMP,

// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// )