package models

import "time"

type Post struct {
	ID uint `gorm:"primaryKey"`
	UserID string `gorm:"type:uuid;not null;index"`

	Caption string `gorm:"size:250;not null"`
	Song string `gorm:"not null"`
	Location string `gorm:"not null"`
	ImageUrl string `gorm:"not null"`
	IsPublic bool `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
}

// CREATE TABLE posts (
// 	id SERIAL PRIMARY KEY,
// 	user_id UUID NOT NULL,
// 	caption VARCHAR(250) NOT NULL,
// 	song TEXT NOT NULL,
// 	location TEXT NOT NULL,
// 	image_url TEXT NOT NULL,
// 	is_public BOOLEAN NOT NULL,
// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// )

// CREATE INDEX idx_posts_user_id ON posts(user_id)