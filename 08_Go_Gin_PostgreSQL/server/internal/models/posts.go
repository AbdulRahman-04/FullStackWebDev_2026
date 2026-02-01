package models

import "time"

type Post struct {
	ID uint `gorm:"primaryKey"`
	UserID string `gorm:"type:uuid;not null;index"`

	Caption string `gorm:"size:250;not null"`
	ImageURL string `gorm:"type:text;not null"`
	Location string `gorm:"type:text;not null"`
    // Song string `gorm:"type:text;not null"`

	MyUser *User `gorm:"foreignKey:UserID;references:ID" json:"user_id,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// CREATE TABLE posts(
// 	id SERIAL PRIMARY KEY,
// 	user_id uuid NOT NULL,
// 	caption VARCHAR(250) NOT NULL,
// 	image_url TEXT NOT NULL,
// 	Location TEXT NOT NULL,

// 	FOREIGN KEY(user_id) REFERENCES users(id)
// )

// CREATE INDEX idx_posts_user_id ON posts(user_id);