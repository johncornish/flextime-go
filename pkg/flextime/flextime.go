package flextime

import (
	"regexp"
	"strconv"
	"time"
)

type Task struct {
	DueDate time.Time
	Repeat  string
}

func (t Task) IsDue() bool {
	return t.DueDate.Before(time.Now().Local())
}

func nextTime(t time.Time, repeat string) time.Time {
	num, err := strconv.ParseInt(repeat[:len(repeat)-1], 10, 64)
	_ = err

	dayRepeat := regexp.MustCompile(`^\d+d$`)
	weekRepeat := regexp.MustCompile(`^\d+w$`)
	monthRepeat := regexp.MustCompile(`^\d+m$`)

	switch {
	case dayRepeat.MatchString(repeat):
		return t.AddDate(0, 0, int(num))
	case weekRepeat.MatchString(repeat):
		return t.AddDate(0, 0, 7*int(num))
	case monthRepeat.MatchString(repeat):
		return t.AddDate(0, int(num), 0)
	default:
		return t
	}
}

func (t Task) Next() Task {
	return Task{
		DueDate: nextTime(t.DueDate, t.Repeat),
		Repeat:  t.Repeat,
	}
}

type TimeBlock struct{}
