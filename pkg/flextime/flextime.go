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

func parseDateChange(dateModifier string)

func (t Task) Next() Task {
	dayRepeat := regexp.MustCompile(`^\d+d$`)

	switch {
	case dayRepeat.MatchString(t.Repeat):
		numStr := t.Repeat[:len(t.Repeat)-1]
		num, err := strconv.ParseInt(numStr, 10, 64)
		_ = err

		return Task{
			DueDate: t.DueDate.AddDate(0, 0, int(num)),
			Repeat:  t.Repeat,
		}
	default:
		return t
	}
}

type TimeBlock struct{}
