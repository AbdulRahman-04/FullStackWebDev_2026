package models

import "time"

type Admin struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Role string `gorm:"size:20;default:'admin'"`

	AdminName      string `gorm:"size:50;not null"`
	Email          string `gorm:"size:50;not null;uniqueIndex"`
	Password       string `gorm:"not null"`
	Phone          string `gorm:"size:20;not null;uniqueIndex"`
	ProfilePicture string `gorm:"default:''"`
	Provider       string `gorm:"size:30;not null"`

	EmailVerified bool `gorm:"default:false"`
	PhoneVerified bool `gorm:"default:false"`

	Email_Verify_Token string

	
	RefreshToken string
	RefreshExpiry time.Time

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// CREATE TABLE admins (
// 	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
// 	role VARCHAR(50) DEFAULT 'admin',
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