package plan

import (
	"strconv"
	"time"
)

var global_id uint64 = 0

type Plan struct {
	id      uint64
	start   time.Time
	end     time.Time
	content string
}

type PlanJson struct {
	Start   string `json:"start_time"`
	End     string `json:"end_time"`
	Content string `json:"content"`
}

func (p Plan) Id() uint64 {
	return p.id
}

func (p Plan) StartTime() time.Time {
	return p.start
}

func (p Plan) EndTime() time.Time {
	return p.end
}

func (p Plan) Content() string {
	return p.content
}

func (p *Plan) SetId(id uint64) {
	p.id = id
}

func (p *Plan) SetStartTime(start time.Time) {
	p.start = start
}

func (p *Plan) SetEndTime(end time.Time) {
	p.end = end
}

func (p *Plan) SetContent(content string) {
	p.content = content
}

func (p Plan) String() string {
	return strconv.FormatUint(uint64(p.id), 10) + " " + p.Content() + " " + p.start.String() + " ~ " + p.end.String()
}

// func (p *Plan) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(PlanJson{
// 		Start:   p.start.String(),
// 		End:     p.end.String(),
// 		Content: p.content,
// 	})
// }

// func (p *Plan) UnmarshalJSON(b []byte) error {
// 	temp := &PlanJson{}
// 	if err := json.Unmarshal(b, &temp); err != nil {
// 		return err
// 	}
// 	var err error
// 	p.start, err = newDate(temp.Start)
// 	if err != nil {
// 		return err
// 	}
// 	p.end, err = newDate(temp.End)
// 	if err != nil {
// 		return err
// 	}
// 	p.content = temp.Content
// 	return nil
// }
