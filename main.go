package main

import (
	"fmt"
	"sea-battle-bot/field"
)

func main() {
	f := field.NewField(4)
	f.SetShip(0, 0)
	f.SetShip(0, 1)
	f.SetShip(0, 2)
	f.SetShip(10, 3)
	f.Shot(0, 10)
	f.Shot(-3, 4)
	f.Shot(0, -4)
	fmt.Println(f)
}
