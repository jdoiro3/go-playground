package main

import "fmt"

type list []interface{}

func newList(c int) *list {
	v := make(list, 0, c)
	return &v
}

func (listPtr *list) append(val interface{}) {
	l := *listPtr
	if len(l) == cap(l) {
		fmt.Println("Adding capacity...")
		newList := make(list, len(l), cap(l)*2)
		for i, _ := range l {
			newList[i] = l[i]
		}
		*listPtr = newList
		l = newList
	}
	*listPtr = l[:len(l)+1]
	(*listPtr)[len(l)] = val
}

func main() {
	v := newList(10)
	for i := 0; i < 20; i++ {
		fmt.Println(len(*v), cap(*v))
		v.append(i + 1)
	}
	fmt.Println(*v)
}
