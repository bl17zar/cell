package drawer

import (
	"fmt"
	"strings"

	mapInternal "github.com/bl17zar/cell/internal/map"
)

type consoleColor string

const (
	reset  consoleColor = "\033[0m"
	red    consoleColor = "\033[31m"
	green  consoleColor = "\033[32m"
	yellow consoleColor = "\033[33m"
	blue   consoleColor = "\033[34m"
	purple consoleColor = "\033[35m"
	cyan   consoleColor = "\033[36m"
	gray   consoleColor = "\033[97m"
	white  consoleColor = "\033[97m"
)

const widthWidener = " "

var colors = map[string]consoleColor{
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

func (d *ConsoleDrawer) handleSign(sym string, c consoleColor, destRow *[]string, currBuilder strings.Builder) strings.Builder {
	*destRow = append(*destRow, currBuilder.String())
	*destRow = append(*destRow, fmt.Sprintf("%s%s%s", c, fmt.Sprint(d.widthWidener, sym), reset))

	return strings.Builder{}
}
