package main

import (
	"fmt"

	"github.com/Ayobami6/linkedlist_implementation/linkedlist"
)

func main() {

	list := linkedlist.NewLinkedList()
	list.InsertAtTheStart(4)
	list.InsertAtTheStart(5)
	list.InsertAtTheStart(9)

	fmt.Println(list.GetCount())
	list.InsertAtTheEnd(3)
	fmt.Println(list.GetCount())
	list.DumpList()

}
