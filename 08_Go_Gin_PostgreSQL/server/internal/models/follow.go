package models

import "time"

type Follow struct {
	ID uint `gorm:"primaryKey"`
	FollowerID string `gorm:"type:uuid;not null;index"`
	FollowingID string `gorm:"type:uuid;not null;index"`

	Follower *Users `gorm:"foreignKey:FollowerID;references:ID" json:"follower,omitempty"`
	Following *Users `gorm:"foreignKey:FollowingID;references:ID" json:"following,omitempty"`
 
	CreatedAt time.Time
}

// CREATE TABLE follows(
// 	id SERIAL PRIMARY KEY,
// 	follower_id uuid NOT NULL,
// 	following_id uuid NOT NULL,

// 	FOREIGN KEY(follower_id) REFERENCES users(id),
// 	FOREIGN KEY(following_id) REFERENCES users(id),

// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );

// CREATE INDEX idx_follows_follower_id ON follows(follower_id);
// CREATE INDEX idx_follows_following_id ON follows(following_id);