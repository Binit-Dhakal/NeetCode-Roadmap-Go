package trees

import (
	"strconv"
	"strings"
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) Serialize(root *TreeNode) string {
	var res []string

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, "N")
			return
		}
		res = append(res, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(res, ",")
}

func (this *Codec) Deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	i := 0
	var rec func() *TreeNode

	rec = func() *TreeNode {
		if vals[i] == "N" {
			i++
			return nil
		}

		val, _ := strconv.Atoi(vals[i])
		node := &TreeNode{Val: val}
		i++

		node.Left = rec()
		node.Right = rec()

		return node
	}
	return rec()
}
