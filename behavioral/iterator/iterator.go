package iterator

import "fmt"

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type ListIterator struct {
	list []interface{}
	index int
}

func NewListIterator(list []interface{}) *ListIterator {
	return &ListIterator{
		list: list,
	}
}

func (l *ListIterator)HasNext() bool {
	return l.index < len(l.list)
}

func (l *ListIterator)Next() interface{} {
	if !l.HasNext() {
		panic("No Data")
	}

	val := l.list[l.index]
	l.index++
	return val
}


type Aggregate interface {
	Iterator() Iterator
}

type ListAggregate struct {
	start int
	end int
}

func NewListAggregate(start, end int) *ListAggregate {
	return &ListAggregate{
		start: start,
		end: end,
	}
}

func (l *ListAggregate)Iterator() Iterator {
	list := make([]interface{}, 0)
	for i := l.start; i <= l.end; i++ {
		list = append(list, i)
	}

	return NewListIterator(list)
}


func printIterator(l *ListAggregate) {
	iter := l.Iterator()
	for iter.HasNext() {
		val := iter.Next()
		fmt.Printf("%v ", val)
	}

	fmt.Println()
}