package main

//Move represents a single move
type Move interface {
	getStart() int
	getDistance() int
	getPawn() Pawn
	//TODO: ADD FINAL SPOT METHOD AND USE EVERYWHERE
}
