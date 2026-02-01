package models

import "time"

type AIExecution struct {
	ID uint `gorm:"primaryKey"`
	UserID string `gorm:"type:uuid;not null;index"`

	FunctionName string `gorm:"size:50;not null"`
	InputUser string `gorm:"type:text;not null"`
	AIResult string `gorm:"type:jsonb;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// CREATE TABLE ai_executions(
// 	id SERIAL PRIMARY KEY,
// 	user_id uuid NOT NULL,

// 	function_name VARCHAR(50) NOT NULL,
// 	input_user TEXT NOT NULL,
// 	ai_result JSONB NOT NULL,

// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

// 	FOREIGN KEY(user_id) REFERENCES users(id)
// )

// CREATE INDEX idx_ai_executions_user_id ON ai_executions(user_id);