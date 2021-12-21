package main

import (
	"fmt"
	"sea-battle-bot/field"
)

func main() {
	f := field.NewField(4)
	f.SetShip(0, 0)
	fmt.Println(f.InBattle())
	fmt.Println(f)

	fmt.Println()

	f.Shot(0, 0)
	fmt.Println(f.InBattle())
	fmt.Println(f)
}
