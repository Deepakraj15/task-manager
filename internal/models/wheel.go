package models

import (
	"time"

	"github.com/Deepakraj15/task-manager/internal/constants"
)

type TimeWheel struct {
	Interval  time.Duration
	Slots     []*Slot
	Current   int
	WheelType string
	Next      *TimeWheel
}

var hoursWheel *TimeWheel
var minutesWheel *TimeWheel
var secondsWheel *TimeWheel

// helper to create a wheel
func newWheel(interval time.Duration, size int, wheelType string) *TimeWheel {

	slots := make([]*Slot, size)

	for i := range slots {
		slots[i] = &Slot{}
	}

	return &TimeWheel{
		Interval:  interval,
		Slots:     slots,
		Current:   0,
		WheelType: wheelType,
	}
}

// Hours Wheel (24 slots)
func GetHoursWheel() *TimeWheel {

	if hoursWheel == nil {
		hoursWheel = newWheel(
			time.Hour,
			24,
			constants.HOURS,
		)
	}

	return hoursWheel
}

// Minutes Wheel (60 slots)
func GetMinutesWheel() *TimeWheel {

	if minutesWheel == nil {
		minutesWheel = newWheel(
			time.Minute,
			60,
			constants.MINUTES,
		)
	}

	return minutesWheel
}

// Seconds Wheel (60 slots)
func GetSecondsWheel() *TimeWheel {

	if secondsWheel == nil {
		secondsWheel = newWheel(
			time.Second,
			60,
			constants.SECONDS,
		)
	}

	return secondsWheel
}

func InitTimeWheels() {

	seconds := GetSecondsWheel()
	minutes := GetMinutesWheel()
	hours := GetHoursWheel()

	seconds.Next = minutes
	minutes.Next = hours
}
