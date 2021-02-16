package snake

import "errors"

const (
	RIGHT = 1 + iota
	LEFT
	UP
	DOWN
)

type (
	direction int
	snake struct {
		body []coord
		direction direction
		length int
	}
)

func (s *snake) changeDirection(d direction) {
	opposites := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}
	if o := opposites[d]; o != 0 && 0 != s.direction {
		s.direction = d
	}
}

func (s *snake) head() coord {
	return s.body[len(s.body) - 1]
}

func (s *snake) isOnPosition(c coord) bool {
	for _, b := range s.body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}

func (s *snake) die() error {
	return errors.New("Died")
}

func (s *snake) move() error {
	h := s.head()
	c := coord{x: h.x, y: h.y}
	switch s.direction {
	case RIGHT:
		c.x++
	case LEFT:
		c.x--
	case UP:
		c.y++
	case DOWN:
		c.y--
	}
	if s.isOnPosition(c) {
		return s.die()
	}
	if s.length > len(s.body) {
		s.body = append(s.body, c)
	} else {
		s.body = append(s.body[1:], c)
	}
	return nil
}