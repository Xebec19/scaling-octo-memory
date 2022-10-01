package db

import "time"

type Group struct {
	GroupId   string    `json:"group_id"`
	GroupName string    `json:"group_name"`
	Teachers  []string  `json:"teachers"`
	Students  []string  `json:"students"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}
