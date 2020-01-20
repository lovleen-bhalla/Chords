package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	chords := `{
		"am": {
			"5": 0,
			"4": 2,
			"3": 2,
			"2": 1
		}
	}`
	var f interface{}
	json.Unmarshal([]byte(chords), &f)
	chordMap := f.(map[string]interface{})

	for _, arg := range os.Args[1:] {
		if c, ok := chordMap[arg]; ok {
			chord := c.(map[string]interface{})
			fmt.Println(getChordTab(chord))
		} else {
			panic("chord not present in system")
		}
	}
}

func getChordTab(chord map[string]interface{}) []string {
	frets := 3
	var board []string
	for stringNO := 0; stringNO < 6; stringNO++ {
		var row string
		for fret := 0; fret < frets; fret++ {
			cell := "I-----"
			column := chord[strconv.Itoa(stringNO+1)]
			switch vv := column.(type) {
			case float64:
				if int(vv)-1 == fret {
					cell = "I--" + strconv.Itoa(int(vv)) + "--"
				}
			}
			row = row + cell

		}
		board = append(board, row)
	}
	return board
}
