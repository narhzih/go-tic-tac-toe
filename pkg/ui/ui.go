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

	fmt.Printf(" %+v | %+v | %+v \n", BoardValue[0][0], BoardValue[0][1], BoardValue[0][2])
	fmt.Printf("----------\n")
	fmt.Printf(" %+v | %+v | %+v \n", BoardValue[1][0], BoardValue[1][1], BoardValue[1][2])
	fmt.Printf("----------\n")
	fmt.Printf(" %+v | %+v | %+v \n", BoardValue[2][0], BoardValue[2][1], BoardValue[2][2])

}

func UpdateBoard(boardNumber int, value string) {
	boardNumberPair := boardValueMap[boardNumber]
	indexOne, indexTwo, parseErr := utils.SplitStringToInt(boardNumberPair)
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	BoardValue[indexOne][indexTwo] = value
}

func PlayerHasWon() bool {
	//  1. Check for rows
	for _, i := range BoardValue {
		var rowValues []string
		for _, j := range i {
			rowValues = append(rowValues, j)
		}
		if rowValues[0] == rowValues[1] && rowValues[1] == rowValues[2] {
			return true
		}
	}

	// 2. Check for columns
	var colValues []string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			colValues = append(colValues, BoardValue[j][i])
		}
		if colValues[0] == colValues[1] && colValues[1] == colValues[2] {
			return true
		}
		colValues = []string{}
	}

	// 3. Check for diagonals
	if BoardValue[0][0] == BoardValue[1][1] && BoardValue[1][1] == BoardValue[2][2] {
		return true
	} else if BoardValue[0][2] == BoardValue[1][1] && BoardValue[1][1] == BoardValue[2][0] {
		return true
	}
	return false
}

func GameTie() bool {
	return false
}
