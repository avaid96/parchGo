package main

import "reflect"

type Board struct {
	allPlayers []int //implemented as a map of pawnIDs to pawns
	board      []Spot
}

//Board constants
const (
	STARTINGSQUARESOFFSET = 101
	ENTERSPOT             = 13
	HOMEROWSTART          = 8
	PERPLAYEROFFSET       = 17
	MAINRINGSIZE          = PERPLAYEROFFSET * 4
	STARTOFHOMEROWS       = 68
	HOMEROWLENGTH         = 7
	STARTOFHOMES          = 97
	SAFETYOFFSET1         = 3
	SAFETYOFFSET2         = 8
	SAFETYOFFSET3         = 13
)

//Return codes for update board
const (
	OK              = 0
	BLOCKADE        = -1
	BOP             = -2
	INVALIDMOVEHOME = -3
)

func (b Board) GetBoardSize() int {
	return len(b.board)
}

func (b Board) GetNumPlayers() int {
	return len(b.allPlayers)
}

func (b Board) GetSpot(spotIdx int) Spot {
	return b.board[spotIdx]
}

func (b Board) putPawn(p Pawn, pos int) {
	b.board[pos].Add(p)
}

const PAWNNOTFOUND = -1

func (b Board) findPawn(p Pawn) int {
	for spotIdx, spot := range b.board {
		if spot.PawnOnSpot(p) {
			return spotIdx
		}
	}
	return PAWNNOTFOUND
}

func NewBoard(allPlayers []int) *Board {
	board := make([]Spot, 105)
	//initialize spots
	for index := range board {
		board[index] = Spot{make(map[Pawn]bool)}
	}
	//populate starting spots
	for _, playerID := range allPlayers {
		startingSqaure := _getStartingSquare(playerID)
		//initializing player pawns
		for i := 0; i < 4; i++ {
			pawni := Pawn{i, mapColor[i]}
			board[startingSqaure].Add(pawni)
		}
	}
	return &Board{allPlayers: allPlayers, board: board}
}

func NewEmptyBoard() *Board {
	board := make([]Spot, 105)
	//initialize spots
	for index := range board {
		board[index] = Spot{make(map[Pawn]bool)}
	}
	return &Board{board: board}
}

