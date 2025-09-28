package main

import (
	"fmt"
)

// TreeNode represents a node in binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Graph represents an adjacency list graph
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// NewGraph creates a new graph with given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

// AddEdge adds an undirected edge between two vertices
func (g *Graph) AddEdge(v1, v2 int) {
	g.AdjList[v1] = append(g.AdjList[v1], v2)
	g.AdjList[v2] = append(g.AdjList[v2], v1)
}

// Binary Tree Operations

// Insert inserts a value into binary search tree
func (root *TreeNode) Insert(val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = root.Left.Insert(val)
	} else {
		root.Right = root.Right.Insert(val)
	}

	return root
}

// Search searches for a value in binary search tree
func (root *TreeNode) Search(val int) bool {
	if root == nil {
		return false
	}

	if val == root.Val {
		return true
	} else if val < root.Val {
		return root.Left.Search(val)
	} else {
		return root.Right.Search(val)
	}
}

// Tree Traversals

// InOrderTraversal performs in-order traversal (Left, Root, Right)
func InOrderTraversal(root *TreeNode) []int {
	var result []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			result = append(result, node.Val)
			inorder(node.Right)
		}
	}
	inorder(root)
	return result
}

// PreOrderTraversal performs pre-order traversal (Root, Left, Right)
func PreOrderTraversal(root *TreeNode) []int {
	var result []int
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node != nil {
			result = append(result, node.Val)
			preorder(node.Left)
			preorder(node.Right)
		}
	}
	preorder(root)
	return result
}

// PostOrderTraversal performs post-order traversal (Left, Right, Root)
func PostOrderTraversal(root *TreeNode) []int {
	var result []int
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node != nil {
			postorder(node.Left)
			postorder(node.Right)
			result = append(result, node.Val)
		}
	}
	postorder(root)
	return result
}

// LevelOrderTraversal performs level-order (breadth-first) traversal
func LevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var level []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}

// Tree Properties

// GetHeight returns the height of the binary tree
func GetHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := GetHeight(root.Left)
	rightHeight := GetHeight(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// IsBalanced checks if binary tree is height-balanced
func IsBalanced(root *TreeNode) bool {
	return checkBalance(root) != -1
}

func checkBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := checkBalance(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := checkBalance(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// IsSymmetric checks if binary tree is symmetric
func IsSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return (t1.Val == t2.Val) &&
		isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left)
}

// Binary Tree Problems

// MaxDepth returns the maximum depth of binary tree
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// MinDepth returns the minimum depth of binary tree
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}

	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}

	leftDepth := MinDepth(root.Left)
	rightDepth := MinDepth(root.Right)

	if leftDepth < rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// DiameterOfBinaryTree returns the diameter of binary tree
func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := height(node.Left)
		rightHeight := height(node.Right)

		// Update diameter
		diameter := leftHeight + rightHeight
		if diameter > maxDiameter {
			maxDiameter = diameter
		}

		// Return height
		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}

	height(root)
	return maxDiameter
}

// InvertTree inverts a binary tree
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}

// LowestCommonAncestor finds LCA of two nodes in BST
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestor(root.Right, p, q)
	}

	return root
}

// Graph Algorithms

// DFSTraversal performs DFS traversal from given vertex
func (g *Graph) DFSTraversal(start int) []int {
	visited := make(map[int]bool)
	var result []int

	var dfs func(int)
	dfs = func(vertex int) {
		visited[vertex] = true
		result = append(result, vertex)

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(start)
	return result
}

// BFSTraversal performs BFS traversal from given vertex
func (g *Graph) BFSTraversal(start int) []int {
	visited := make(map[int]bool)
	var result []int
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// HasPath checks if there's a path between two vertices using DFS
func (g *Graph) HasPath(start, end int) bool {
	visited := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(vertex int) bool {
		if vertex == end {
			return true
		}

		visited[vertex] = true

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				if dfs(neighbor) {
					return true
				}
			}
		}

		return false
	}

	return dfs(start)
}

// ShortestPath finds shortest path using BFS (unweighted graph)
func (g *Graph) ShortestPath(start, end int) []int {
	visited := make(map[int]bool)
	parent := make(map[int]int)
	queue := []int{start}
	visited[start] = true
	parent[start] = -1

	found := false
	for len(queue) > 0 && !found {
		vertex := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = vertex
				queue = append(queue, neighbor)

				if neighbor == end {
					found = true
					break
				}
			}
		}
	}

	if !found {
		return nil
	}

	// Reconstruct path
	var path []int
	current := end
	for current != -1 {
		path = append([]int{current}, path...)
		current = parent[current]
	}

	return path
}

