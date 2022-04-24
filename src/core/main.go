package main

import (
	"fmt"
	"originalGod/src/game"
)

func main() {
	player := game.NewTestPlayer()
	fmt.Println(player.ModIdcard)
	player.RecvSetIcon(1)   // 胡桃icon
	player.RecvSetIdcard(2) // 枫叶
}
