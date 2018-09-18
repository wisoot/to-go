package item

import "time"

type Item struct {
	Id          uint
	Description string
	IsFinished  bool
	CreatedAt   time.Time
	FinishedAt  time.Time
}

func (item *Item) Finish() {
	if item.IsFinished {
		return
	}

	item.IsFinished = true
	item.FinishedAt = time.Now()
}
