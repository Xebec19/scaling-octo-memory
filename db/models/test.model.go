package db

import "time"

type Test struct {
	TestId      string `gorm:"primaryKey"`
	Title       string
	Description string
	CreatedBy   string
	Enrolled    []string
	Schedule    time.Time
	StartTime   time.Time
	EndTime     time.Time
	Link        string
	CreatedOn   time.Time
	UpdatedOn   time.Time
}
