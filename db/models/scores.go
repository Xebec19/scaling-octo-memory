package db

import "time"

type Scores struct {
	TestId      string `gorm:"primaryKey"`
	TestChecker string
	CheckedBy   string
	CheckedOn   time.Time
	Status      string
}
