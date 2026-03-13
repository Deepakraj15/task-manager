package models

import "time"

type Task struct {
	TName        string
	TDescription string
	RunAt        time.Time
}
