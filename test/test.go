package main

import (
	"log"
	"tinList"
)

func main() {
	l := tinList.NewList()

	l.
		PushBack("ttt").
		PushBack("bbb").
		PushBack("fff").
		PushFront(0)

	ne := l.PushGetBack("aaa")
	printList(l)

	ne = l.MoveFront(ne)
	log.Println("================")
	printList(l)
	l.Delete(ne)
	log.Println("================")
	printList(l)

}

func printList(l *tinList.TinList) {
	for e := l.GetFirst(); e != nil; e = e.GetNext() {

		intV, intOk := e.Value.(int)
		if intOk {
			log.Println("int", intV)
		}
		v, ok := e.Value.(string)
		if ok {
			log.Println("string", v)
		}
	}
}
