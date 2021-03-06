package singly

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func getStdoutLogs(fn func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return fmt.Sprintf("%s", out)
}

func linkNodes(length int) *Node {
	if length == 0 {
		return nil
	}
	node := Node{value: 1}
	for i, initial := 1, &node; i < length; i++ {
		nextNode := initial.Push(&Node{value: i + 1})
		initial = nextNode
	}
	return &node
}

func TestSingly(t *testing.T) {
	first := linkNodes(3)
	output := getStdoutLogs(first.Print)

	if output != "1\n2\n3\n" {
		t.Errorf("expected %s got %s", "1\n2\n3", output)
	}
}

func TestSinglyLinkedList_HeadIntersion(t *testing.T) {
	first := linkNodes(3)

	newHead := Node{value: 4, next: first}

	output := getStdoutLogs(newHead.Print)

	if output != "4\n1\n2\n3\n" {
		t.Fatal("failed to insert node at head")
	}
}

func TestSinglyLinkedList_Push(t *testing.T) {
	first := Node{value: 1, next: nil}
	second := first.Push(&Node{value: 2})
	_ = second.Push(&Node{value: 3})
	output := getStdoutLogs(first.Print)
	expectedOutput := "1\n2\n3\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", output, expectedOutput)
	}
}

func TestSinglyLinkedList_DeleteKeyInMiddle(t *testing.T) {
	first := linkNodes(4)
	// should delete an element in the middle
	if first.DeleteKey(2) != true {
		t.Errorf("Failed to delete key %d\n", 2)
	}
	output := getStdoutLogs(first.Print)
	expectedOutput := "1\n3\n4\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyLinkedList_DeleteKeyAtEnd(t *testing.T) {
	first := linkNodes(4)
	// should delete an element at the end
	if first.DeleteKey(4) != true {
		t.Errorf("Failed to delete key %d\n", 4)
	}
	output := getStdoutLogs(first.Print)
	expectedOutput := "1\n2\n3\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyLinkedList_DeleteKeyAtStart(t *testing.T) {
	first := linkNodes(4)
	// should delete an element at the end
	if first.DeleteKey(1) != true {
		t.Errorf("Failed to delete key %d\n", 1)
	}
	output := getStdoutLogs(first.Print)
	expectedOutput := "2\n3\n4\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyLinkedList_ShouldNotDelete(t *testing.T) {
	first := linkNodes(4)
	// should delete an element at the end
	if first.DeleteKey(6) != false {
		t.Errorf("Accidentally deleted one key")
	}
	output := getStdoutLogs(first.Print)
	expectedOutput := "1\n2\n3\n4\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}
func TestSinglyLinkedList_DeleteElementAtIndexInMiddle(t *testing.T) {
	first := linkNodes(4)
	// should delete an element in the middle
	if first.DeleteElementAtIndex(1) != true {
		t.Errorf("Failed to delete key %d\n", 2)
	}
	output := getStdoutLogs(first.Print)
	expectedOutput := "1\n3\n4\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyLinkedList_DeleteElementAtIndexAtEnd(t *testing.T) {
	first := linkNodes(4)
	// should delete an element at the end
	if first.DeleteElementAtIndex(3) != true {
		t.Errorf("Failed to delete key %d\n", 4)
	}
	output := getStdoutLogs(first.Print)
	expectedOutput := "1\n2\n3\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyLinkedList_DeleteElementAtIndexAtStart(t *testing.T) {
	first := linkNodes(4)
	// should delete an element at the end
	if first.DeleteElementAtIndex(0) != true {
		t.Errorf("Failed to delete key %d\n", 1)
	}
	output := getStdoutLogs(first.Print)
	expectedOutput := "2\n3\n4\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyLinkedList_ShouldNotDeleteAtIndex(t *testing.T) {
	first := linkNodes(4)
	// should delete an element at the end
	if first.DeleteElementAtIndex(6) != false {
		t.Errorf("Accidentally deleted one key")
	}
	output := getStdoutLogs(first.Print)
	expectedOutput := "1\n2\n3\n4\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyListList_MakeMiddleNodeFirst(t *testing.T) {
	first := linkNodes(5)
	newHead := first.MakeMiddleNodeFirst()
	output := getStdoutLogs(newHead.Print)
	expectedOutput := "3\n1\n2\n4\n5\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyListList_MakeMiddleNodeFirstWithTwoNodes(t *testing.T) {
	first := linkNodes(2)
	newHead := first.MakeMiddleNodeFirst()
	output := getStdoutLogs(newHead.Print)
	expectedOutput := "2\n1\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyListList_MakeMiddleNodeFirstWithSingleNodes(t *testing.T) {
	first := linkNodes(1)
	newHead := first.MakeMiddleNodeFirst()
	output := getStdoutLogs(newHead.Print)
	expectedOutput := "1\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}

func TestSinglyListList_MergeTwoSortedListCorrectly(t *testing.T) {
	first := linkNodes(3)
	second := linkNodes(5)
	newHead := MergeTwoSortedList(first, second)
	output := getStdoutLogs(newHead.Print)
	expectedOutput := "1\n1\n2\n2\n3\n3\n4\n5\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}
func TestSinglyListList_MergeTwoSortedListWithTwoElement(t *testing.T) {
	first := &Node{value: 3}
	first.Push(&Node{value: 5})

	second := &Node{value: 7}
	second.Push(&Node{value: 11})
	newHead := MergeTwoSortedList(first, second)
	output := getStdoutLogs(newHead.Print)
	expectedOutput := "3\n5\n7\n11\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", expectedOutput, output)
	}
}
