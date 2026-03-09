package models

import "time"

type FollowRequest struct {
	ID uint `gorm:"primaryKey"`
	FromUserID string `gorm:"type:uuid;not null;index"`
	ToUserID string `gorm:"type:uuid;not null;index"`

	Status string `gorm:"size:30;default:'pending'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	FromUser *User `gorm:"foreignKey:FromUserID;references:ID"`
	ToUser *User `gorm:"foreignKey:ToUserID;references:ID"`
}

// CREATE TABLE follow_requests (
// 	id SERIAL PRIMARY KEY,
// 	from_user_id UUID NOT NULL,
// 	to_user_id UUID NOT NULL,
// 	status VARCHAR(30) DEFAULT 'pending',
// 	FOREIGN KEY(from_user_id) REFERENCES users(id),
// 	FOREIGN KEY (to_user_id) REFERENCES users(id)
// )

// CREATE INDEX idx_follow_request_from_user_id ON follow_request(from_user_id);
// CREATE INDEX idx_follow_request_to_user_id ON follow_request(to_user_id);