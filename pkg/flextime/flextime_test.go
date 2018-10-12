package flextime_test

import (
	"time"

	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/johncornish/flextime-go/pkg/flextime"
)

var _ = Describe("Flextime", func() {
	var (
		now      time.Time
		location time.Location
	)

	BeforeSuite(func() {
		now = time.Now().Local()
		location = time.Location{}
	})

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

		DescribeTable("time repetition",
			func(years, months, days int, repeat string) {
				nextDate := now.AddDate(years, months, days)

				task := flextime.Task{
					DueDate: now,
					Repeat:  repeat,
				}

				next, err := task.Next()
				Expect(next.DueDate).To(Equal(nextDate))
				Expect(err).ToNot(HaveOccurred())
			},
			Entry("days", 0, 0, 1, "1d"),
			Entry("days", 0, 0, 21, "21d"),
			Entry("weeks", 0, 0, 7, "1w"),
			Entry("weeks", 0, 0, 14, "2w"),
			Entry("months", 0, 1, 0, "1m"),
			Entry("months", 0, 2, 0, "2m"),
		)

		DescribeTable("invalid repetitions",
			func(repeat string) {
				task := flextime.Task{
					DueDate: now,
					Repeat:  repeat,
				}

				_, err := task.Next()

				Expect(err).To(HaveOccurred())
			},
			Entry("empty", ""),
			Entry("repeated character", "1dd"),
		)
	})

	Describe("TaskCategory", func() {
		It("should allow user to add tasks", func() {
			tc := flextime.TaskCategory{
				// This vs. separate test to make sure it can be named?
				Name:     "Messaging",
				Contexts: []string{"computer", "phone"},
			}

			task := flextime.Task{
				Name:    "email",
				Repeat:  "1d",
				DueDate: now.AddDate(0, 0, 1),
			}

			tc.AddTask(task)

			Expect(tc.Tasks).To(ContainElement(task))
		})
	})

	Describe("TimeBlock", func() {
		var (
			tb    flextime.TimeBlock
			tc    flextime.TaskCategory
			tasks []flextime.Task
		)

		BeforeEach(func() {
			tb = flextime.TimeBlock{
				Name:  "home",
				Start: now.Local(),
				End:   now.Local().Add(time.Hour),
			}

			tasks = []flextime.Task{
				{
					Name:     "Clean room",
					Estimate: 15 * time.Minute,
					Repeat:   "1d",
					DueDate:  time.Date(2018, 10, 11, 24, 0, 0, 0, &location),
				},
				{
					Name:     "Clean kitchen",
					Repeat:   "6w",
					DueDate:  time.Date(2018, 10, 6, 24, 0, 0, 0, &location),
					Estimate: time.Hour,
				},
				{
					Name:     "Mail",
					Estimate: 20 * time.Minute,
				},
				{
					Name:     "Vacuum",
					Repeat:   "1m",
					DueDate:  time.Date(2018, 10, 15, 24, 0, 0, 0, &location),
					Estimate: 15 * time.Minute,
				},
			}

			tc = flextime.TaskCategory{
				Name:     "Upkeep",
				Contexts: []string{"home"},
				Tasks:    tasks,
			}

			tb.Schedule(tc)
		})

		It("should order by due date, with non-due tasks at the end", func() {
			orderedTasks := []flextime.Task{
				{
					Name:     "Clean kitchen",
					Repeat:   "6w",
					DueDate:  time.Date(2018, 10, 6, 24, 0, 0, 0, &location),
					Estimate: time.Hour,
				},
				{
					Name:     "Clean room",
					Estimate: 15 * time.Minute,
					Repeat:   "1d",
					DueDate:  time.Date(2018, 10, 11, 24, 0, 0, 0, &location),
				},
				{
					Name:     "Vacuum",
					Repeat:   "1m",
					DueDate:  time.Date(2018, 10, 15, 24, 0, 0, 0, &location),
					Estimate: 15 * time.Minute,
				},
				{
					Name:     "Mail",
					Estimate: 20 * time.Minute,
				},
			}

			Expect(tb.Scheduled).To(Equal(orderedTasks))
		})
	})
})
