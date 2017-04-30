package main

//MoveHome implements Move
type MoveHome struct {
	pawn  Pawn
	start int
}

func (m MoveHome) getStart() int {
	return m.start
}

func (m MoveHome) getDistance() int {
	return _getPlayerHome(m.getPawn().getPlayerID()) - m.start
}

func (m MoveHome) getPawn() Pawn {
	return m.pawn
}
