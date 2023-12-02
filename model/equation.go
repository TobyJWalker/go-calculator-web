package model

import "time"

// equation model
type EquationModel struct {
	Equation string `json:"equation"`
	Result string `json:"result"`
	CalculatedAt *time.Time `json:"calculated_at"`
}