package model

import "time"

type Employee struct {
	ID        int
	Name      string
	Salary    float64
	Position  string
	DeletedAt *time.Time
}
