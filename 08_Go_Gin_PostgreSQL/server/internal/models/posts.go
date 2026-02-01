package models

import "time"

type Post struct {
	ID     string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID string `gorm:"type:uuid;not null;index"`

	Caption  string `gorm:"size:250;not null"`
	ImageURL string `gorm:"not null"`
	Location string `gorm:"not null"`

	IsPublic bool `gorm:"default:true;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// CREATE TABLE posts(
// 	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
// 	user_id uuid NOT NULL,
// 	caption VARCHAR(250) NOT NULL,
// 	image_url TEXT NOT NULL,
// 	location TEXT NOT NULL,

// 	is_public BOOLEAN DEFAULT TRUE,
// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );

// CREATE INDEX idx_posts_user_id ON posts(user_id);

// CREATE INDEX idx_posts_is_public ON posts(is_public);