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

const (
	spacer     = " "
	line_up    = "\033[1A"
	line_clear = "\x1b[2K"
)

var colors = map[string]consoleColor{
	mapInternal.SymbolNode.String():           cyan,
	mapInternal.SymbolEdgeHorizontal.String(): cyan,
	mapInternal.SymbolEdgeVertical.String():   cyan,
	mapInternal.SymbolObstacle.String():       red,
}

type ConsoleDrawer struct {
	opts      *Opts
	spaceSize int
	spacer    string
}

type Opts struct {
	UseClear  bool
	ClearSize int
}

func NewConsoleDrawer(spacerSize int, opts *Opts) *ConsoleDrawer {
	wwList := make([]string, 0, spacerSize)
	for i := 0; i < spacerSize; i++ {
		wwList = append(wwList, spacer)
	}

	return &ConsoleDrawer{
		opts:      opts,
		spaceSize: spacerSize,
		spacer:    strings.Join(wwList, ""),
	}
}

func (d *ConsoleDrawer) Draw(m *mapInternal.Map) {
	if d.opts != nil && d.opts.UseClear {
		d.clear(d.opts.ClearSize)
	}

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
}

func (d *ConsoleDrawer) handleSign(sym string, c consoleColor, destRow *[]string, currBuilder strings.Builder) strings.Builder {
	*destRow = append(*destRow, currBuilder.String())
	*destRow = append(*destRow, fmt.Sprintf("%s%s%s", c, fmt.Sprint(d.spacer, sym), reset))

	return strings.Builder{}
}

func (d *ConsoleDrawer) clear(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print(line_clear)
		fmt.Print(line_up)
	}
}
