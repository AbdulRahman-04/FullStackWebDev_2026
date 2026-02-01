package models

import "time"

type Follow struct {
	ID uint `gorm:"primaryKey"`

	FollowerID string `gorm:"type:uuid;index;not null"`
	FollowingID string `gorm:"type:uuid;index;not null"`

	Follower *User `gorm:"foreignKey:FollowerID;references:ID" json:"follower,omitempty"`
	Following *User `gorm:"foreignKey:FollowingID;references:ID" json:"following,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

CREATE TABLE follow(
	id SERIAL PRIMARY KEY,

	follower_id uuid NOT NULL,
	following_id uuid NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY(follower_id) REFERENCES  users(id),
	FOREIGN KEY(following_id) REFERENCES users(id)
);

CREATE INDEX idx_follow_follower_id ON follow(follower_id);
CREATE INDEX idx_follow_following_id ON follow(following_id);