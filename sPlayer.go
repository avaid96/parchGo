package main

import (
	"fmt"
	"log"
	"reflect"
)

type SPlayer struct {
	color string
	name  string
}

func (sp *SPlayer) startGame(color string) string {
	sp.name = "jimmy"
	sp.color = color
	return sp.name
}

func (sp SPlayer) doMove(board Board, rolls []int) []Move {
	t1, _ := NewTurn(board, colorMap[sp.name], rolls)
	m1 := t1.makeMove(0, Pawn{0, sp.color})
	m2 := t1.makeMove(1, Pawn{1, sp.color})
	fmt.Println(reflect.TypeOf(m1))
	fmt.Println(reflect.TypeOf(m2))
	return []Move{m1, m2}
}

func (sp *SPlayer) DoublesPenalty() {
	log.Println("!!!!!!!! not implemented : doubles penalty")
}
