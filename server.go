package main

import "fmt"
import "reflect"

func main() {
	allPlayers := []int{0, 1, 2, 3}
	b1 := NewBoard(allPlayers)

	for player := range allPlayers {
		t1, rCode := NewTurn(*b1, player)
		fmt.Println("your rolls are: ", t1.rolls)

		//TAKE THE TURN HERE
		m1 := (t1.makeMove(0, Pawn{0, player}))
		fmt.Println(reflect.TypeOf(m1))
		fmt.Println(m1)
		rc := b1.UpdateBoard(m1)
		fmt.Println(rc)
		//TAKE THE NEXT TURN HERE
		m1 = (t1.makeMove(1, Pawn{1, player}))
		fmt.Println(reflect.TypeOf(m1))
		fmt.Println(m1)
		rc = b1.UpdateBoard(m1)
		fmt.Println(rc)

		//used all rolls?
		for _, roll := range t1.rolls {
			if roll != -1 {
				panic("Not used all rolls")
			}
		}

		numDoubles := 0
		for numDoubles < 3 {
			if rCode == DOUBLESCODE {
				fmt.Println("You have a doubles bonus")
				if numDoubles == 3 {
					//IF THIRD - NO NEW TURN - MOVE ALL PAWNS HOME
					panic("Unimplemented")
				} else {
					//ELSE - GIVE NEW TURN
					t2, rC := NewTurn(*b1, player)
					fmt.Println(t2.rolls)
					rCode = rC

					//TAKE THE TURN HERE
					m1 := (t1.makeMove(0, Pawn{0, player}))
					fmt.Println(reflect.TypeOf(m1))
					fmt.Println(m1)
					rc := b1.UpdateBoard(m1)
					fmt.Println(rc)
					//TAKE THE NEXT TURN HERE
					m1 = (t1.makeMove(1, Pawn{1, player}))
					fmt.Println(reflect.TypeOf(m1))
					fmt.Println(m1)
					rc = b1.UpdateBoard(m1)
					fmt.Println(rc)

					//used all rolls?
					for _, roll := range t1.rolls {
						if roll != -1 {
							panic("Not used all rolls")
						}
					}
				}
				numDoubles++
			} else {
				break
			}
		}

		break
	}

	// t1, rCode := NewTurn(*b1, 0)
	// fmt.Println(t1.rolls)
	// fmt.Println(t1.doublesP)
	// fmt.Println(rCode)

	// p1 := Pawn{playerID: 0, pawnID: 2}
	// t1.makeMove(0, p1)
	// fmt.Println(t1.rolls)
}
