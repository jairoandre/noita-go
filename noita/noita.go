package noita

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mazznoer/colorgrad"
	"image/color"
	"noita-go/utils"
)

const (
	scale   = 2
	width   = 800
	height  = 600
	wScaled = width / scale
	hScaled = height / scale
)

type Scene struct {
	Title        string
	Img          *ebiten.Image
	Grid         *Grid
	FireGradient colorgrad.Gradient
	GifWriter    *utils.GifWriter
	IsPainting   bool
	PaintType    CellType
}

func NewScene() *Scene {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)
	gradient := colorgrad.Inferno()
	grid := Grid{
		Cells: make([][]*Cell, 0),
	}
	for y := 0; y < hScaled; y++ {
		row := make([]*Cell, 0)
		for x := 0; x < wScaled; x++ {
			xf64 := float64(x)
			yf64 := float64(y)
			row = append(row, NewCell(xf64, yf64, img, empty))
		}
		grid.Cells = append(grid.Cells, row)
	}
	return &Scene{
		Title:        "Noita Go",
		Img:          img,
		FireGradient: gradient,
		Grid:         &grid,
		PaintType:    sand,
	}
}

func (s *Scene) GetDimensions() (int, int) {
	return width, height
}

func (s *Scene) Painting(cType CellType) {
	mx, my := ebiten.CursorPosition()
	rx := mx / scale
	ry := my / scale
	if rx > 0 && rx < wScaled && ry > 0 && ry < hScaled {
		for j := -3; j <= 3; j++ {
			for i := -3; i <= 3; i++ {
				cell := s.Grid.Get(rx+i, ry+j)
				if cell == nil {
					continue
				}
				cell.Type = cType
				cell.Alpha = 1.0
			}
		}
	}
}

func (s *Scene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		for _, row := range s.Grid.Cells {
			for _, cell := range row {
				cell.Type = empty
			}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		s.PaintType = (s.PaintType + 1) % 4
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.IsPainting = true
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		s.IsPainting = false
	}
	if s.IsPainting {
		s.Painting(s.PaintType)
	} else {
		if s.Grid.Tick%10 == 0 {
			s.Grid.Cells[0][wScaled/2].Type = sand
		}
	}
	s.Grid.Update()
	return nil
}

func (s *Scene) BrushLabel() string {
	switch s.PaintType {
	case empty:
		return "Eraser"
	case sand:
		return "Sand"
	case water:
		return "Water"
	case ground:
		return "Ground"
	default:
		return "-"
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.Grid.Draw(screen)
	utils.DebugInfo(screen)
	utils.DebugInfoMessage(screen, fmt.Sprintf("\n\nPress [A] to change brush: %s", s.BrushLabel()))
}

func (s *Scene) Layout(oW, oH int) (int, int) {
	return s.GetDimensions()
}
