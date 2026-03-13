package models

import "time"

type TimeWheel struct {
	interval  time.Duration
	slots     []*Slot
	current   int
	wheelType string
}
