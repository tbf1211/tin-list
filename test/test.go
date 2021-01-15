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
