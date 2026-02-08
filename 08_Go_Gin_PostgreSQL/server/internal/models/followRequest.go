package models

import "time"

type FollowRequest struct {
	ID uint `gorm:"primaryKey"`
	FromUserID string `gorm:"type:uuid;not null;index"`
	ToUserID string `gorm:"type:uuid;not null;index"`

	Status string `gorm:"default:'pending'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	FromUser *User `gorm:"foreignKey:FromUserID;references:ID"`
	ToUser *User `gorm:"foreignKey:ToUserID;references:ID"` 
}

// CREATE TABLE follow_requests (
// 	id SERIAL PRIMARY KEY,
// 	from_user_id uuid NOT NULL,
// 	to_user_id uuid NOT NULL,

// 	status TEXT DEFAULT 'pending',

// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

// 	FOREIGN KEY(from_user_id) REFERENCES users(id),
// 	FOREIGN KEY (to_user_id) REFERENCES users(id)
// )

// CREATE INDEX idx_follow_requests_from_user_id ON follow_requests(from_user_id);
// CREATE INDEX idx_follow_requests_to_user_id ON follow_requests(to_user_id)