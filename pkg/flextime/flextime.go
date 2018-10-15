package flextime

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/stretchr/stew/slice"
)

type Task struct {
	Name     string
	DueDate  time.Time
	Estimate time.Duration
	Repeat   string
}

func (t Task) IsDue() bool {
	return t.DueDate.Before(time.Now().Local())
}

func (t Task) Next() (Task, error) {
	var task Task

	next, err := nextTime(t.DueDate, t.Repeat)
	if err != nil {
		return task, err
	}

	task = Task{
		DueDate: next,
		Repeat:  t.Repeat,
	}

	return task, nil
}

type TimeBlock struct {
	Name      string
	Start     time.Time
	End       time.Time
	Scheduled []Task
}

func (tb TimeBlock) AvailableTime() time.Duration {
	return tb.End.Sub(tb.Start)
}

func (tb *TimeBlock) Schedule(tCats ...TaskCategory) {
	var taskTime time.Duration
	var tasks []Task

	tb.Scheduled = []Task{}

	for _, tc := range tCats {
		if !slice.Contains(tc.Contexts, tb.Name) {
			continue
		}
		tasks = append(tasks, tc.Tasks...)
	}
	sort.Sort(byDue(tasks))

	for _, task := range tasks {
		taskTime += task.Estimate
		if taskTime > tb.AvailableTime() {
			break
		}
		tb.Scheduled = append(tb.Scheduled, task)
	}
}

type TaskCategory struct {
	Name     string
	Contexts []string
	Tasks    []Task
}

func (tc *TaskCategory) AddTask(task Task) {
	tc.Tasks = append(tc.Tasks, task)
}

func repeatNum(repeat string) int {
	num64, _ := strconv.ParseInt(repeat[:len(repeat)-1], 10, 64)
	return int(num64)
}

func nextDay(t time.Time, repeat string) (time.Time, error) {
	num64, _ := strconv.ParseInt(repeat, 10, 64)
	day := int(num64)

	newTime := time.Date(
		t.Year(),
		t.Month()+1,
		day,
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		t.Location(),
	)
	return newTime, nil
}

func nextTime(t time.Time, repeat string) (time.Time, error) {
	var nt time.Time

	dayRepeat := regexp.MustCompile(`^\d+d$`)
	weekRepeat := regexp.MustCompile(`^\d+w$`)
	monthRepeat := regexp.MustCompile(`^\d+m$`)

	dayOfMonthRepeat := regexp.MustCompile(`^\d+$`)

	switch {
	case dayRepeat.MatchString(repeat):
		return t.AddDate(0, 0, repeatNum(repeat)), nil
	case weekRepeat.MatchString(repeat):
		return t.AddDate(0, 0, 7*repeatNum(repeat)), nil
	case monthRepeat.MatchString(repeat):
		return t.AddDate(0, repeatNum(repeat), 0), nil
	case dayOfMonthRepeat.MatchString(repeat):
		next, err := nextDay(t, repeat)
		_ = err
		nt = next
		return nt, nil
	default:
		return t, errors.New("Unrecognized repeat string")
	}
}

type byDue []Task

func (a byDue) Len() int      { return len(a) }
func (a byDue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byDue) Less(i, j int) bool {
	if a[j].DueDate.IsZero() {
		return true
	}
	if a[i].DueDate.IsZero() && !a[j].DueDate.IsZero() {
		return false
	}

	return a[i].DueDate.Before(a[j].DueDate)
}
