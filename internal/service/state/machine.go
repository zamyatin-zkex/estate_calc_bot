package state

import "github.com/zamyatin-zkex/estate_calc_bot/internal/entity"

type Machine struct {
	Moves map[entity.State][]entity.State
}

func NewMachine() Machine {
	return Machine{
		Moves: map[entity.State][]entity.State{
			entity.Root:          {entity.Root, entity.PlanTotalCost},
			entity.PlanTotalCost: {entity.Root, entity.PlanTotalCost},
		},
	}
}
