package main

import "testing"

func TestDll(t *testing.T) {

	list := &List{}

	if list.size != 0 {
		t.Errorf("List is not empty")
	}

	list.InsertAtFront("item1")

	if list.size != 1 {
		t.Errorf("List size is not correct")
	}

	if item := list.Find("item1"); item != nil && item.Value != "item1" {
		t.Errorf("Item has not been found")
	}

	list.RemoveFromFront()

	if list.size != 0 {
		t.Errorf("List is not empty")
	}

	list.InsertAtFront("item0")
	list.InsertAtFront("item2")

	if list.size != 2 {
		t.Errorf("List size is not correct")
	}

	list.RemoveFromFront()
	list.RemoveFromFront()

	if list.size != 0 {
		t.Errorf("List is not empty")
	}

	list.InsertAtBack("item3")
	list.InsertAtBack("item4")
	list.InsertAtBack("item5")

	item3 := "item3"
	if item := list.Find(item3); item != nil && item.Value != item3 {
		t.Errorf("Item has not been found")
	}

	item4 := "item4"
	if item := list.Find(item4); item != nil && item.Value != item4 {
		t.Errorf("Item has not been found")
	}

	item5 := "item5"
	if item := list.Find(item5); item != nil && item.Value != item5 {
		t.Errorf("Item has not been found")
	}

	list.Remove(item3)

	if item := list.Find(item3); item != nil {
		t.Errorf("Item has been found when it should not")
	}

	list.RemoveFromBack()
	list.RemoveFromBack()

	if list.size != 0 {
		t.Errorf("List is not empty")
	}

	list.Print()

	list.InsertAtFront("item0")
	list.InsertAtFront("item2")

	list.InsertNode("item1", list.head, list.head.next)

	list.InsertNode("item11", list.head, nil)

	list.InsertAfter("item8", 1)

	if err := list.InsertAfter("item90", 10); err == nil {
		t.Errorf("Error was not raised %v", err)
	}

	if list.size != 5 {
		t.Errorf("List has not the correct size: %v", list.size)
	}

	list.Remove("item0")
	list.Remove("item8")
	list.Remove("item11")

	list = &List{}

	list.InsertAtFront("head")
	list.RemoveFromBack()
	if list.size != 0 {
		t.Errorf("List is not empty")
	}

	list.InsertAtFront("init")
	list.InsertAtFront("mid")
	list.InsertAtFront("end")
	list.RemoveNode(list.head.next)
	if list.size != 2 {
		t.Errorf("List is not empty")
	}
}
