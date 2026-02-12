package models

import "time"

type Post struct {
	ID uint `gorm:"primaryKey"`
	UserID string `gorm:"type:uuid;not null;index"`

	Caption string `gorm:"size:250;not null"`
	Song string `gorm:"type:text;not null"`
	ImageUrl string `gorm:"type:text;not null"`
	Location string `gorm:"type:text;not null"`
	IsPublic bool `gorm:"type:boolean;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
}


// CREATE TABLE posts (
// 	id SERIAL PRIMARY KEY,
// 	user_id uuid NOT NULL,

// 	caption VARCHAR(250) not null,
// 	song TEXT NOT NULL,
// 	image_url TEXT NOT NULL,
// 	location TEXT NOT NULL,
// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

// 	FOREIGN KEY(user_id) REFERENCES users(id)
// )

// CREATE INDEX idx_posts_user_id ON posts(user_id)