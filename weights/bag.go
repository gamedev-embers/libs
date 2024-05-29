package weights

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

type Bag[T Item] struct {
	src         []T
	items       []*WeightItem[T]
	totalWeight int64
}

func New[T Item](items []T) *Bag[T] {
	sort.Slice(items, func(i, j int) bool {
		return items[i].GetWeight() < items[j].GetWeight()
	})
	obj := &Bag[T]{src: items}
	_items, totalWeight := obj.buildItems(items)
	obj.items = _items
	obj.totalWeight = totalWeight
	return obj
}

func (bag *Bag[T]) buildItems(items []T) ([]*WeightItem[T], int64) {
	rs := make([]*WeightItem[T], len(items))
	var totalWeight int64
	for i, item := range items {
		totalWeight += item.GetWeight()
		rs[i] = &WeightItem[T]{
			Item:              item,
			AccumulatedWeight: totalWeight,
		}
	}
	return rs, totalWeight
}

func (bag *Bag[T]) DropOne(r *rand.Rand) (T, error) {
	randNum := rand.Int63n(bag.totalWeight)
	idx := bag.binarySearch(bag.items, randNum)
	if idx >= len(bag.items) {
		var _default T
		return _default, errors.New("index out of range")
	}
	return bag.items[idx].Item, nil
}

func (bag *Bag[T]) DropMany(r *rand.Rand, count int) ([]T, error) {
	if count <= 0 {
		return nil, fmt.Errorf("invalid count %d", count)
	}
	var items, totalWeight = bag.buildItems(bag.src)
	droppeds := make([]T, 0)
	for i := 0; i < count; i++ {
		if totalWeight <= 0 {
			break
		}
		flag := rand.Int63n(totalWeight)
		idx := bag.binarySearch(items, flag)
		if idx >= len(items) {
			return nil, errors.New("index out of range")
		}
		droppedItem := items[idx].Item
		droppeds = append(droppeds, droppedItem)

		totalWeight -= items[idx].Item.GetWeight()
		items = append(items[:idx], items[idx+1:]...)

		// 更新剩余项的累计权重
		reduceVal := droppedItem.GetWeight()
		for j := idx; j < len(items); j++ {
			items[j].AccumulatedWeight -= reduceVal
		}
	}
	return droppeds, nil
}

// binarySearch 二分查找函数，基于累计权重
func (bag *Bag[T]) binarySearch(items []*WeightItem[T], target int64) int {
	low, high := 0, len(items)-1
	for low <= high {
		mid := low + (high-low)/2
		if items[mid].AccumulatedWeight == target {
			return mid
		} else if items[mid].AccumulatedWeight < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func (bag *Bag[T]) Dump() []*dumpItem[T] {
	rs := make([]*dumpItem[T], len(bag.items))
	for i, item := range bag.items {
		row := &dumpItem[T]{
			Item: item.Item,
		}
		row.Probs.Expect = float64(item.Item.GetWeight()) / float64(bag.totalWeight)
		rs[i] = row
	}
	return rs
}

func (bag *Bag[T]) DryRun(r *rand.Rand, times int) ([]*dumpItem[T], error) {
	if bag.totalWeight <= 0 {
		return nil, errors.New("total weight is less than or equal to 0")
	}

	var drops = make(map[T]int, len(bag.items))
	for i := 0; i < times; i++ {
		item, err := bag.DropOne(r)
		if err != nil {
			return nil, fmt.Errorf("drop one error: %w", err)
		}
		drops[item]++
	}
	dumps := bag.Dump()
	if len(drops) != len(dumps) {
		return nil, errors.New("len(drops) != len(dumps)")
	}
	for _, row := range dumps {
		row.Drops = drops[row.Item]
		row.Probs.Actual = float64(row.Drops) / float64(times)
	}
	return dumps, nil
}
