package models

import (
	"time"

	"github.com/Deepakraj15/task-manager/internal/constants"
)

type TimeWheel struct {
	interval  time.Duration
	slots     []*Slot
	current   int
	wheelType string
}

var hoursWheel *TimeWheel
var minutesWheel *TimeWheel
var secondsWheel *TimeWheel

func GetHoursWheel() *TimeWheel {
	if hoursWheel == nil {
		hoursWheel = &TimeWheel{
			wheelType: constants.HOURS,
		}
	}
	return hoursWheel
}

func GetMinutesWheel() *TimeWheel {
	if minutesWheel == nil {
		minutesWheel = &TimeWheel{
			wheelType: constants.MINUTES,
		}
	}
	return minutesWheel
}

func GetSecondsWheel() *TimeWheel {
	if secondsWheel == nil {
		secondsWheel = &TimeWheel{
			wheelType: constants.SECONDS,
		}
	}
	return secondsWheel
}
