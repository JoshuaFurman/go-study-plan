# 16. Trees & Graphs

This example demonstrates binary trees, tree traversals, and graph algorithms essential for coding interviews. Trees and graphs are fundamental data structures used extensively in computer science problems.

## Key Concepts

### Binary Trees
- **Binary Tree**: Each node has at most two children (left, right)
- **Binary Search Tree (BST)**: Left subtree < root < right subtree
- **Complete Tree**: All levels filled except possibly last level
- **Balanced Tree**: Height difference between subtrees ≤ 1

### Tree Traversals
- **In-order**: Left → Root → Right (sorted order for BST)
- **Pre-order**: Root → Left → Right (copy tree structure)
- **Post-order**: Left → Right → Root (delete tree)
- **Level-order**: Breadth-first traversal using queue

### Graph Representations
- **Adjacency List**: Most common, space-efficient
- **Adjacency Matrix**: Dense graphs, fast lookups
- **Edge List**: Simple, good for some algorithms

## Binary Tree Operations

### Tree Construction
```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Insert into BST
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
```

### Tree Traversals
```go
// In-order traversal (recursive)
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

// Level-order traversal (iterative with queue)
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
```

## Tree Properties & Algorithms

### Tree Height & Balance
```go
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

func IsBalanced(root *TreeNode) bool {
    return checkBalance(root) != -1
}

func checkBalance(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := checkBalance(root.Left)
    if left == -1 {
        return -1
    }
    right := checkBalance(root.Right)
    if right == -1 {
        return -1
    }
    if abs(left-right) > 1 {
        return -1
    }
    return max(left, right) + 1
}
```

### Tree Symmetry & Diameter
```go
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

func DiameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    var height func(*TreeNode) int
    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := height(node.Left)
        right := height(node.Right)
        diameter := left + right
        if diameter > maxDiameter {
            maxDiameter = diameter
        }
        return max(left, right) + 1
    }
    height(root)
    return maxDiameter
}
```

## Graph Algorithms

### Graph Representation
```go
type Graph struct {
    Vertices int
    AdjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
    return &Graph{
        Vertices: vertices,
        AdjList:  make(map[int][]int),
    }
}

func (g *Graph) AddEdge(v1, v2 int) {
    g.AdjList[v1] = append(g.AdjList[v1], v2)
    g.AdjList[v2] = append(g.AdjList[v2], v1) // Undirected
}
```

### Depth-First Search (DFS)
```go
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
```

### Breadth-First Search (BFS)
```go
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
```

### Shortest Path (BFS for unweighted graphs)
```go
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
```

### Connected Components
```go
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
```

### Topological Sort (DAG)
```go
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
```

## Time & Space Complexity

### Tree Operations
| Operation | Time | Space |
|-----------|------|-------|
| Insert (BST) | O(h) | O(h) |
| Search (BST) | O(h) | O(h) |
| In-order traversal | O(n) | O(h) |
| Level-order traversal | O(n) | O(w) |

### Graph Operations
| Operation | Time | Space |
|-----------|------|-------|
| DFS | O(V + E) | O(V) |
| BFS | O(V + E) | O(V) |
| Shortest Path (BFS) | O(V + E) | O(V) |
| Connected Components | O(V + E) | O(V) |
| Topological Sort | O(V + E) | O(V) |

## Common Patterns

1. **Tree Recursion**: Base case + recursive calls on subtrees
2. **Graph Traversal**: DFS for exploration, BFS for shortest paths
3. **Backtracking**: DFS with state restoration
4. **Tree Construction**: Build from traversals or arrays
5. **Graph Construction**: Build adjacency list from edges

## When to Use

### Trees
- Hierarchical data
- Sorted data (BST)
- Expression parsing
- Decision trees
- File system representation

### Graphs
- Social networks
- Web crawling
- Network routing
- Dependency resolution
- State machines

## Practice Problems

### Easy Trees
- Maximum Depth of Binary Tree
- Same Tree
- Symmetric Tree
- Path Sum

### Medium Trees
- Binary Tree Level Order Traversal
- Validate Binary Search Tree
- Lowest Common Ancestor
- Binary Tree Right Side View

### Easy Graphs
- Number of Islands
- Clone Graph
- Course Schedule (topological sort)

### Medium Graphs
- Pacific Atlantic Water Flow
- Word Ladder
- Network Delay Time

### Hard
- Serialize and Deserialize Binary Tree
- Word Search II (Trie + DFS)
- Alien Dictionary (topological sort)

## Running the Example

```bash
cd 16-trees-graphs
go run main.go
```

This example provides comprehensive coverage of tree and graph algorithms, essential for technical interviews at top companies.