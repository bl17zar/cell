package drawers

import (
	"github.com/bl17zar/cell/cell"
)

type Drawer interface {
	Draw(*cell.Map)
}
