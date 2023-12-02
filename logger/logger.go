package logger

import (
	"context"
	"time"
	"web-calculator/model"

	"gorm.io/gorm"
)	

type Logger struct {
	Client *gorm.DB
}

// logging function
func (l *Logger) Log(ctx context.Context, eq string, total float64) error {

	// time now
	utc := time.Now().UTC()

	// create equation model
	equation := &model.EquationModel{
		Equation: eq,
		Result: total,
		CalculatedAt: &utc,
	}

	// add to db
	err := l.Client.Create(equation).Error; if err != nil {
		return err
	}

	return nil
}
