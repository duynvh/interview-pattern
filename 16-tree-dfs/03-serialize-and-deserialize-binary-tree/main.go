package main

import (
	"fmt"
	"strconv"
	"strings"
)

var MARKER = "M"
var m = 1

func serializeRec(node *EduBinaryTreeNode, stream *[]string) {
	if node == nil {
		(*stream) = append((*stream), MARKER+strconv.Itoa(m))
		m += 1
		return
	}

	(*stream) = append((*stream), strconv.Itoa(node.data))
	serializeRec(node.left, stream)
	serializeRec(node.right, stream)
}

func serialize(root *EduBinaryTreeNode) []string {
	stream := make([]string, 0)
	serializeRec(root, &stream)
	return stream
}

func deserialize(stream *[]string) *EduBinaryTreeNode {
	val := (*stream)[0]
	*stream = (*stream)[1:len(*stream)]

	if string(val[0]) == MARKER {
		return nil
	}

	tempVal, _ := strconv.Atoi(val)
	node := &EduBinaryTreeNode{data: tempVal}

	node.left, node.right = deserialize(stream), deserialize(stream)

	return node
}

func main() {
	inputTrees := [][]int{
		{1, 2, 3, 4, 5},
		{100, 50, 200, 25, 75, 350},
		{100, 200, 75, 50, 25, 350},
		{200, 350, 100, 75, 50, 200, 25},
		{100, 50, 200, 25, 75, 350},
		{25, 50, 75, 100, 200, 350},
		{350, 75, 25, 200, 50, 100},
	}

	for i, input := range inputTrees {
		tree := new(EduBinaryTree)
		tree.BinaryTreeInitWithSlice(input)
		if i > 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%d.\tBinary tree:\n", i+1)
		fmt.Printf("\n\tMarker used for NULL nodes in serialization/deserialization: %s*\n", MARKER)

		// Serialization
		ostream := serialize(tree.root)
		fmt.Print("\n\tSerialized integer array:\n")
		fmt.Print("\t", strings.Replace(fmt.Sprint(ostream), " ", ", ", -1))

		// Deserialization
		deserializedRoot := deserialize(&ostream)
		fmt.Print("\n\n\tDeserialized binary tree:\n", deserializedRoot)

		fmt.Printf("%s\n", strings.Repeat("-", 100))
		m = 1
	}
}
