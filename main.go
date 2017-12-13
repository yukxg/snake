package main

import (
	"bufio"
	"os"
	"snake/src"
)

func main()  {
	s := src.NewSnake("D", []src.Coord{src.Coord{X:1,Y:1}},)
	screen := src.NewScreen(s)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for scanner.Scan() {
		switch scanner.Text() {
		default:
			if err:= s.Move(scanner.Text(),screen); err !=nil {
				panic(err)
				break
			}
		case "EXIT":
			return
		}
	}
}


