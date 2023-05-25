package plan

import (
	"time"

	model "leebee.io/planner/rest/models/plan"
	repo "leebee.io/planner/rest/repositories/plan"
)

type PlanService interface {
	Create(start, end time.Time, content string) (*model.Plan, error)
	FindAll() ([]model.Plan, error)
	FindById(id uint64) (*model.Plan, error)
	Update(old, new *model.Plan) (*model.Plan, error)
}

type PlanServiceImpl struct {
	prepo repo.PlanRepository
}

func (ps *PlanServiceImpl) SetPlanRepository(r repo.PlanRepository) *PlanServiceImpl {
	ps.prepo = r
	return ps
}

func (ps *PlanServiceImpl) Create(start, end time.Time, content string) (*model.Plan, error) {
	p := new(model.Plan)
	p.SetStartTime(start)
	p.SetEndTime(end)
	p.SetContent(content)
	saved, err := ps.prepo.Save(p)
	return saved, err
}

func (ps *PlanServiceImpl) FindAll() ([]model.Plan, error) {
	plans, err := ps.prepo.FindAll()
	return plans, err
}

func (ps *PlanServiceImpl) FindById(id uint64) (*model.Plan, error) {
	plan, err := ps.prepo.FindById(id)
	return plan, err
}

func (ps *PlanServiceImpl) Update(old, new *model.Plan) (*model.Plan, error) {
	trg, err := ps.prepo.FindById(old.Id())
	if err != nil {
		return nil, err
	}
	trg.SetStartTime(new.StartTime())
	trg.SetEndTime(new.EndTime())
	trg.SetContent(new.Content())
	saved, err := ps.prepo.Save(trg)
	return saved, err
}
