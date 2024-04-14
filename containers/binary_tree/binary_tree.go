package binary_tree

import (
	"cmp"
	"fmt"
)

type TreeNode[T cmp.Ordered] struct {
	value T
	left  *TreeNode[T]
	right *TreeNode[T]
}

type BinaryTree[T cmp.Ordered] struct {
	root *TreeNode[T]
}

func (t *BinaryTree[T]) Add(value T) {
	t.root = add(t.root, value)
}

func add[T cmp.Ordered](node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		node = &TreeNode[T]{value: value}
	}
	if node.value < value {
		node.right = add(node.right, value)
	} else if node.value > value {
		node.left = add(node.left, value)
	}
	return node
}

func (t *BinaryTree[T]) Invert() {
	t.root = invert(t.root)
}

func invert[T cmp.Ordered](node *TreeNode[T]) *TreeNode[T] {
	if node == nil {
		return node
	}

	left := invert(node.left)
	right := invert(node.right)

	node.left = right
	node.right = left

	return node
}

func (t *BinaryTree[T]) Depth() int {
	return depth(t.root, 0)
}

func depth[T cmp.Ordered](node *TreeNode[T], level int) int {
	if node == nil {
		return level
	}
	rightLevel := depth(node.right, level+1)
	leftLevel := depth(node.left, level+1)

	if rightLevel > leftLevel {
		return rightLevel
	}
	return leftLevel
}

func (t *BinaryTree[T]) Print() {
	if t.root == nil {
		return
	}

	queue := []*TreeNode[T]{t.root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[i]
			fmt.Printf("%v ", node.value)

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		fmt.Println() // Print newline after each level

		// Update the queue to contain only nodes of the next level
		queue = queue[levelSize:]
	}
}

func New[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}
