package models

import "time"

type FollowRequest struct {

	ID uint `gorm:"primaryKey"`
	FromUserID string `gorm:"type:uuid;index;not null"`
	ToUserID string `gorm:"type:uuid;index;not null"`

	Status string `gorm:"default:pending"`

	FromUser *User `gorm:"foreignKey:FromUserID;references:ID" json:"from_user_id,omitempty"`
	ToUser *User `gorm:"foreignKey:ToUserID;references:ID" json:"to_user_id,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// CREATE TABLE follow_requests(
// 	id SERIAL PRIMARY KEY,
// 	from_user_id uuid NOT NULL,
// 	to_user_id uuid NOT NULL,
// 	status TEXT DEFAULT 'pending',

// 	FOREIGN KEY(from_user_id) REFERENCES users(id),
// 	FOREIGN KEY(to_user_id) REFERENCES users(id)
// )

// CREATE INDEX idx_follow_requests_from_user_id ON follow_requests(from_user_id);
// CREATE INDEX idx_follow_requests_to_user_id ON follow_requests(to_user_id);