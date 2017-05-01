package main

import (
	"bytes"
	"encoding/xml"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"fmt"

	"strings"

	"github.com/gorilla/mux"
)

//init player
var player MyPlayer

type Request struct {
	XMLName xml.Name `xml:"Message"`
	Color   string   `xml:"start-game"`
	Die     string   `xml:"do-move>die"`
}

type PawnIntermediateStr struct {
	Color string `xml:"color"`
	ID    string `xml:"id"`
}

func interpretPawns(s string) []Pawn {
	pawns := make([]Pawn, 0)
	re, _ := regexp.Compile("<pawn>((.|\\s) |[^(pawn)])*pawn>")
	PawnsStr := re.FindAllStringSubmatch(s, -1)
	for _, pawn := range PawnsStr {
		v := PawnIntermediateStr{}
		xml.Unmarshal([]byte(pawn[0]), &v)
		i, err := strconv.Atoi(strings.TrimSpace(v.ID))
		if err != nil {
			fmt.Println(err)
		}
		thisPawn := Pawn{pawnID: i, color: v.Color}
		pawns = append(pawns, thisPawn)
	}
	return pawns
}

func interpret(w http.ResponseWriter, r *http.Request) {
	// get request body as string
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()
	// encode as xml
	v := Request{}
	err := xml.Unmarshal([]byte(s), &v)
	if err != nil {
		panic(err)
	}
	log.Println("Message received-----------------")
	// log.Println(s)
	log.Println(v)
	//branch off
	if strings.Contains(s, "start-game") {
		log.Println("this tells me to start a game")
		name := player.startGame(v.Color)
		response := fmt.Sprintf("<name>%s</name>", name)
		w.Write([]byte(response))
		log.Println("I have started a game told the server my name and am done")
	}
	if strings.Contains(s, "do-move") {
		b1 := NewEmptyBoard()
		log.Println("this tells me to do a move")
		re, _ := regexp.Compile("<board>(.|\\s)*board>")
		boardStrings := re.FindStringSubmatch(s)
		boardString := boardStrings[0]
		re, _ = regexp.Compile("<die>\\s*\\d\\s*</die>")
		dieStrings := (re.FindAllStringSubmatch(s, -1))
		re, _ = regexp.Compile("\\d")
		dieRoll1 := (re.FindStringSubmatch(dieStrings[0][0]))[0]
		dieRoll2 := (re.FindStringSubmatch(dieStrings[1][0]))[0]
		log.Println("dierolls: ", dieRoll1, dieRoll2)
		return
		log.Println("!!!!!!!Start")
		//building the start
		re, _ = regexp.Compile("<start>(.|\\s)*start>")
		startStrings := re.FindStringSubmatch(boardString)
		startString := startStrings[0]
		startPawns := interpretPawns(startString)
		var startingSpot Spot
		for _, pawn := range startPawns {
			pawnColor := strings.TrimSpace(pawn.color)
			playerID := colorMap[pawnColor]
			startingSpot = b1.GetSpot(_getStartingSquare(playerID))
			startingSpot.Add(pawn)
		}
		log.Println(b1)
		log.Println("!!!!!!!Main")
		//building the main
		re, _ = regexp.Compile("<main>(.|\\s)*main>")
		mainStrings := re.FindStringSubmatch(boardString)
		mainString := mainStrings[0]
		re, _ = regexp.Compile("<piece-loc>.+?piece-loc>")
		pieceLocs := re.FindAllStringSubmatch(mainString, -1)
		log.Println(pieceLocs)
		for _, pl := range pieceLocs {
			pieceLoc := pl[0]
			re, _ = regexp.Compile("<pawn>((.|\\s) |[^(pawn)])*pawn>")
			pawnString := (re.FindAllStringSubmatch(pieceLoc, -1))[0][0] //assuming each piece loc will only have one pawn
			pawn := interpretPawns(pawnString)[0]
			log.Println(pawn)
			re, _ = regexp.Compile("<loc>.+?loc>")
			locTag := re.FindStringSubmatch(pieceLoc) //assuming each piece loc will only have one loc
			re, _ = regexp.Compile("\\d+")
			loc := re.FindStringSubmatch(locTag[0]) //assuming each piece loc will only have one loc
			rawIndex := (loc[0][0])
			log.Println(rawIndex)
		}
		return
		log.Println("!!!!!!!Home")
		//building the home
		re, _ = regexp.Compile("<home>(.|\\s)*home>")
		homeStrings := re.FindStringSubmatch(boardString)
		homeString := homeStrings[0]
		re, _ = regexp.Compile("<pawn>((.|\\s) |[^(pawn)])*pawn>")
		homePawns := re.FindAllStringSubmatch(homeString, -1)
		log.Println(homePawns)
		for _, pawn := range homePawns {
			log.Println("+++++++")
			log.Println(pawn[0])
		}

	}
	if strings.Contains(s, "doubles-penalty") {
		log.Println("this tells me I got a doubles penalty")
		player.DoublesPenalty()
		response := fmt.Sprintf("<void></void>")
		w.Write([]byte(response))
		log.Println("I have replied with a doubles penalty")
	}

	// w.WriteHeader(http.StatusOK)
}

func main() {
	const playerName = "Jimmy"
	player.name = playerName
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", interpret).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

/////////////ACTUAL PLAYER IMPLEMENTATION

type MyPlayer struct {
	color string
	name  string
	id    int
}

func (sp *MyPlayer) startGame(color string) string {
	sp.color = color
	return sp.name
}

// func (sp SPlayer) doMove(board Board, rolls []int) []Move {
// 	return nil
// }

func (sp *MyPlayer) DoublesPenalty() {
	log.Println("!!!!!!!! not implemented : doubles penalty")
}
