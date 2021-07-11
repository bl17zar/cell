package cell

type orderedSet struct {
	values map[string]int
	orders []string
}

func newOrderedSet(vals ...string) orderedSet {
	os := orderedSet{
		values: make(map[string]int, len(vals)),
		orders: make([]string, 0, len(vals)),
	}

	if len(vals) != 0 {
		for _, v := range vals {
			if _, ok := os.values[v]; !ok {
				os.values[v] = len(os.orders)
				os.orders = append(os.orders, v)
			}
		}
	}

	return os
}

func (o orderedSet) Put(v string) orderedSet {
	newOrders := make([]string, 0, len(o.orders) + 1)
	newOrders = append(newOrders, o.orders...)
	newOrders = append(newOrders, v)

	return newOrderedSet(newOrders...)
}

func (o orderedSet) HasNotLast(v string) bool {
	if i, ok := o.values[v]; ok && i != len(o.orders)-1 {
		return true
	}

	return false
}

func (o orderedSet) GetLast() string {
	if len(o.orders) == 0 {
		return ""
	}

	return o.orders[len(o.orders)-1]
}

func (o orderedSet) Unwind(v string) []string {
	return o.orders[o.values[v]:]
}
