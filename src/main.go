package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ntakahashi016/parallel_shooter/parallel_shooter"
)
func main() {
	game, err := parallel_shooter.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Parallel Shooter")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
