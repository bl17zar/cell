package _map

type Symbol int

const (
	SymbolEmpty Symbol = iota
	SymbolNode
	SymbolEdgeHorizontal
	SymbolEdgeVertical
	SymbolObstacle
	SymbolCycleNode
	SymbolCycleEdgeHorizontal
	SymbolCycleEdgeVertical
)

var symbolMap = map[int]string{
	0: " ",
	1: "◼",
	2: "—",
	3: "|",
	4: "⊠",
	5: "□",
	6: "⋯",
	7: "⋮",
}

func (s Symbol) String() string {
	return symbolMap[int(s)]
}
