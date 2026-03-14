package algo

import (
	"fmt"
	"time"

	"github.com/Deepakraj15/task-manager/internal/models"
)

func StartWheel() {

	seconds := models.GetSecondsWheel()

	seconds.Interval = time.Second
	seconds.Slots = make([]*models.Slot, 60)

	for i := range seconds.Slots {
		seconds.Slots[i] = &models.Slot{}
	}

	ticker := time.NewTicker(time.Second)

	for {
		<-ticker.C
		tick(seconds)
	}
}

func tick(wheel *models.TimeWheel) {

	slot := wheel.Slots[wheel.Current]

	for _, task := range slot.Tasks {
		go executeTask(task)
	}

	slot.Tasks = nil

	wheel.Current = (wheel.Current + 1) % len(wheel.Slots)
}

func SubmitTask(task models.Task) {

	wheel := models.GetSecondsWheel()

	delay := time.Until(task.RunAt)

	slot := (wheel.Current + int(delay.Seconds())) % len(wheel.Slots)

	wheel.Slots[slot].Tasks = append(wheel.Slots[slot].Tasks, &task)
}

func executeTask(task *models.Task) {
	fmt.Printf("Executed the task %v at %v \n Description: %v", task.TName, task.RunAt, task.TDescription)

}
