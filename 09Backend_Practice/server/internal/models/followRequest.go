package models

import "time"

type FollowRequest struct {
	ID uint `gorm:"primaryKey"`

	FromUserID string `gorm:"index;not null"`
	ToUserID   string `gorm:"index;not null"`

	Status string `gorm:"type:varchar(20);default:'pending'"`
	// pending | rejected

	CreatedAt time.Time
	UpdatedAt time.Time

	FromUser *User `gorm:"foreignKey:FromUserID;references:ID"`
	ToUser   *User `gorm:"foreignKey:ToUserID;references:ID"`
}

// CREATE TABLE follow_requests (
// 	id SERIAL PRIMARY KEY,
// 	from_user_id UUID NOT NULL,
// 	to_user_id UUID NOT NULL,
// 	status VARCHAR(20) DEFAULT 'pending',
//
// 	FOREIGN KEY (from_user_id) REFERENCES users(id) ON DELETE CASCADE,
// 	FOREIGN KEY (to_user_id) REFERENCES users(id) ON DELETE CASCADE
// );
//
// CREATE INDEX ON follow_requests(from_user_id);
// CREATE INDEX ON follow_requests(to_user_id);