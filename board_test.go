package main

import "testing"

func TestBoardInit(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	// board := make([]Spot, 105)
	b1 := NewBoard(allPlayers)
	got := b1.GetNumPlayers()
	if got != 4 {
		t.Error("Expected 4 got", got)
	}
	got = b1.GetBoardSize()
	if got != 105 {
		t.Error("Expected 105 got", got)
	}
	//testing board contents
	player1Start := _getStartingSquare(0)
	player1StartSpot := b1.GetSpot(player1Start)
	got = len(player1StartSpot.GetPawns())
	if got != 4 {
		t.Error("Expected 4 got ", got)
	}
	startingPawns := player1StartSpot.GetPawns()
	for i := 0; i < 4; i++ {
		pawni := Pawn{i, 0}
		if _, ok := startingPawns[pawni]; !ok {
			t.Error(pawni, " not initialized in starting spot")
		}
	}
}

func TestStartingSquare(t *testing.T) {
	got := _getStartingSquare(0)
	if got != 101 {
		t.Error("Expected 101 got", got)
	}
	got = _getStartingSquare(1)
	if got != 102 {
		t.Error("Expected 102 got", got)
	}
}

func TestEnterSpot(t *testing.T) {
	got := _getEnterSpot(0)
	if got != 13 {
		t.Error("Expected 13 got", got)
	}
	got = _getEnterSpot(1)
	if got != 30 {
		t.Error("Expected 30 got", got)
	}
}

func TestEntering(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)
	firstPawn := Pawn{pawnID: 0, playerID: 0}
	m1 := EnterPiece{}
	b1.UpdateBoard(m1)
	//check pawn 0,0 moved out
	player1StartSpot := b1.GetSpot(_getStartingSquare(0))
	pawns := player1StartSpot.GetPawns()
	got := len(pawns)
	if got != 3 {
		t.Error("Expected 3 got ", got)
	}
	if _, ok := pawns[firstPawn]; ok {
		t.Error("Expect pawn to have been bounced")
	}
	//check pawn 0,0 moved to entering spot
	player1EnterSpot := b1.GetSpot(_getEnterSpot(0))
	pawns = player1EnterSpot.GetPawns()
	got = len(pawns)
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	if _, ok := pawns[firstPawn]; !ok {
		t.Error("Expect pawn to have come to the entering spot")
	}
}

func TestRegularMove(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)
	firstPawn := Pawn{pawnID: 0, playerID: 0}
	m1 := EnterPiece{firstPawn}
	b1.UpdateBoard(m1)
	m3 := MoveMain{firstPawn, _getEnterSpot(0), 5}
	b1.UpdateBoard(m3)
	finalSpot := b1.GetSpot(18)
	pawns := finalSpot.GetPawns()
	got := len(pawns)
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	if _, ok := pawns[firstPawn]; !ok {
		t.Error("Expect pawn to have come to the entering spot")
	}
}

func TestCheckHREnter(t *testing.T) {
	m3 := MoveMain{Pawn{pawnID: 0, playerID: 0}, 8, 3}
	if !(checkEnterHomeRow(m3)) {
		t.Error("Should be true, it enters home row")
	}
	m3 = MoveMain{Pawn{pawnID: 0, playerID: 0}, 7, 1}
	if checkEnterHomeRow(m3) {
		t.Error("Should be false, it's still on shared ring")
	}
}

func TestHREnter(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)
	firstPawn := Pawn{pawnID: 0, playerID: 0}
	b1.putPawn(firstPawn, 8)
	m3 := MoveMain{Pawn{pawnID: 0, playerID: 0}, 8, 3}
	b1.UpdateBoard(m3)
	wrongSpot := b1.GetSpot(11)
	pawns := wrongSpot.GetPawns()
	got := len(pawns)
	if got != 0 {
		t.Error("Expected 0 got ", got)
	}
	correctSpot := b1.GetSpot(68 + 3)
	pawns = correctSpot.GetPawns()
	got = len(pawns)
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	if _, ok := pawns[firstPawn]; !ok {
		t.Error("Expected pawn to have come to this spot")
	}
}

