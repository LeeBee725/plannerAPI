package plan

import (
	"encoding/json"
	"strconv"
	"time"
)

var global_id uint32 = 0

type Plan struct {
	id      uint32
	start   time.Time
	end     time.Time
	content string
}

type PlanJson struct {
	Start   string `json:"start_time"`
	End     string `json:"end_time"`
	Content string `json:"content"`
}

func NewPlan() *Plan {
	var p Plan
	global_id++
	p.id = global_id
	return &p
}

func newPlanParam(start, end time.Time, content string) (p *Plan) {
	p = NewPlan()
	p.SetStartTime(start.Year(), start.Month(), start.Day(),
		start.Hour(), start.Minute(), start.Second(), start.Nanosecond(), start.Location())
	p.SetEndTime(end.Year(), end.Month(), end.Day(),
		end.Hour(), end.Minute(), end.Second(), end.Nanosecond(), end.Location())
	p.SetContent(content)
	return p
}

func NewPlanParam(start, end, content string) (p *Plan) {
	startDate, err := newDate(start)
	if err != nil {
		panic(err)
	}
	endDate, err := newDate(end)
	if err != nil {
		panic(err)
	}
	p = newPlanParam(startDate, endDate, content)
	return p
}

func newDate(dateString string) (t time.Time, err error) {
	if dateString == "" {
		return t, err
	}
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return t, err
	}
	t, err = time.ParseInLocation(time.RFC3339, dateString, loc)
	return
}

func (p Plan) Id() uint32 {
	return p.id
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

func (p Plan) String() string {
	return strconv.FormatUint(uint64(p.id), 10) + " " + p.Content() + " " + p.start.String() + " ~ " + p.end.String()
}

func (p *Plan) MarshalJSON() ([]byte, error) {
	return json.Marshal(PlanJson{
		Start:   p.StartTime(),
		End:     p.EndTime(),
		Content: p.content,
	})
}

func (p *Plan) UnmarshalJSON(b []byte) error {
	temp := &PlanJson{}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	var err error
	p.start, err = newDate(temp.Start)
	if err != nil {
		return err
	}
	p.end, err = newDate(temp.End)
	if err != nil {
		return err
	}
	p.content = temp.Content
	return nil
}

func (p *Plan) Update(newP Plan) {
	var temp Plan
	if temp.start != newP.start {
		p.start = newP.start
	}
	if temp.end != newP.end {
		p.end = newP.end
	}
	if temp.content != newP.content {
		p.content = newP.content
	}
}
