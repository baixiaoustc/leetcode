package leetcode

import (
	"fmt"
	"testing"
)

func TestMyLinkedList1(t *testing.T) {
	obj := ConstructorMyLinkedList()
	obj.AddAtHead(0)
	fmt.Println(obj.Get(0))
	obj.AddAtHead(1)
	fmt.Println(obj.Get(0))
	obj.Print()
	obj.AddAtHead(2)
	obj.Print()
	obj.AddAtTail(0)
	obj.AddAtTail(1)
	obj.AddAtTail(2)
	obj.Print()
	obj.AddAtIndex(0, 100)
	obj.AddAtIndex(6, 200)
	obj.Print()
	obj.AddAtIndex(8, 300)
	obj.Print()
	obj.DeleteAtIndex(8)
	obj.Print()
	obj.DeleteAtIndex(7)
	obj.DeleteAtIndex(0)
	obj.Print()
}

//["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
//[[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
func TestMyLinkedList2(t *testing.T) {
	obj := ConstructorMyLinkedList()
	obj.AddAtHead(7)
	obj.Print()
	obj.AddAtHead(2)
	obj.Print()
	obj.AddAtHead(1)
	obj.Print()
	obj.AddAtIndex(3, 0)
	obj.Print()
	obj.DeleteAtIndex(2)
	obj.Print()
	obj.AddAtHead(6)
	obj.Print()
	obj.AddAtTail(4)
	obj.Print()
	fmt.Println(obj.Get(4))
	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)
}
