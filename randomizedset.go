// Package randomizedset provides a data structure that supports insert, delete, search and getRandom in constant time.
package randomizedset

import (
	"math/rand"
	"time"
)

type Set struct {
	valueIndexes map[interface{}]int
	values       []interface{}
	size         int
	rand         *rand.Rand
}

func New() *Set {
	newSrc := rand.NewSource(time.Now().UnixNano())
	return &Set{
		valueIndexes: make(map[interface{}]int),
		values:       make([]interface{}, 10),
		size:         0,
		rand:         rand.New(newSrc),
	}
}

func (set *Set) Add(element interface{}) {
	if _, ok := set.valueIndexes[element]; !ok {
		set.valueIndexes[element] = set.size

		// when set.values is full
		if set.size == len(set.values) {
			set.values = append(set.values, make([]interface{}, 10))
		}
		set.values[set.size] = element
		set.size++
	}
}

func (set *Set) Search(element interface{}) bool {
	_, ok := set.valueIndexes[element]
	return ok
}

func (set *Set) GetRandom() interface{} {
	if set.size == 0 {
		return nil
	}
	index := set.rand.Intn(set.size)
	element := set.values[index]
	return element
}

func (set *Set) Delete(element interface{}) {
	index, ok := set.valueIndexes[element]
	if ok {
		lastIndex := set.size - 1
		lastValue := set.values[lastIndex]

		// removing an element from an array is O(n) while  removing the last element is constant
		// swap these two
		// element in set.value is not exactly removed, but it will be overwritten next time inserting an element
		set.values[index] = lastValue
		set.valueIndexes[lastValue] = index
		delete(set.valueIndexes, element)
		set.size--
	}
}
