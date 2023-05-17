package plan

import (
	"time"
)

var global_id uint = 0

type Plan struct {
	id      uint
	start   time.Time
	end     time.Time
	content string
}

func NewPlan() *Plan {
	var p Plan
	p.id = global_id + 1
	global_id++
	return &p
}

func (p Plan) StartTime() string {
	return p.start.String()
}

func (p Plan) EndTime() string {
	return p.end.String()
}

func (p Plan) Content() string {
	return p.content
}

func (p *Plan) SetStartTime(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) {
	p.start = time.Date(year, month, day, hour, min, sec, nsec, loc)
}

func (p *Plan) SetEndTime(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) {
	p.end = time.Date(year, month, day, hour, min, sec, nsec, loc)
}

func (p *Plan) SetContent(content string) {
	p.content = content
}