func TestGoHome(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)
	firstPawn := Pawn{pawnID: 0, playerID: 0}
	b1.putPawn(firstPawn, 71)
	m3 := MoveHome{Pawn{pawnID: 0, playerID: 0}, 71}
	b1.UpdateBoard(m3)
	correctSpot := b1.GetSpot(_getPlayerHome(firstPawn.getPlayerID()))
	pawns := correctSpot.GetPawns()
	got := len(pawns)
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	if _, ok := pawns[firstPawn]; !ok {
		t.Error("Expected pawn to have come home")
	}
}

func TestMoveInHomeRow(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)
	firstPawn := Pawn{pawnID: 0, playerID: 0}
	b1.putPawn(firstPawn, 71)
	m3 := MoveMain{Pawn{pawnID: 0, playerID: 0}, 71, 3}
	b1.UpdateBoard(m3)
	correctSpot := b1.GetSpot(74)
	pawns := correctSpot.GetPawns()
	got := len(pawns)
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	if _, ok := pawns[firstPawn]; !ok {
		t.Error("Expected pawn to have come home")
	}
}

func TestBop(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)
	firstPawn := Pawn{pawnID: 0, playerID: 0}
	m1 := EnterPiece{}
	b1.UpdateBoard(m1)
	//check pawn 0,0 moved out
	player1StartSpot := b1.GetSpot(_getStartingSquare(0))
	pawns := player1StartSpot.GetPawns()
	got := len(pawns)
	if got != 3 {
		t.Error("Expected 3 got ", got)
	}
	b1.putPawn(firstPawn, 21)
	secondPawn := Pawn{pawnID: 0, playerID: 1}
	b1.putPawn(secondPawn, 19)
	m3 := MoveMain{secondPawn, 19, 2}
	b1.UpdateBoard(m3)
	pawns = b1.GetSpot(21).GetPawns()
	got = len(pawns)
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	if _, ok := pawns[secondPawn]; !ok {
		t.Error("Expected second pawn to have moved in")
	}
	pawns = b1.GetSpot(_getStartingSquare(firstPawn.getPlayerID())).GetPawns()
	got = len(pawns)
	if got != 4 {
		t.Error("Expected 1 got ", got)
	}
}

func TestSafety(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)
	firstPawn := Pawn{pawnID: 0, playerID: 0}
	m1 := EnterPiece{}
	b1.UpdateBoard(m1)
	//check pawn 0,0 moved out
	player1StartSpot := b1.GetSpot(_getStartingSquare(0))
	pawns := player1StartSpot.GetPawns()
	got := len(pawns)
	if got != 3 {
		t.Error("Expected 3 got ", got)
	}
	b1.putPawn(firstPawn, 20)
	secondPawn := Pawn{pawnID: 0, playerID: 1}
	b1.putPawn(secondPawn, 19)
	m3 := MoveMain{secondPawn, 19, 1}
	errCode := b1.UpdateBoard(m3)
	if errCode != BLOCKADE {
		t.Error("Expected safety single pawn blockade")
	}
	pawns = b1.GetSpot(20).GetPawns()
	got = len(pawns)
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	if _, ok := pawns[firstPawn]; !ok {
		t.Error("Expected second pawn to have moved in")
	}
}

func TestBlockade(t *testing.T) {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)
	firstPawn := Pawn{pawnID: 0, playerID: 0}
	b1.putPawn(firstPawn, 20)
	secondPawn := Pawn{pawnID: 1, playerID: 0}
	b1.putPawn(secondPawn, 20)
	thirdPawn := Pawn{pawnID: 2, playerID: 2}
	m3 := MoveMain{thirdPawn, 19, 1}
	errCode := b1.UpdateBoard(m3)
	if errCode != BLOCKADE {
		t.Error("Expected blockade")
	}
}
