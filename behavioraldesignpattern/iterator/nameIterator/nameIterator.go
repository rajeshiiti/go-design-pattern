package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() string
}

type NameIterator struct {
	Names []string
	Index int
}

func (ni *NameIterator) HasNext() bool {
	return ni.Index < len(ni.Names)
}

func (ni *NameIterator) Next() string {
	if ni.HasNext() {
		name := ni.Names[ni.Index]
		ni.Index++
		return name
	}
	return ""
}

type NameCollection struct {
	names []string
}

func (nc *NameCollection) Add(val string) {
	nc.names = append(nc.names, val)
}

func (nc *NameCollection) GetIterator() Iterator {
	return &NameIterator{Names: nc.names, Index: 0}
}

func main() {
	collection := NameCollection{}
	collection.Add("ram")
	collection.Add("ramesh")
	collection.Add("suresh")
	collection.Add("shyam")

	iterator := collection.GetIterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

}
