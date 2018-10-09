package flextime_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/johncornish/flextime-go/pkg/flextime"
)

var _ = Describe("Flextime", func() {
	Describe("Task", func() {
		It("should know whether it's due or not", func() {
			task1 := flextime.Task{
				DueDate: time.Now().Local().AddDate(0, 0, 1),
			}
			task2 := flextime.Task{
				DueDate: time.Now().Local().AddDate(0, 0, -1),
			}

			Expect(task1.IsDue()).To(Equal(false))
			Expect(task2.IsDue()).To(Equal(true))
		})

		// It("should support repetition by day, week, and month", func() {
		It("should support repetition by number of days", func() {
			today := time.Now().Local().AddDate(0, 0, 0)
			tomorrow := today.AddDate(0, 0, 1)

			task := flextime.Task{
				DueDate: today,
				Repeat:  "1d",
			}

			nextTask := task.Next()
			Expect(nextTask.DueDate).To(Equal(tomorrow))
		})
	})

	Describe("TimeBlock", func() {
		It("should be able to be instantiated", func() {
			tb := flextime.TimeBlock{}
			_ = tb
		})
	})
})
