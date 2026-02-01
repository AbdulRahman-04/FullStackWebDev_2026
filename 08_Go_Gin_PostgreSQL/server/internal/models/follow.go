package models

import "time"

type Follow struct {
	ID uint `gorm:"primaryKey"`
	FollowerID string `gorm:"type:uuid;index;not null"`
	FollowingID string `gorm:"type:uuid;index;not null"`

	MyFollower *User `gorm:"foreignKey:FollowerID;references:ID" json:"follower_id,omitempty"`
	MyFollowing *User `gorm:"foreignKey:FollowingID;references:ID" json:"following_id,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// CREATE TABLE follows(
// 	id SERIAL PRIMARY KEY,
// 	follower_id uuid NOT NULL,
// 	following_id uuid NOT NULL,

// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

// 	FOREIGN KEY(follower_id) REFERENCES users(id),
// 	FOREIGN KEY(following_id) REFERENCES users(id)
// )

// CREATE INDEX idx_follows_follower_id ON follows(follower_id);
// CREATE INDEX idx_follows_following_id ON follows(following_id);