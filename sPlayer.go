package main

import (
	"fmt"
	"reflect"
)

type SPlayer struct {
	color string
	name  string
	id    int
}

func (sp *SPlayer) startGame(color string) string {
	sp.name = "jimmy"
	sp.color = color
	return sp.name
}

func (sp SPlayer) doMove(board Board, rolls []int) []Move {
	t1, _ := NewTurn(board, sp.id, rolls)
	m1 := t1.makeMove(0, Pawn{0, mapColor[sp.id]})
	m2 := t1.makeMove(1, Pawn{1, mapColor[sp.id]})
	fmt.Println(reflect.TypeOf(m1))
	fmt.Println(reflect.TypeOf(m2))
	return []Move{m1, m2}
}
