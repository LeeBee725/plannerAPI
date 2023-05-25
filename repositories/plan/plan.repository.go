package plan

import (
	"errors"

	model "leebee.io/planner/rest/models/plan"
)

type PlanRepository interface {
	Save(p *model.Plan) (*model.Plan, error)
	FindAll() ([]model.Plan, error)
	FindById(id uint64) (*model.Plan, error)
	Delete(p *model.Plan) error
}

type MemoryPlanRepository struct {
	idSequence uint64
	plans      map[uint64]*model.Plan
}

func NewMemoryPlanRepository() *MemoryPlanRepository {
	mp := new(MemoryPlanRepository)
	mp.idSequence = 0
	mp.plans = map[uint64]*model.Plan{}
	return mp
}

func (mp *MemoryPlanRepository) Save(p *model.Plan) (*model.Plan, error) {
	if p.Id() == 0 {
		mp.idSequence++
		p.SetId(mp.idSequence)
	}
	mp.plans[uint64(p.Id())-1] = p
	return mp.plans[mp.idSequence-1], nil
}

func (mp *MemoryPlanRepository) FindAll() ([]model.Plan, error) {
	if mp.plans == nil {
		return nil, errors.New("the data is nil")
	}
	slice := make([]model.Plan, 0)
	for _, val := range mp.plans {
		slice = append(slice, *val)
	}
	return slice, nil
}

func (mp *MemoryPlanRepository) FindById(id uint64) (*model.Plan, error) {
	if _, exist := mp.plans[id]; !exist {
		return nil, errors.New("do not exist")
	}
	return mp.plans[id], nil
}

func (mp *MemoryPlanRepository) Delete(p *model.Plan) error {
	if _, exist := mp.plans[uint64(p.Id())]; !exist {
		return errors.New("do not exist")
	}
	delete(mp.plans, uint64(p.Id()))
	return nil
}
