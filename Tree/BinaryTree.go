package Tree

import "strconv"

type binaryTreeNode struct {
	Val   interface{}
	Left  *binaryTreeNode
	Right *binaryTreeNode
}

func GetBinaryTree(s string) *binaryTreeNode {
	if len(s) == 0 {
		return &binaryTreeNode{}
	}
	return DeSerializeBinaryTreePre(s)
}

//	反序列化二叉树
func DeSerializeBinaryTreePre(str string) *binaryTreeNode {
	var strIndex int = 0
	return DeSerializeBinaryTreePreHelper(str, &strIndex)
}

func DeSerializeBinaryTreePreHelper(str string, index *int) *binaryTreeNode {
	if len(str) == *index {
		return nil
	}
	if (str)[*index] == '#' {
		(*index) = *index + 1
		return nil
	}
	v, _ := strconv.Atoi(string((str)[*index]))
	root := &binaryTreeNode{Val: v}
	(*index) = *index + 1
	root.Left = DeSerializeBinaryTreePreHelper(str, index)
	root.Right = DeSerializeBinaryTreePreHelper(str, index)

	return root
}
