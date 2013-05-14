package list

// This file only contains code for the list<string>, as defined in the thrift file lists.thrift
// !!!!!!!!!!!!! For information and explaining comments, please see listInt32.go !!!!!!

import ()

type ListString struct {
	list []string
}

func NewListString() *ListString {
	return &ListString{list: make([]string, 0)}
}

func (l *ListString) Len() int {
	return len(l.list)
}

func (l *ListString) At(index int) string {
	return l.list[index]
}

func (l *ListString) Set(index int, data string) {
	if len(l.list) >= index {
		l.list[index] = data
	} else {
		l.list = append(l.list, data)
	}
}

func (l *ListString) Push(data string) {
	l.list = append(l.list, data)
}

func (l *ListString) Pop() string {
	var popval string
	popval, l.list = l.list[len(l.list)-1], l.list[:len(l.list)-1]
	return popval
}

func (l *ListString) Swap(indexA, indexB int) {
	l.list[indexB], l.list[indexA] = l.list[indexA], l.list[indexB]
}

func (l *ListString) Insert(index int, data string) {
	newl := make([]string, 1)
	newl[0] = data
	l.list = append(l.list[:index], append(newl, l.list[index:]...)...)
}

func (l *ListString) Delete(index int) {
	l.list = append(l.list[:index], l.list[index+1:]...)
}

func (l *ListString) Contains(data string) bool {
	return l.IndexOf(data) >= 0
}

func (l *ListString) Less(i, j int) bool {
	return l.list[i] < l.list[j]
}

func (l *ListString) Iter() <-chan string {
	c := make(chan string)
	go l.iterate(c)
	return c
}

func (l *ListString) iterate(c chan<- string) {
	for _, elem := range l.list {
		c <- elem
	}
	close(c)
}

func (l *ListString) IndexOf(data string) int {
	size := len(l.list)
	for i := 0; i < size; i++ {
		if l.list[i] == data {
			return i
		}
	}
	return -1
}
