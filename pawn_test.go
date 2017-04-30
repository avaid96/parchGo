package main

import "testing"

func TestPawn(t *testing.T) {
	p := Pawn{pawnID: 1, playerID: 1}
	gotID := p.getPawnID()
	if gotID != 1 {
		t.Error("Expected 1 got", gotID)
	}
	gotID = p.getPlayerID()
	if gotID != 1 {
		t.Error("Expected 1 got", gotID)
	}
}
