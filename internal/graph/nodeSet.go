package graph

type nodeSet struct {
	values map[NodeID]int
	orders []NodeID
}

func newNodeSet(vals ...NodeID) nodeSet {
	os := nodeSet{
		values: make(map[NodeID]int, len(vals)),
		orders: make([]NodeID, 0, len(vals)),
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

func (o nodeSet) Put(v NodeID) nodeSet {
	newOrders := make([]NodeID, 0, len(o.orders)+1)
	newOrders = append(newOrders, o.orders...)
	newOrders = append(newOrders, v)

	return newNodeSet(newOrders...)
}

func (o nodeSet) HasNotLast(v NodeID) bool {
	if i, ok := o.values[v]; ok && i != len(o.orders)-1 {
		return true
	}

	return false
}

func (o nodeSet) GetLast() NodeID {
	if len(o.orders) == 0 {
		return ""
	}

	return o.orders[len(o.orders)-1]
}

func (o nodeSet) Unwind(v NodeID) []NodeID {
	return o.orders[o.values[v]:]
}
