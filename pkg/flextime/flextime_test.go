package flextime_test

import (
	"time"

	. "github.com/onsi/ginkgo/extensions/table"

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

		DescribeTable("repetition",
			func(years, months, days int, repeat string) {
				today := time.Now().Local()
				nextDate := today.AddDate(years, months, days)

				task := flextime.Task{
					DueDate: today,
					Repeat:  repeat,
				}

				nextTask := task.Next()
				Expect(nextTask.DueDate).To(Equal(nextDate))
			},
			Entry("days", 0, 0, 1, "1d"),
			Entry("days", 0, 0, 21, "21d"),
			Entry("weeks", 0, 0, 7, "1w"),
			Entry("weeks", 0, 0, 14, "2w"),
			Entry("months", 0, 1, 0, "1m"),
			Entry("months", 0, 2, 0, "2m"),
		)
	})

	Describe("TimeBlock", func() {
		It("should be able to be instantiated", func() {
			tb := flextime.TimeBlock{}
			_ = tb
		})
	})
})
