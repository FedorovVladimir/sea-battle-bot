package main

import (
	"fmt"
	"sea-battle-bot/field"
)

func main() {
	f, err := field.NewField(field.Min)
	if err != nil {
		panic(err)
	}
	f.SetShip(0, 0)
	fmt.Println(f.InBattle())
	fmt.Println(f)

	fmt.Println()

	f.Shot(0, 0)
	fmt.Println(f.InBattle())
	fmt.Println(f)
}
