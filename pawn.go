package main

type Pawn struct {
	pawnID, playerID int
}

func (p Pawn) getPawnID() int {
	return p.pawnID
}

func (p Pawn) getPlayerID() int {
	return p.playerID
}
