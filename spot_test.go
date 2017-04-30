package main

import "testing"

func TestSpot(t *testing.T) {
	pawn1 := Pawn{0, 1}
	pawn2 := Pawn{0, 2}
	pawn3 := Pawn{0, 3}
	s := Spot{make(map[Pawn]bool)}
	got := len(s.GetPawns())
	if got != 0 {
		t.Error("Expected 0 got ", got)
	}
	//Adding pawn1
	err := s.Add(pawn1)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	got = len(s.GetPawns())
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	gotPawn := s.GetPawns()[pawn1]
	if gotPawn != true {
		t.Error("Expected pawn1")
	}
	//Adding pawn2
	err = s.Add(pawn2)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	got = len(s.GetPawns())
	if got != 2 {
		t.Error("Expected 2 got ", got)
	}
	gotPawn = s.GetPawns()[pawn2]
	if gotPawn != true {
		t.Error("Expected pawn2 ", got)
	}
	// THIS SHOULD BE ADDED IF WE MAKE A DIFFERENT REGULAR AND STARTING/HOME SPOT
	// //Adding pawn3
	// err = s.Add(pawn3)
	// if err == nil {
	// 	t.Error("Expected \"Spot already full\", got ", nil)
	// }
	//Removing pawn2
	s.Remove(pawn2)
	got = len(s.GetPawns())
	if got != 1 {
		t.Error("Expected 1 got ", got)
	}
	_, ok := s.GetPawns()[pawn2]
	if ok != false {
		t.Error("Expected pawn2 to have been removed")
	}
	//Adding pawn3
	err = s.Add(pawn3)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	got = len(s.GetPawns())
	if got != 2 {
		t.Error("Expected 2 got ", got)
	}
	gotPawn = s.GetPawns()[pawn3]
	if gotPawn != true {
		t.Error("Expected pawn3 ", got)
	}
}
