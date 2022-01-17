package noita

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mazznoer/colorgrad"
	"image"
	"image/color"
	"noita-go/model"
	"noita-go/model/solid"
	"noita-go/utils"
)

const (
	scale      = 4
	width      = 1024
	height     = 768
	bufferSize = 64
	wScaled    = width / scale
	hScaled    = height / scale
)

type CellType uint8

type Scene struct {
	Title        string
	Img          *ebiten.Image
	Grid         *model.Grid
	FireGradient colorgrad.Gradient
	IsPainting   bool
	ImgBuffer    [][]*image.Image
	PaintType    CellType
	PreviousImg  *image.Image
	Paused       bool
}

func NewScene() *Scene {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.RGBA{})
	gradient := colorgrad.Inferno()
	grid := model.NewGrid()
	for y := 0; y < hScaled; y++ {
		row := make([]*model.Cell, 0)
		for x := 0; x < wScaled; x++ {
			xf64 := float64(x)
			yf64 := float64(y)
			row = append(row, model.NewCell(xf64, yf64, scale, img, grid))
		}
		grid.Cells = append(grid.Cells, row)
	}
	grid.FillCellNeighbor()
	return &Scene{
		Title:        "Noita Go",
		Img:          img,
		FireGradient: gradient,
		Grid:         grid,
		PaintType:    1,
	}
}

func (s *Scene) GetDimensions() (int, int) {
	return width, height
}

func (s *Scene) PaintElement(cType CellType) model.Element {
	switch cType {
	case 0:
		return model.NewEmpty()
	case 1:
		return solid.NewSand()
	case 2:
		return solid.NewSand()
	case 3:
		return solid.NewGround()
	default:
		return solid.NewSand()
	}

}

func (s *Scene) Painting(cType CellType) {
	mx, my := ebiten.CursorPosition()
	rx := mx / scale
	ry := my / scale
	if rx > 0 && rx < wScaled && ry > 0 && ry < hScaled {
		for j := -1; j <= 1; j++ {
			for i := -1; i <= 1; i++ {
				cell := s.Grid.Get(rx+i, ry+j)
				if cell == nil {
					continue
				}
				cell.Tick = s.Grid.Tick
				cell.Element = s.PaintElement(cType)
			}
		}
	}
}

func (s *Scene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		for _, row := range s.Grid.Cells {
			for _, cell := range row {
				cell.Element = model.NewEmpty()
			}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		s.Paused = !s.Paused
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
		// drop a grain of sand every 10 ticks
		if s.Grid.Tick%10 == 0 {
			s.Grid.Cells[0][wScaled/2].Element = solid.NewSand()
		}
	}
	if !s.Paused {
		s.Grid.Update()
	}
	return nil
}

func (s *Scene) BrushLabel() string {
	switch s.PaintType {
	case 0:
		return "Eraser"
	case 1:
		return "Sand"
	case 2:
		return "Water"
	case 3:
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
