package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"noita-go/noita"
)

func main() {
	scene := noita.NewScene()
	//w, h := scene.GetDimensions()
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle(scene.Title)
	if err := ebiten.RunGame(scene); err != nil {
		log.Fatal(err)
	}
}
