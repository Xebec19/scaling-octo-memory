package db

import "time"

type Group struct {
	GroupId   string `gorm:"primaryKey"`
	GroupName string
	Teachers  []string
	Students  []string
	CreatedOn time.Time
	UpdatedOn time.Time
}
