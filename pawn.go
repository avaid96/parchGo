package main

type Pawn struct {
	pawnID int    `xml:"pawn>id"`
	color  string `xml:"pawn>color"`
}

var colorMap = map[string]int{
	"red":    0,
	"green":  1,
	"blue":   2,
	"yellow": 3,
}

var mapColor = map[int]string{
	0: "red",
	1: "green",
	2: "blue",
	3: "yellow",
}

func (p Pawn) getPawnID() int {
	return p.pawnID
}

func (p Pawn) getPlayerID() int {
	return colorMap[p.color]
}
