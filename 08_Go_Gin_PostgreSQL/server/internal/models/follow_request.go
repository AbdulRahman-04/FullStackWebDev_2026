package models

import "time"

type FollowRequest struct {
	ID uint `gorm:"primaryKey"`

	FromUserID string `gorm:"type:uuid;index;not null"`
	ToUserID string `gorm:"type:uuid;index;not null"`

	Status string `gorm:"size:30;default:pending"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// CREATE TABLE follow_request(
// 	id SERIAL PRIMARY KEY,

// 	from_user_id uuid  NOT NULL,
// 	to_user_id uuid NOT NULL,

// 	status VARCHAR(30) DEFAULT 'pending',

// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );

// CREATE INDEX idx_follow_request_from_user_id ON follow_request(from_user_id);
// CREATE INDEX idx_follow_request_to_user_id ON follow_request(to_user_id);