package noita

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mazznoer/colorgrad"
	"image"
	"math"
	"noita-go/model"
	"noita-go/model/liquid"
	"noita-go/model/solid"
	"noita-go/utils"
)

const (
	scale   = 2
	width   = 640
	height  = 480
	wScaled = width / scale
	hScaled = height / scale
)

type CellType uint8

type Scene struct {
	Title        string
	Grid         *model.Grid
	FireGradient colorgrad.Gradient
	IsPainting   bool
	ImgBuffer    [][]*image.Image
	PaintType    CellType
	PreviousImg  *image.Image
	Paused       bool
	Canvas       *image.RGBA
	prevX        int
	prevY        int
}

func NewScene() *Scene {
	gradient := colorgrad.Inferno()
	grid := model.NewGrid()
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < hScaled; y++ {
		row := make([]*model.Cell, 0)
		for x := 0; x < wScaled; x++ {
			xf64 := float64(x)
			yf64 := float64(y)
			row = append(row, model.NewCell(xf64, yf64, scale, grid))
		}
		grid.Cells = append(grid.Cells, row)
	}
	grid.FillCellNeighbor()
	return &Scene{
		Title:        "Noita Go",
		FireGradient: gradient,
		Grid:         grid,
		PaintType:    1,
		Canvas:       canvas,
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
		return liquid.NewWater()
	case 3:
		return solid.NewGround()
	default:
		return solid.NewSand()
	}

}

const paintSize = 1

func (s *Scene) PaintAt(x, y int, cType CellType) {
	rx := x / scale
	ry := y / scale
	if rx < 0 || rx > wScaled || ry < 0 || ry > hScaled {
		return
	}
	for j := -paintSize; j <= paintSize; j++ {
		for i := -paintSize; i <= paintSize; i++ {
			cell := s.Grid.Get(rx+i, ry+j)
			if cell == nil {
				continue
			}
			cell.Tick = s.Grid.Tick
			cell.Element = s.PaintElement(cType)
		}
	}
}

func (s *Scene) PaintSloped(mx, my int, cType CellType) {
	if mx < 0 || mx > width || my < 0 || my > height {
		return
	}
	dx := float64(mx - s.prevX)
	dy := float64(my - s.prevY)
	absDx := math.Abs(dx)
	absDy := math.Abs(dy)
	if dx == 0 || dy == 0 {
		return
	}
	xDiffIsLarger := absDx > absDy
	stepX := 1.0
	if dx < 0 {
		stepX = -1.0
	}
	stepY := 1.0
	if dy < 0 {
		stepY = -1.0
	}
	longerSideLength := math.Max(absDx, absDy)
	shorterSideLength := math.Min(absDx, absDy)
	slope := shorterSideLength / longerSideLength

	for i := 1.0; i <= longerSideLength; i++ {
		shorterSideIncrease := math.Round(i * slope)
		xIncrease := 0.0
		yIncrease := 0.0
		if xDiffIsLarger {
			xIncrease = i
			yIncrease = shorterSideIncrease
		} else {
			xIncrease = shorterSideIncrease
			yIncrease = i
		}
		toX := s.prevX + int(math.Round(xIncrease*stepX))
		toY := s.prevY + int(math.Round(yIncrease*stepY))
		s.PaintAt(toX, toY, cType)
	}

	s.prevX = mx
	s.prevY = my

}

func (s *Scene) Painting(cType CellType) {
	mx, my := ebiten.CursorPosition()
	s.PaintAt(mx, my, cType)
	s.PaintSloped(mx, my, cType)
}

func (s *Scene) Update() error {
	s.Canvas = image.NewRGBA(image.Rect(0, 0, width, height))
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
		s.prevX, s.prevY = ebiten.CursorPosition()
		s.IsPainting = true
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		s.IsPainting = false
	}
	if s.IsPainting {
		s.Painting(s.PaintType)
	} else {
		// drop a grain of sand every 10 ticks
		//if s.Grid.Tick%10 == 0 {
		//	s.Grid.Cells[0][wScaled/2].Element = solid.NewSand()
		//}
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
	s.Grid.Draw(screen, s.Canvas)
	//s.Canvas = image.NewRGBA(image.Rect(0, 0, width, height))
	utils.DebugInfo(screen)
	utils.DebugInfoMessage(screen, fmt.Sprintf("\n\nPress [A] to change brush: %s", s.BrushLabel()))
}

func (s *Scene) Layout(oW, oH int) (int, int) {
	return s.GetDimensions()
}
