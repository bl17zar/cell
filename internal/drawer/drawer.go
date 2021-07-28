package drawer

import (
	mapInternal "github.com/bl17zar/cell/internal/map"
)

type Drawer interface {
	Draw(m *mapInternal.Map)
}
