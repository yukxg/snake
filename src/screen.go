package src

import (
	"fmt"
	"math/rand"
)

type Screen struct {
	snake *Snake
	height int
	width int
	food *Food
}
func NewScreen(snake *Snake) *Screen {
	screen := &Screen{
		height:10,
		width:10,
		snake:snake,
	}
	screen.placeFood()
	return screen
}




func PrintScreen(s *Screen) error{
	for i:= 0; i < s.width; i++ {
		for j := 0; j < s.height; j++ {
			if s.food.x == i && s.food.y == j {
				fmt.Print("$")
				continue
			}
			if s.snake.isOnposition(Coord{X:i, Y:j}) {
				fmt.Print("*")
				continue
			}else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	return nil
}

func (screen *Screen) placeFood () {
	var x, y int
	for {
		x = rand.Intn(screen.width)
		y = rand.Intn(screen.height)
		if(!screen.snake.isOnposition(Coord{X:x, Y:y})) {
			break;
		}
	}
	screen.food = NewFood(x, y)
	fmt.Printf("food coor {%d, %d}\n", x, y)
}
