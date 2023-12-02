package model

import (
	"time"

	"gorm.io/gorm"
)

// equation model
type EquationModel struct {
	gorm.Model
	Equation string `json:"equation"`
	Result float64 `json:"result"`
	CalculatedAt *time.Time `json:"calculated_at"`
}