package flextime

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"time"
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

func (tb *TimeBlock) Schedule(tCats ...TaskCategory) {
	for _, tc := range tCats {
		tb.Scheduled = append(tb.Scheduled, tc.Tasks...)
	}
	sort.Sort(byDue(tb.Scheduled))
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

func nextTime(t time.Time, repeat string) (time.Time, error) {
	dayRepeat := regexp.MustCompile(`^\d+d$`)
	weekRepeat := regexp.MustCompile(`^\d+w$`)
	monthRepeat := regexp.MustCompile(`^\d+m$`)

	switch {
	case dayRepeat.MatchString(repeat):
		return t.AddDate(0, 0, repeatNum(repeat)), nil
	case weekRepeat.MatchString(repeat):
		return t.AddDate(0, 0, 7*repeatNum(repeat)), nil
	case monthRepeat.MatchString(repeat):
		return t.AddDate(0, repeatNum(repeat), 0), nil
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
