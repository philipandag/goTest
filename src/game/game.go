package game

import "strconv"

type Object struct {
	X float64
	Y float64
}

func (o *Object) ToString() string {
	return "{" + strconv.FormatFloat(o.X, 'f', 2, 64) + ", " + strconv.FormatFloat(o.Y, 'f', 2, 64) + "}"
}

type Game struct {
	GameObject Object
}