func (b Board) UpdateBoard(m Move) int {
	if reflect.TypeOf(m) == reflect.TypeOf(EnterPiece{}) {
		enterSpotIdx := _getEnterSpot(m.getPawn().getPlayerID())
		if b.checkBlockade(m) {
			//Doing nothing - returning -1 as error code
			return BLOCKADE
		}
		if b.checkBop(m) {
			//get pawn on spot
			boppedPawn := b.board[enterSpotIdx].GetPawns()
			//need loop to access map
			for pawn := range boppedPawn {
				//remove pawn from the spot
				b.board[enterSpotIdx].Remove(pawn)
				//add boppedPawn to starting spot
				b.board[_getStartingSquare(pawn.getPlayerID())].Add(pawn)
				break //it is guaranteed to be a one element loop
			}
			b.board[m.getStart()].Remove(m.getPawn())
			b.board[enterSpotIdx].Add(m.getPawn())
			return BOP
		}
		//perform entering move
		//remove pawn from the starting spot
		//@TODO: ERROR IF PAWN NOT AT SPOT
		b.board[m.getStart()].Remove(m.getPawn())
		//put the pawn in the enter spot
		b.board[enterSpotIdx].Add(m.getPawn())
		return OK
	}
	if reflect.TypeOf(m) == reflect.TypeOf(MoveMain{}) {
		if checkEnterHomeRow(m) {
			//remove pawn from the original spot
			b.board[m.getStart()].Remove(m.getPawn())
			//calculate new spot
			distTillEntry := _getHomeRowEnterSpot(m.getPawn().getPlayerID()) - m.getStart()
			distAfterEntry := m.getDistance() - distTillEntry
			//put the pawn in the new spot
			newSpotIdx := _getPlayerHomeRowStart(m.getPawn().getPlayerID()) + distAfterEntry
			b.board[newSpotIdx].Add(m.getPawn())
			return OK
		}
		if checkInHomeRow(m) {
			//remove pawn from the original spot
			b.board[m.getStart()].Remove(m.getPawn())
			//put the pawn in the new spot
			b.board[m.getStart()+m.getDistance()].Add(m.getPawn())
			//TODO: CHECK IT DOESN'T GO HOME AND CHECK FOR BLOCKADES
			return OK
		}
		//REGULAR MAIN RING MOVE
		finalSpot := ((m.getStart() + m.getDistance()) % MAINRINGSIZE)
		if b.checkBlockade(m) {
			//Doing nothing - returning -1 as error code
			return BLOCKADE
		}
		if b.checkBop(m) {
			//get pawn on spot
			boppedPawn := b.board[finalSpot].GetPawns()
			//need loop to access map
			for pawn := range boppedPawn {
				//remove pawn from the spot
				b.board[finalSpot].Remove(pawn)
				//add boppedPawn to starting spot
				b.board[_getStartingSquare(pawn.getPlayerID())].Add(pawn)
				break //it is guaranteed to be a one element loop
			}
			b.board[m.getStart()].Remove(m.getPawn())
			b.board[finalSpot].Add(m.getPawn())
			return BOP
		}
		//remove pawn from the original spot
		b.board[m.getStart()].Remove(m.getPawn())
		//put the pawn in the new spot
		b.board[finalSpot].Add(m.getPawn())
		return OK
	}
	if reflect.TypeOf(m) == reflect.TypeOf(MoveHome{}) {
		finalPos := m.getStart() + m.getDistance()
		isValid := (finalPos == _getPlayerHome(m.getPawn().getPlayerID()))
		//TODO: CHECK IT DOESN'T GO HOME AND CHECK FOR BLOCKADES
		if isValid {
			//remove pawn from the original spot
			b.board[m.getStart()].Remove(m.getPawn())
			//put the pawn in the new spot
			b.board[finalPos].Add(m.getPawn())
			return OK
		}
		return INVALIDMOVEHOME
	}
	return OK
}

func checkEnterHomeRow(m Move) bool {
	playerID := m.getPawn().getPlayerID()
	homeRowEntrySpot := _getHomeRowEnterSpot(playerID)
	finalSpot := m.getStart() + m.getDistance()
	if m.getStart() <= homeRowEntrySpot {
		if finalSpot > homeRowEntrySpot {
			return true
		}
	}
	return false
}

func checkInHomeRow(m Move) bool {
	playerID := m.getPawn().getPlayerID()
	homeRowStart := _getPlayerHomeRowStart(playerID)
	playerHome := _getPlayerHome(playerID)
	if m.getStart() >= homeRowStart && m.getStart() < playerHome {
		return true
	}
	return false
}

func (b Board) checkBop(m Move) bool {
	var finalSpot int
	playerID := m.getPawn().getPlayerID()
	if reflect.TypeOf(m) == reflect.TypeOf(EnterPiece{}) {
		finalSpot := _getEnterSpot(m.getPawn().getPlayerID())
		pawnsAtDest := b.GetSpot(finalSpot).GetPawns()
		if len(pawnsAtDest) == 1 {
			for pawn := range pawnsAtDest {
				if pawn.getPlayerID() == playerID {
					//it will be able to make a blockade at that spot
					return false
				}
			}
			return true
		}
	}
	finalSpot = (m.getDistance() + m.getStart()) % 68
	pawnsAtDest := b.GetSpot(finalSpot).GetPawns()

	if len(pawnsAtDest) == 0 {
		return false
	}
	if len(pawnsAtDest) == 1 {
		for pawn := range pawnsAtDest {
			if pawn.getPlayerID() == playerID {
				//it will be able to make a blockade at that spot
				return false
			}
		}
	}
	if len(pawnsAtDest) > 1 {
		//check blockade called earlier - shouldn't get here if blockade
		panic("There are too many pawns at this spot")
	}
	return true
}

