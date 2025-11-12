package repository

import (
	"github.com/zamyatin-zkex/estate_calc_bot/internal/entity"
	"sync"
)

type State struct {
	mx sync.Mutex

	accounts map[string]entity.State
}

func NewState() *State {
	return &State{
		accounts: make(map[string]entity.State),
	}
}

func (s *State) Set(account string, state entity.State) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.accounts[account] = state
}

func (s *State) Get(account string) entity.State {
	s.mx.Lock()
	defer s.mx.Unlock()
	
	return s.accounts[account]
}
