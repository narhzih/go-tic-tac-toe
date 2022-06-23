package ui

import (
	"fmt"
	"github.com/narhzih/go-tic-tac-toe/pkg/utils"
	"log"
)

var BoardValue *[3][3]string = &[3][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}
var boardValueMap map[int]string = map[int]string{
	1: "0:0",
	2: "0:1",
	3: "0:2",
	4: "1:0",
	5: "1:1",
	6: "1:2",
	7: "2:0",
	8: "2:1",
	9: "2:2",
}

func PrintBoard() {
	fmt.Printf("----------\n")
	fmt.Printf(" %+v | %+v | %+v \n", BoardValue[0][0], BoardValue[0][1], BoardValue[0][2])
	fmt.Printf("----------\n")
	fmt.Printf(" %+v | %+v | %+v \n", BoardValue[1][0], BoardValue[1][1], BoardValue[1][2])
	fmt.Printf("----------\n")
	fmt.Printf(" %+v | %+v | %+v \n", BoardValue[2][0], BoardValue[2][1], BoardValue[2][2])
	fmt.Printf("----------")
}

func UpdateBoard(boardNumber int, value string) {
	boardNumberPair := boardValueMap[boardNumber]
	indexOne, indexTwo, parseErr := utils.SplitStringToInt(boardNumberPair)
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	BoardValue[indexOne][indexTwo] = value
}
