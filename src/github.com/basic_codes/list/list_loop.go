package list

import (
	"container/list"
	"fmt"
)

func PrintList() {
	var mylist list.List
	mylist.PushBack(1)
	mylist.PushBack(2)
	mylist.PushBack(3)
	mylist.PushBack(4)
	mylist.PushBack(5)

	for elem := mylist.Front(); elem != nil; elem = elem.Next() {
		fmt.Println(elem.Value.(int))
	}

}

// package main

// // importing fmt and container list packages
// import (
// 	"container/list"
// 	"fmt"
// )

// // main method
// func main() {
// 	var intList list.List
// 	intList.PushBack(11)
// 	intList.PushBack(23)
// 	intList.PushBack(34)

// 	for element := intList.Front(); element != nil; element = element.Next() {
// 		fmt.Println(element.Value.(int))
// 	}
// }
