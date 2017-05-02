package main

import "fmt"

func main() {
	colorOrd := []string{"red", "green", "blue", "yellow"}
	//player initialize
	allPlayerIDs := []int{0, 1, 2, 3}
	var allPlayers []SPlayer
	for i := range allPlayerIDs {
		playerI := SPlayer{}
		_ = playerI.startGame(colorOrd[i])
		allPlayers = append(allPlayers, playerI)
	}
	b1 := NewBoard(allPlayerIDs)

	// let's go
	for player := range allPlayerIDs {
		r1 := throwDie()
		r2 := throwDie()
		rollArr := []int{r1, r2}
		moves := allPlayers[player].doMove(*b1, rollArr)
		fmt.Println(moves)

		//used all rolls?
		for _, roll := range rollArr {
			exists := false
			for _, move := range moves {
				if move.getDistance() != roll {
					exists = true
				}
			}
			if !exists {
				panic("Roll not used")
			}
		}
	}
	// //BELOW IS DOUBLE ROLL IMPLEMENTATION
	// numDoubles := 0
	// for numDoubles < 3 {
	// 	if rCode == DOUBLESCODE {
	// 		fmt.Println("You have a doubles bonus")
	// 		if numDoubles == 3 {
	// 			//IF THIRD - NO NEW TURN - MOVE ALL PAWNS HOME
	// 			panic("Unimplemented")
	// 		} else {
	// 			//ELSE - GIVE NEW TURN
	// 			t2, rC := NewTurn(*b1, player)
	// 			fmt.Println(t2.rolls)
	// 			rCode = rC

	// 			//TAKE THE TURN HERE
	// 			m1 := (t1.makeMove(0, Pawn{0, player}))
	// 			fmt.Println(reflect.TypeOf(m1))
	// 			fmt.Println(m1)
	// 			rc := b1.UpdateBoard(m1)
	// 			fmt.Println(rc)
	// 			//TAKE THE NEXT TURN HERE
	// 			m1 = (t1.makeMove(1, Pawn{1, player}))
	// 			fmt.Println(reflect.TypeOf(m1))
	// 			fmt.Println(m1)
	// 			rc = b1.UpdateBoard(m1)
	// 			fmt.Println(rc)

	// 			//used all rolls?
	// 			for _, roll := range t1.rolls {
	// 				if roll != -1 {
	// 					panic("Not used all rolls")
	// 				}
	// 			}
	// 		}
	// 		numDoubles++
	// 	} else {
	// 		break
	// 	}
}
