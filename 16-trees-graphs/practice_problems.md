# Practice Problems - Trees and Graphs

These practice problems reinforce tree and graph algorithms. Focus on traversal techniques and recursive problem solving.

## Problem 1: Binary Tree Inorder Traversal

**Description:** Perform inorder traversal of a binary tree and return the node values in the correct order.

**Requirements:**
- Implement inorder traversal (left, root, right)
- Handle empty trees and single-node trees
- Use recursive approach
- Return slice of node values

**Examples:**
- Tree:     1
         /   \
        2     3
       / \
      4   5
- Inorder: [4, 2, 5, 1, 3]

**Constraints:**
- Tree nodes contain integer values
- Tree can be empty or have up to 100 nodes

### Solution

```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorderHelper(root, &result)
	return result
}

func inorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	// Left subtree
	inorderHelper(node.Left, result)

	// Current node
	*result = append(*result, node.Val)

	// Right subtree
	inorderHelper(node.Right, result)
}

func createTree() *TreeNode {
	// Create tree:     1
	//               /   \
	//              2     3
	//             / \
	//            4   5

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	return root
}

func main() {
	tree := createTree()

	fmt.Println("Tree structure:")
	fmt.Println("     1")
	fmt.Println("   /   \\")
	fmt.Println("  2     3")
	fmt.Println(" / \\")
	fmt.Println("4   5")

	result := inorderTraversal(tree)
	fmt.Printf("\nInorder traversal: %v\n", result)
	fmt.Println("Expected: [4, 2, 5, 1, 3]")
}
```

**Explanation:**
1. Use recursive helper function for inorder traversal
2. Visit left subtree, then current node, then right subtree
3. Pass result slice by reference to avoid copying
4. Handle nil nodes (base case)

## Problem 2: Maximum Depth of Binary Tree

**Description:** Find the maximum depth (height) of a binary tree. The depth of a tree is the number of nodes along the longest path from root to leaf.

**Requirements:**
- Calculate tree height recursively
- Handle empty trees (depth 0)
- Handle single-node trees (depth 1)
- Use recursive approach

**Examples:**
- Tree:     1
         /   \
        2     3
       / \
      4   5
- Maximum depth: 3

**Constraints:**
- Tree can be empty or unbalanced
- Node values are integers

### Solution

```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// Return the larger depth plus 1 for current node
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func createUnbalancedTree() *TreeNode {
	// Create unbalanced tree:
	//     1
	//   /   \
	//  2     3
	//   \
	//    4
	//     \
	//      5

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Right.Right = &TreeNode{Val: 5}

	return root
}

func main() {
	testCases := []*TreeNode{
		createUnbalancedTree(),
		&TreeNode{Val: 1}, // Single node
		nil,                // Empty tree
	}

	for i, tree := range testCases {
		depth := maxDepth(tree)
		fmt.Printf("Test %d: Maximum depth = %d\n", i+1, depth)
	}

	// Detailed output for the unbalanced tree
	fmt.Println("\nUnbalanced tree structure:")
	fmt.Println("     1")
	fmt.Println("   /   \\")
	fmt.Println("  2     3")
	fmt.Println("   \\")
	fmt.Println("    4")
	fmt.Println("     \\")
	fmt.Println("      5")
	fmt.Printf("Maximum depth: %d\n", maxDepth(createUnbalancedTree()))
}
```

**Explanation:**
1. Base case: nil node has depth 0
2. Recursively calculate depth of left and right subtrees
3. Return max(leftDepth, rightDepth) + 1
4. Demonstrates recursive tree traversal