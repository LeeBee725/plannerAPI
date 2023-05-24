package plan

import (
	"errors"

	"leebee.io/planner/rest/models/plan"
)

type PlanRepository interface {
	Save(p *plan.Plan) (*plan.Plan, error)
	FindAll() ([]plan.Plan, error)
	FindById(id uint64) (*plan.Plan, error)
	Delete(p *plan.Plan) error
}

type MemoryPlanRepository struct {
	idSequence uint64
	plans      map[uint64]*plan.Plan
}

func NewMemoryPlanRepository() *MemoryPlanRepository {
	mp := new(MemoryPlanRepository)
	mp.idSequence = 0
	mp.plans = map[uint64]*plan.Plan{}
	return mp
}

func (mp *MemoryPlanRepository) Save(p *plan.Plan) (*plan.Plan, error) {

	mp.idSequence++
	// p.
	mp.plans[mp.idSequence]
	return mp.plans[mp.idSequence-1], nil
}

func (mp *MemoryPlanRepository) FindAll() ([]plan.Plan, error) {
	if mp.plans == nil {
		return nil, errors.New("the data is nil")
	}
	slice := make([]plan.Plan, 0)
	for _, val := range mp.plans {
		slice = append(slice, *val)
	}
	return slice, nil
}

func (mp *MemoryPlanRepository) FindById(id uint64) (*plan.Plan, error) {
	if id > uint64(len(mp.plans)) {
		return nil, errors.New("the id is out of range")
	}
	return &mp.plans[id], nil
}

func (mp *MemoryPlanRepository) Delete(p *plan.Plan) error {

}
