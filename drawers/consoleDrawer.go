package drawers

import (
	"fmt"
	"strings"

	"github.com/bl17zar/cell/cell"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[97m"
	colorWhite  = "\033[97m"
)

type ConsoleDrawer struct{}

func (d *ConsoleDrawer) Draw(m *cell.Map) {
	for _, r := range m.Rows() {
		imageRow := []string{}
		b := strings.Builder{}

		for _, sym := range r {
			if string(sym) == string(cell.SignObstacle) {
				imageRow = append(imageRow, b.String())

				imageRow = append(imageRow, fmt.Sprintf("%s%s%s", colorRed, string(sym), colorReset))

				b = strings.Builder{}

				continue
			}

			if string(sym) == string(cell.SignNode) || string(sym) == string(cell.SignEdgeHorizontal) || string(sym) == string(cell.SignEdgeVertical) {
				imageRow = append(imageRow, b.String())

				imageRow = append(imageRow, fmt.Sprintf("%s%s%s", colorCyan, string(sym), colorReset))

				b = strings.Builder{}

				continue
			}

			b.WriteString(string(sym))
		}

		imageRow = append(imageRow, b.String())

		for _, el := range imageRow {
			fmt.Print(el)
		}

		fmt.Print("\n")
	}
	fmt.Print("\n")
}
