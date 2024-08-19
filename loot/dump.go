package loot

import "fmt"

type dumpItem[T Item] struct {
	Item  T
	Drops int
	Probs struct {
		Expect float64
		Actual float64
	}
}

func (d *dumpItem[T]) String() string {
	p := &d.Probs
	s := fmt.Sprintf("%5.2f%%: %+v", p.Expect*100, d.Item)
	if d.Drops > 0 {
		s += fmt.Sprintf("  drops=%5.2f%%(%d)", p.Actual*100, d.Drops)
	}
	return s
}
