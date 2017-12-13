package src

import (
	"fmt"
	"errors"
)


const (
	RIGHT = "D"
	LEFT = "A"
	UP = "W"
	DOWN = "S"
	HEIGHT = 10
	WIDTH = 10
)
type Snake struct {
	body []Coord
	len int
	direction string
}

func (snake *Snake) isOnposition(c Coord) bool{
	for _, b := range snake.body {
		if b.X == c.X && b.Y == c.Y {
			return true
		}
	}
	return false
}

func NewSnake(d string, b []Coord) *Snake {
	return &Snake{
		len: len(b),
		body: b,
		direction: d,
	}
}

func (s *Snake) head() Coord {
	return s.body[len(s.body)-1]
}

func (s *Snake) eatFood (c Coord) bool {
	return s.head().X == c.X && s.head().Y == c.Y
}



func (s *Snake) Move(dir string, screen *Screen) error{
	h := s.head()
	if h.X > WIDTH || h.Y > HEIGHT || h.X < 0 || h.Y < 0 {
		return errors.New("GAME OVER")
	}
	c := Coord{X: h.X, Y: h.Y}
	switch dir {
	case RIGHT:
		c.Y += 1
	case LEFT:
		c.Y -= 1
	case UP:
		c.X -= 1
	case DOWN:
		c.X += 1
	}
	if h.X == screen.food.x && h.Y == screen.food.y{
		s.len++
		fmt.Printf("eat! + len: %d\n",s.len)
		s.body = append(s.body, c)
		screen.placeFood()
		screen.snake = s
		PrintScreen(screen)
	}
	if s.len > len(s.body) {

	} else {
		s.body = append(s.body[1:], c)
	}
	fmt.Printf("hahahahhah- x : %d y : %d\n", h.X, h.Y)
	PrintScreen(screen)
	return nil
}
