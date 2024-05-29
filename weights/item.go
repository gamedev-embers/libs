package weights

type Item interface {
	comparable
	GetWeight() int64
}

// WeightItem 包含 Item 及其累计权重
type WeightItem[T Item] struct {
	Item              T
	AccumulatedWeight int64
}

// func _sort[T Item](items []T) {

// }
