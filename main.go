package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"noita-go/noita"
)

func main() {
	scene := noita.NewScene()
	//w, h := scene.GetDimensions()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle(scene.Title)
	if err := ebiten.RunGame(scene); err != nil {
		log.Fatal(err)
	}
}
