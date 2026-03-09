package models

import "time"

type AIExecution struct {
	ID uint `gorm:"primaryKey"`
    UserID string `gorm:"type:uuid;not null;index"`
	FunctionName string `gorm:"size:50;not null"`
	InputUser string `gorm:"not null"`
	AiResult string `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
}

// CREATE TABLE ai_executions (
// 	id SERIAL PRIMARY KEY,
// 	user_id UUID NOT NULL,
// 	function_name VARCHAR(50) NOT NULL,
// 	input_user TEXT NOT NULL,
// 	ai_result TEXT NOT NULL,
// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

// 	FOREIGN KEY(user_id) REFERENCES users(id)
// )

// CREATE INDEX idx_ai_executions_user_id ON ai_executions(user_id);