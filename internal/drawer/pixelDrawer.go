package drawer

import (
	mapInternal "github.com/bl17zar/cell/internal/map"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"image/color"
)

var pixelColors = map[string][]uint8{
	mapInternal.SymbolNode.String():                colorToPixel(colornames.Cyan),
	mapInternal.SymbolEdgeHorizontal.String():      colorToPixel(colornames.Cyan),
	mapInternal.SymbolEdgeVertical.String():        colorToPixel(colornames.Cyan),
	mapInternal.SymbolObstacle.String():            colorToPixel(colornames.Red),
	mapInternal.SymbolCycleEdgeVertical.String():   colorToPixel(colornames.Gray),
	mapInternal.SymbolCycleEdgeHorizontal.String(): colorToPixel(colornames.Gray),
	mapInternal.SymbolCycleNode.String():           colorToPixel(colornames.Gray),
}

type PixelDrawer struct {
	win *pixelgl.Window
}

func NewPixelDrawer() *PixelDrawer {
	cfg := pixelgl.WindowConfig{
		Title:  "Cell Evolution",
		Bounds: pixel.R(0, 0, 512, 512),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	go func() {
		for !win.Closed() {
			win.Update()
		}
		win.Destroy()
	}()

	return &PixelDrawer{
		win: win,
	}
}

func (p *PixelDrawer) Draw(m *mapInternal.Map) {
	p.win.Canvas().SetPixels(p.drawPixels(m))
}

func (p *PixelDrawer) drawPixels(m *mapInternal.Map) []uint8 {
	size := int(p.win.Bounds().Max.X)

	rowSize := 4 * size

	pixels := make([]uint8, rowSize*size, rowSize*size)

	cellSize := size / len(m.Values)
	cellSize = cellSize - cellSize/4

	for x, row := range m.Values {
		for y, el := range row {
			drawElem(x*(1+cellSize), y*(1+cellSize), el, cellSize, &pixels, rowSize)
		}
	}

	return pixels
}

func writePixel(x, y int, pix []uint8, dest *[]uint8, rowSize int) {
	for i, v := range pix {
		(*dest)[x*rowSize+y*4+i] = v
	}
}

func drawElem(x, y int, elem mapInternal.Symbol, size int, dest *[]uint8, rowSize int) {
	switch elem {
	case mapInternal.SymbolNode, mapInternal.SymbolObstacle, mapInternal.SymbolCycleNode:
		drawSquare(x, y, pixelColors[elem.String()], size, dest, rowSize)
	case mapInternal.SymbolCycleEdgeHorizontal, mapInternal.SymbolEdgeHorizontal:
		drawHorizontalLine(x, y, pixelColors[elem.String()], size, dest, rowSize)
	case mapInternal.SymbolCycleEdgeVertical, mapInternal.SymbolEdgeVertical:
		drawVerticalLine(x, y, pixelColors[elem.String()], size, dest, rowSize)
	}
}

func drawVerticalLine(x, y int, pix []uint8, size int, dest *[]uint8, rowSize int) {
	for i := 0; i < size; i++ {
		writePixel(x+i, y+size/2, pix, dest, rowSize)
	}
}

func drawHorizontalLine(x, y int, pix []uint8, size int, dest *[]uint8, rowSize int) {
	for i := 0; i < size; i++ {
		writePixel(x+size/2, y+i, pix, dest, rowSize)
	}
}

func drawSquare(x, y int, pix []uint8, size int, dest *[]uint8, rowSize int) {
	for i := 0; i < size; i++ {
		writePixel(x+i, y+size, pix, dest, rowSize)
		writePixel(x+size, y+i, pix, dest, rowSize)
		writePixel(x, y+i, pix, dest, rowSize)
		writePixel(x+i, y, pix, dest, rowSize)
	}
}

func colorToPixel(c color.Color) []uint8 {
	res := make([]uint8, 4, 4)
	v1, v2, v3, v4 := c.RGBA()
	res[0], res[1], res[2], res[3] = uint8(v1), uint8(v2), uint8(v3), uint8(v4)
	return res
}
