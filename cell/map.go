package cell

import (
	"fmt"
	"strings"
)

type SignType string

const (
	SignNode           SignType = "◼"
	SignEdgeHorizontal SignType = "—"
	SignEdgeVertical   SignType = "|"

	// SignEmpty    SignType = "·"
	SignEmpty    SignType = " "
	SignObstacle SignType = "⊠"

	SignCycleNode                    = "□"
	SignCycleEdgeHorizontal SignType = "⋯"
	SignCycleEdgeVertical   SignType = "⋮"
)

type Map struct {
	xExtender string
	size      int
	Values    [][]SignType
}

func NewMap(size, xMult int, features []*Display) *Map {
	values := make([][]SignType, 0, size)
	for i := 0; i < size; i++ {
		values = append(values, make([]SignType, size, size))
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			values[i][j] = SignEmpty
		}
	}

	if len(features) != 0 {
		for _, f := range features {
			values[f.Row][f.Col] = f.Sign
		}
	}

	xMultList := make([]string, 0, xMult)
	for i := 0; i < xMult; i++ {
		xMultList = append(xMultList, " ")
	}

	return &Map{
		xExtender: strings.Join(xMultList, ""),
		size:      size,
		Values:    values,
	}
}

func (m *Map) Copy() *Map {
	res := NewMap(m.size, len(strings.Split(m.xExtender, "")), nil)

	for i, row := range m.Values {
		res.Values[i] = row
	}

	return res
}

func (m *Map) Set(row int, column int, v SignType) {
	m.Values[row-1][column-1] = v
}

func (m *Map) String() string {
	b := strings.Builder{}

	for _, row := range m.Values {
		for _, el := range row {
			b.WriteString(fmt.Sprint(m.xExtender, el))
		}

		b.WriteString("\n")
	}

	return b.String()
}

func (m *Map) Rows() []string {
	res := make([]string, 0, len(m.Values))

	for _, row := range m.Values {
		b := strings.Builder{}
		for _, el := range row {
			b.WriteString(fmt.Sprint(m.xExtender, el))
		}

		res = append(res, b.String())
	}

	return res
}

func (m *Map) IsInsideBorders(row, col int) bool {
	return row-1 >= 0 && row-1 < m.size && col-1 >= 0 && col-1 < m.size
}

func (m *Map) IsEmpty(row int, col int) bool {
	return m.IsInsideBorders(row, col) && m.Values[row-1][col-1] == SignEmpty
}
