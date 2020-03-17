package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	curr := root
	var succ *TreeNode
	for curr != nil {
		if curr.Val == p.Val {
			if curr.Right != nil {
				return getMinNode(curr.Right)
			}
			break
		} else if p.Val < curr.Val {
			succ = curr
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return succ
}

func getMinNode(node *TreeNode) *TreeNode {

	if node.Left == nil {
		return node
	}

	return getMinNode(node.Left)
}

/*func binarySearch(root *TreeNode, parent *TreeNode, target *TreeNode) (*TreeNode, *TreeNode) {

  if root.Val == target.Val {
    return root, parent
  }


}*/
