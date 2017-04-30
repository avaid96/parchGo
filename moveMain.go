package main

//MoveMain implements Move
type MoveMain struct {
	pawn            Pawn
	start, distance int
}

func (m MoveMain) getStart() int {
	return m.start
}

func (m MoveMain) getDistance() int {
	return m.distance
}

func (m MoveMain) getPawn() Pawn {
	return m.pawn
}
