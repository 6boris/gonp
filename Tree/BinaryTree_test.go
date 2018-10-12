package Tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		root := GetBinaryTree("124##5##36##")

		fmt.Println(root.Val)
	})
}
