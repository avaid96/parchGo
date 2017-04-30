package main

//EnterPiece implements Move
type EnterPiece struct {
	pawn Pawn
}

func (m EnterPiece) getStart() int {
	return _getStartingSquare(m.pawn.getPlayerID())
}

func (m EnterPiece) getDistance() int {
	return 5
}

func (m EnterPiece) getPawn() Pawn {
	return m.pawn
}
