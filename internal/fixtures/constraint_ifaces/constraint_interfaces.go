package skipconstraintifaces

import "golang.org/x/exp/constraints"

type Skip1 constraints.Ordered

type Skip2 interface {
	~int
}

type Skip3 interface {
	constraints.Float
}

type Skip4 interface {
	constraints.Float | constraints.Integer
}
