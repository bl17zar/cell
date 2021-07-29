package cell

import (
	"errors"
	"fmt"
)

type Hasher interface {
	Hash() string
}

type history struct {
	values map[string]Hasher
	orders []string
}

func newHistory() *history {
	return &history{
		values: make(map[string]Hasher),
	}
}

func (h *history) Put(v Hasher) {
	hash := v.Hash()

	h.orders = append(h.orders, hash)
	h.values[hash] = v
}

func (h *history) GetByIdx(i int) (Hasher, error) {
	if i >= len(h.orders) {
		return nil, errors.New(fmt.Sprint("no such index in history: ", i))
	}

	return h.values[h.orders[i]], nil
}

func (h *history) Exists(v Hasher) (exists bool) {
	_, exists = h.values[v.Hash()]
	return
}

func (h *history) Distance(v Hasher) int {
	hash := v.Hash()
	for i, oh := range h.orders {
		if oh == hash {
			return i
		}
	}

	return -1
}
