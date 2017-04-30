package main

import (
	"errors"
)

type Spot struct {
	pawnsHere map[Pawn]bool
}

func (s *Spot) Add(p Pawn) error {
	if len(s.GetPawns()) < 4 {
		s.pawnsHere[p] = true
		return nil
	}
	return errors.New("Spot already full")
}

func (s Spot) GetPawns() map[Pawn]bool {
	return s.pawnsHere
}

func (s *Spot) Remove(p Pawn) {
	//design choice: no error return if p isn't in the spot to start with
	delete(s.pawnsHere, p)
}