func (b Board) isSafety(spot int) bool {
	for wingIdx := 0; wingIdx < b.GetNumPlayers(); wingIdx++ {
		safety1 := PERPLAYEROFFSET*wingIdx + SAFETYOFFSET1
		safety2 := PERPLAYEROFFSET*wingIdx + SAFETYOFFSET2
		safety3 := PERPLAYEROFFSET*wingIdx + SAFETYOFFSET3
		if spot == safety1 || spot == safety2 || spot == safety3 {
			return true
		}
	}
	return false
}

func (b Board) checkBlockade(m Move) bool {
	route, finalSpot := b.getRoute(m)
	for _, spot := range route {
		pawnsAtDest := spot.GetPawns()
		if len(pawnsAtDest) == 2 {
			//TODO: CHECK THAT THEY BELONG TO SAME PLAYER
			return true
		}
	}
	if b.isSafety(finalSpot) {
		pawnsAtDest := b.GetSpot(finalSpot).GetPawns()
		if len(pawnsAtDest) == 1 {
			if reflect.TypeOf(m) == reflect.TypeOf(EnterPiece{}) {
				return false
			}
			return true
		}
	}
	return false
}

func (b Board) getRoute(m Move) ([]Spot, int) {
	start := m.getStart()
	playerID := m.getPawn().getPlayerID()
	dist := m.getDistance()
	var spots []Spot
	if reflect.TypeOf(m) == reflect.TypeOf(EnterPiece{}) {
		finalSpot := b.GetSpot(_getEnterSpot(m.getPawn().getPlayerID()))
		spots = append(spots, finalSpot)
		return spots, _getEnterSpot(m.getPawn().getPlayerID())
	}
	if checkEnterHomeRow(m) {
		homeRowEntrySpot := _getHomeRowEnterSpot(playerID)
		distTillEntry := homeRowEntrySpot - start
		distAfterEntry := dist - distTillEntry
		finalSpot := homeRowEntrySpot + distAfterEntry
		step := 1
		for distTillEntry >= 0 {
			nowSpot := b.GetSpot(start + step)
			spots = append(spots, nowSpot)
			step++
			distTillEntry--
		}
		step = 1
		for distAfterEntry >= 0 {
			nowSpot := b.GetSpot(homeRowEntrySpot + step)
			spots = append(spots, nowSpot)
			step++
			distAfterEntry--
		}
		return spots, finalSpot
	}
	if checkInHomeRow(m) {
		distToGo := dist
		finalSpot := start + distToGo
		for distToGo >= 0 {
			nowSpot := b.GetSpot(finalSpot)
			spots = append(spots, nowSpot)
			distToGo--
		}
		return spots, finalSpot
	}
	//normal
	distToGo := dist
	finalSpot := (start + distToGo) % MAINRINGSIZE
	for distToGo >= 0 {
		nowSpot := b.GetSpot(finalSpot)
		spots = append(spots, nowSpot)
		distToGo--
	}
	return spots, finalSpot
}

func _getStartingSquare(playerID int) int {
	return STARTINGSQUARESOFFSET + playerID
}

func _getEnterSpot(playerID int) int {
	return PERPLAYEROFFSET*playerID + ENTERSPOT
}

func _getHomeRowEnterSpot(playerID int) int {
	return PERPLAYEROFFSET*playerID + HOMEROWSTART
}

func _getPlayerHomeRowStart(playerID int) int {
	return STARTOFHOMEROWS + playerID*HOMEROWLENGTH
}

func _getPlayerHome(playerID int) int {
	return STARTOFHOMES + playerID
}

// func main() {
// 	//Simulates server actions
// 	// allPlayers := []int{0, 1, 2, 3}
// 	// b1 := NewBoard(allPlayers)
// 	// m1 := EnterPiece{Pawn{pawnID: 0, playerID: 0}}
// 	// b1.UpdateBoard(m1)
// 	// m2 := EnterPiece{Pawn{pawnID: 1, playerID: 0}}
// 	// b1.UpdateBoard(m2)
// 	// m3 := MoveMain{Pawn{pawnID: 0, playerID: 0}, 8, 3}
// 	// b1.UpdateBoard(m3)

// }
