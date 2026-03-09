package models

import "time"

type User struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Role string `gorm:"size:50;default:'user'"`

	FullName string `gorm:"size:50;not null"`
	Email    string `gorm:"size:50;not null;uniqueIndex"`
	Password string `gorm:"not null"`

	Phone          string `gorm:"size:20;not null;uniqueIndex"`
	Provider       string `gorm:"size:30;not null"`
	ProfilePicture string `gorm:"default:''"`

	EmailVerified bool `gorm:"default:false"`
	PhoneVerified bool `gorm:"default:false"`

	EmailVerifyToken string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// CREATE TABLE users (
// 	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
// 	role VARCHAR(50) DEFAULT 'user',
// 	full_name VARCHAR(50) NOT NULL,
// 	email VARCHAR(50) NOT NULL UNIQUE,
// 	password TEXT NOT NULL UNIQUE,
// 	phone VARCHAR(20) NOT NULL,
// 	provider VARCHAR(30) NOT NULL,
// 	profile_picture TEXT DEFAULT '',
// 	email_verified BOOLEAN DEFAULT false,
// 	phone_verified BOOLEAN DEFAULT false,

//  email_verify_token TEXT,

// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// )