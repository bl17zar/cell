package drawer

import (
	"fmt"
	"strings"

	mapInternal "github.com/bl17zar/cell/internal/map"
)

type color string

const (
	reset  color = "\033[0m"
	red    color = "\033[31m"
	green  color = "\033[32m"
	yellow color = "\033[33m"
	blue   color = "\033[34m"
	purple color = "\033[35m"
	cyan   color = "\033[36m"
	gray   color = "\033[97m"
	white  color = "\033[97m"
)

const widthWidener = " "

var colors = map[string]color{
	mapInternal.SymbolNode.String():           cyan,
	mapInternal.SymbolEdgeHorizontal.String(): cyan,
	mapInternal.SymbolEdgeVertical.String():   cyan,
	mapInternal.SymbolObstacle.String():       red,
}

type ConsoleDrawer struct {
	WidthWidenerSize int
	widthWidener     string
}

func NewConsoleDrawer(wwSize int) *ConsoleDrawer {
	wwList := make([]string, 0, wwSize)
	for i := 0; i < wwSize; i++ {
		wwList = append(wwList, widthWidener)
	}

	return &ConsoleDrawer{
		WidthWidenerSize: wwSize,
		widthWidener:     strings.Join(wwList, ""),
	}
}

func (d *ConsoleDrawer) Draw(m *mapInternal.Map) {
	for _, r := range m.Values {
		b := strings.Builder{}
		imageRow := []string{}
		for _, sym := range r {
			b = d.handleSign(sym.String(), colors[sym.String()], &imageRow, b)

			continue
		}

		imageRow = append(imageRow, b.String())
		for _, iR := range imageRow {
			fmt.Print(iR)
		}

		fmt.Print("\n")
	}

	fmt.Print("\n")
}

func (d *ConsoleDrawer) handleSign(sym string, c color, destRow *[]string, currBuilder strings.Builder) strings.Builder {
	*destRow = append(*destRow, currBuilder.String())
	*destRow = append(*destRow, fmt.Sprintf("%s%s%s", c, fmt.Sprint(d.widthWidener, sym), reset))

	return strings.Builder{}
}
