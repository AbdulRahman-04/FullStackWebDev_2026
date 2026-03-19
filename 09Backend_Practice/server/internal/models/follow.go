package models

import "time"

type Follow struct {
	ID uint `gorm:"primaryKey"`

	FollowerID  string `gorm:"index;not null"`
	FollowingID string `gorm:"index;not null"`

	Follower  *User `gorm:"foreignKey:FollowerID;references:ID" json:"follower,omitempty"`
	Following *User `gorm:"foreignKey:FollowingID;references:ID" json:"following,omitempty"`

	CreatedAt time.Time
}

// CREATE TABLE follows (
// 	id SERIAL PRIMARY KEY,
// 	follower_id UUID NOT NULL,
// 	following_id UUID NOT NULL,
// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//
// 	FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
// 	FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
// );
//
// CREATE INDEX ON follows(follower_id);
// CREATE INDEX ON follows(following_id);