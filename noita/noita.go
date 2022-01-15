package noita

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mazznoer/colorgrad"
	"image"
	"image/color"
	"math/rand"
	"noita-go/utils"
)

const (
	scale      = 2
	width      = 640
	height     = 480
	bufferSize = 64
	wScaled    = width / scale
	hScaled    = height / scale
)

type Scene struct {
	Title        string
	Img          *ebiten.Image
	Grid         *Grid
	FireGradient colorgrad.Gradient
	IsPainting   bool
	ImgBuffer    [][]*image.Image
	PaintType    CellType
}

func NewScene() *Scene {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)
	gradient := colorgrad.Inferno()
	grid := NewGrid()
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
		Grid:         grid,
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
		for j := -2; j <= 2; j++ {
			for i := -2; i <= 2; i++ {
				cell := s.Grid.Get(rx+i, ry+j)
				if cell == nil {
					continue
				}
				cell.Tick = s.Grid.Tick
				cell.Type = cType
				cell.Alpha = rand.Float64() + 0.3
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
	screen.Fill(color.Transparent)
	s.Grid.Draw(screen)
	imgBuffer := make([][]*image.Image, 0)
	for j := 0; j < height/bufferSize; j++ {
		row := make([]*image.Image, 0)
		for i := 0; i < width/bufferSize; i++ {
			x0 := i * bufferSize
			y0 := i * bufferSize
			bounds := image.Rect(x0, y0, x0+bufferSize, y0+bufferSize)
			img := screen.SubImage(bounds)
			row = append(row, &img)
		}
		imgBuffer = append(imgBuffer, row)
	}
	s.ImgBuffer = imgBuffer
	utils.DebugInfo(screen)
	utils.DebugInfoMessage(screen, fmt.Sprintf("\n\nPress [A] to change brush: %s", s.BrushLabel()))
}

func (s *Scene) Layout(oW, oH int) (int, int) {
	return s.GetDimensions()
}
