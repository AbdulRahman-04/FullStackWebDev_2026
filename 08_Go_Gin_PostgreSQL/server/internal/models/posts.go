package models

import "time"

type Posts struct {
	ID uint `gorm:"primaryKey"`
	UserID string `gorm:"type:uuid;not null;index"`

	Caption string `gorm:"type:text;not null"`
	ImageURL string `gorm:"not null"`
	Location string `gorm:"not null"`

	IsPublic bool `gorm:"not null;default:true;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}


// CREATE TABLE posts(
// 	id SERIAL PRIMARY KEY,
// 	user_id uuid NOT NULL,
// 	caption TEXT NOT NULL,
// 	image_url TEXT NOT NULL,
// 	location TEXT NOT NULL,
// 	is_public BOOLEAN DEFAULT TRUE,

// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 
// 	FOREIGN KEY(user_id) REFERENCES users(id) 
// )

// CREATE INDEX idx_posts_user_id ON posts(user_id);
// CREATE INDEX idx_posts_is_public ON posts(is_public)