// ConnectedComponents finds all connected components using DFS
func (g *Graph) ConnectedComponents() [][]int {
	visited := make(map[int]bool)
	var components [][]int

	for vertex := 0; vertex < g.Vertices; vertex++ {
		if !visited[vertex] {
			var component []int

			var dfs func(int)
			dfs = func(v int) {
				visited[v] = true
				component = append(component, v)

				for _, neighbor := range g.AdjList[v] {
					if !visited[neighbor] {
						dfs(neighbor)
					}
				}
			}

			dfs(vertex)
			components = append(components, component)
		}
	}

	return components
}

// TopologicalSort performs topological sort using DFS
func (g *Graph) TopologicalSort() []int {
	visited := make(map[int]bool)
	stack := []int{}

	var dfs func(int)
	dfs = func(vertex int) {
		visited[vertex] = true

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}

		stack = append([]int{vertex}, stack...)
	}

	for vertex := 0; vertex < g.Vertices; vertex++ {
		if !visited[vertex] {
			dfs(vertex)
		}
	}

	return stack
}

// Helper functions
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("=== Trees & Graphs Demo ===\n")

	// 1. Binary Search Tree Operations
	fmt.Println("1. Binary Search Tree Operations:")
	var bst *TreeNode
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}

	for _, val := range values {
		bst = bst.Insert(val)
	}

	fmt.Printf("Inserted values: %v\n", values)
	fmt.Printf("In-order traversal: %v\n", InOrderTraversal(bst))
	fmt.Printf("Pre-order traversal: %v\n", PreOrderTraversal(bst))
	fmt.Printf("Post-order traversal: %v\n", PostOrderTraversal(bst))
	fmt.Printf("Level-order traversal: %v\n", LevelOrderTraversal(bst))
	fmt.Printf("Search for 6: %t\n", bst.Search(6))
	fmt.Printf("Search for 9: %t\n\n", bst.Search(9))

	// 2. Tree Properties
	fmt.Println("2. Tree Properties:")
	fmt.Printf("Tree height: %d\n", GetHeight(bst))
	fmt.Printf("Is balanced: %t\n", IsBalanced(bst))
	fmt.Printf("Is symmetric: %t\n", IsSymmetric(bst))
	fmt.Printf("Maximum depth: %d\n", MaxDepth(bst))
	fmt.Printf("Minimum depth: %d\n", MinDepth(bst))
	fmt.Printf("Diameter: %d\n\n", DiameterOfBinaryTree(bst))

	// 3. Tree Transformations
	fmt.Println("3. Tree Transformations:")
	fmt.Printf("Original in-order: %v\n", InOrderTraversal(bst))
	inverted := InvertTree(bst)
	fmt.Printf("Inverted in-order: %v\n\n", InOrderTraversal(inverted))

	// 4. Graph Operations
	fmt.Println("4. Graph Operations:")
	graph := NewGraph(6)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 5)
	graph.AddEdge(4, 5)

	fmt.Printf("DFS from 0: %v\n", graph.DFSTraversal(0))
	fmt.Printf("BFS from 0: %v\n", graph.BFSTraversal(0))
	fmt.Printf("Has path 0->5: %t\n", graph.HasPath(0, 5))
	fmt.Printf("Shortest path 0->5: %v\n", graph.ShortestPath(0, 5))
	fmt.Printf("Connected components: %v\n\n", graph.ConnectedComponents())

	// 5. DAG Topological Sort
	fmt.Println("5. DAG Topological Sort:")
	dag := NewGraph(6)
	dag.AddEdge(5, 2)
	dag.AddEdge(5, 0)
	dag.AddEdge(4, 0)
	dag.AddEdge(4, 1)
	dag.AddEdge(2, 3)
	dag.AddEdge(3, 1)

	fmt.Printf("Topological sort: %v\n\n", dag.TopologicalSort())

	// 6. Lowest Common Ancestor
	fmt.Println("6. Lowest Common Ancestor:")
	// Create a simple BST for LCA
	var lcaTree *TreeNode
	lcaValues := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}
	for _, val := range lcaValues {
		lcaTree = lcaTree.Insert(val)
	}

	node2 := &TreeNode{Val: 2}
	node8 := &TreeNode{Val: 8}
	lca := LowestCommonAncestor(lcaTree, node2, node8)
	if lca != nil {
		fmt.Printf("LCA of 2 and 8: %d\n", lca.Val)
	}

	node2_2 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4}
	lca2 := LowestCommonAncestor(lcaTree, node2_2, node4)
	if lca2 != nil {
		fmt.Printf("LCA of 2 and 4: %d\n\n", lca2.Val)
	}

	fmt.Println("=== Demo Complete ===")
}
