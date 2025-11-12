package entity

import "strings"

type State string

func (s State) String() string {
	return string(s)
}

func (s State) Nil() bool {
	return s == ""
}

const Root = State("root")
const PlanTotalCost = State("plan_total_cost")
const BankRates = State("bank_rates")

var states = map[State]State{
	Root:          Root,
	PlanTotalCost: PlanTotalCost,
	BankRates:     BankRates,
}

func (s State) Parse() State {
	return states[State(strings.Trim(string(s), " \t\r\n/"))]
}
