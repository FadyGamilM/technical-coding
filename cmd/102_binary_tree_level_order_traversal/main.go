package main

import (
	"log"

	"github.com/golang-collections/collections/queue"
)

func main() {
	t1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	log.Println(levelOrder(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	out := [][]int{}
	if root == nil {
		return out
	}

	Q := queue.New()
	Q.Enqueue(root)

	for Q.Len() > 0 {
		// we define a queue level that will represents a level in the tree
		// we loop through the original queue which we are traversing the tree in BFS manner, with the length of the number of nodes in this origianl queue currently
		levelQ := []int{}
		levelLength := Q.Len()
		for i := 0; i < levelLength; i++ {
			n := Q.Dequeue().(*TreeNode)
			// add it to the level queue
			// push their childs (left then right) to the origianl queue for the next level iteration
			// continue
			levelQ = append(levelQ, n.Val)

			if n.Left != nil {
				Q.Enqueue(n.Left)
			}
			if n.Right != nil {
				Q.Enqueue(n.Right)
			}
		}
		// after we finished processing this level, we can stsore it into the output result
		out = append(out, levelQ)
	}

	return out
}

// [3,9,20,null,null,15,7]
