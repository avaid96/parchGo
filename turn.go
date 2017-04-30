package main

type Turn struct {
	rolls    []int
	doublesP bool
	board    Board
	playerID int
}

//new turn return codes
const (
	DOUBLESCODE = 1
	FINE        = 0
)

//creates a new turn
func NewTurn(b Board, playerID int) (*Turn, int) {
	r1 := throwDie()
	r2 := throwDie()
	rolls := []int{r1, r2}
	doublesP := false
	if r1 == r2 {
		doublesP = true
		return &Turn{rolls: rolls, doublesP: doublesP, board: b, playerID: playerID}, DOUBLESCODE
	}
	return &Turn{rolls: rolls, doublesP: doublesP, board: b, playerID: playerID}, FINE
}

func (t Turn) makeMove(whichRoll int, whichPawn Pawn) Move {
	rollValue := t.rolls[whichRoll]
	t.rolls[whichRoll] = -1

	pawnLoc := (t.board.findPawn(whichPawn))
	if pawnLoc == _getStartingSquare(t.playerID) {
		return EnterPiece{whichPawn}
	}
	if pawnLoc+rollValue == _getPlayerHome(t.playerID) {
		return MoveHome{whichPawn, pawnLoc}
	}
	return MoveMain{whichPawn, pawnLoc, rollValue}
}
